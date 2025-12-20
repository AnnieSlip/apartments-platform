package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
	"github.com/elastic/go-elasticsearch/v8"
)

type EsRepo struct {
	Client *elasticsearch.Client
}

func NewEsRepo(client *elasticsearch.Client) *EsRepo {
	return &EsRepo{Client: client}
}

type Filter struct {
	UserID string                 `json:"user_id"`
	Query  map[string]interface{} `json:"query"`
}

// SaveFilter stores a user’s filter as a percolator query in Elasticsearch
func (r *EsRepo) SaveFilter(ctx context.Context, userID string, f models.ApartmentFilter) error {
	// Prepare the document to be stored in the percolator index
	filterDoc := map[string]interface{}{
		"user_id":    userID,
		"created_at": time.Now().UTC().Format(time.RFC3339),
		"query": map[string]interface{}{
			"percolate": map[string]interface{}{
				"field": "query", // the percolator field
				"document": map[string]interface{}{
					"city":             f.City,
					"district":         f.District,
					"price_per_month":  f.MaxPrice, // max price, can add min logic if needed
					"room_numbers":     f.RoomNumbers,
					"bedroom_numbers":  f.BedroomNumbers,
					"bathroom_numbers": f.BathroomNumbers,
				},
			},
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(filterDoc)
	if err != nil {
		return fmt.Errorf("failed to marshal filter: %w", err)
	}

	// Index into Elasticsearch
	res, err := r.Client.Index(
		"filters",
		strings.NewReader(string(data)),
	)
	if err != nil {
		return fmt.Errorf("failed to index filter: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing filter: %s", res.String())
	}

	return nil
}

// MatchApartment checks which saved filters match a given apartment using Elasticsearch percolator.
// runs a percolate query on filters index and lists of userIDs whose filters match this apartment.
func (r *EsRepo) MatchApartment(ctx context.Context, apartment map[string]interface{}) ([]string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"percolate": map[string]interface{}{
				"field":    "query",
				"document": apartment,
			},
		},
	}

	data, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Search(
		r.Client.Search.WithContext(ctx),
		r.Client.Search.WithIndex("filters"),
		r.Client.Search.WithBody(strings.NewReader(string(data))),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result struct {
		Hits struct {
			Hits []struct {
				Source Filter `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	userIDs := make([]string, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		userIDs = append(userIDs, hit.Source.UserID)
	}

	return userIDs, nil
}
