package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Filter struct {
	UserID string                 `json:"user_id"`
	Query  map[string]interface{} `json:"query"`
}

// SaveFilter stores a user’s filter as a percolator query in Elasticsearch
func SaveFilter(ctx context.Context, es *elasticsearch.Client, userID string, filter map[string]interface{}) error {
	filterDoc := Filter{
		UserID: userID,
		Query:  filter,
	}

	data, err := json.Marshal(filterDoc)
	if err != nil {
		return fmt.Errorf("failed to marshal filter: %w", err)
	}

	res, err := es.Index(
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
func MatchApartment(ctx context.Context, es *elasticsearch.Client, apartment map[string]interface{}) ([]string, error) {
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

	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex("filters"),
		es.Search.WithBody(strings.NewReader(string(data))),
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
