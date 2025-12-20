package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

func (r *EsRepo) IndexApartment(ctx context.Context, a models.Apartment) error {
	doc := map[string]interface{}{
		"price_per_month":  a.PricePerMonth,
		"room_numbers":     a.RoomNumbers,
		"bedroom_numbers":  a.BedroomNumbers,
		"bathroom_numbers": a.BathroomNumbers,
		"city":             a.City,
		"district":         a.District,
		"created_at":       time.Now().UTC(),
	}

	data, _ := json.Marshal(doc)

	res, err := r.Client.Index(
		"apartments",
		strings.NewReader(string(data)),
		r.Client.Index.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("failed to index apartment: %s", res.String())
	}

	return nil
}

func (r *EsRepo) PercolateApartment(ctx context.Context, a models.Apartment) ([]string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"percolate": map[string]interface{}{
				"field": "query",
				"document": map[string]interface{}{
					"price_per_month":  a.PricePerMonth,
					"room_numbers":     a.RoomNumbers,
					"bedroom_numbers":  a.BedroomNumbers,
					"bathroom_numbers": a.BathroomNumbers,
					"city":             a.City,
					"district":         a.District,
				},
			},
		},
	}

	data, _ := json.Marshal(query)

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
				Source struct {
					UserID string `json:"user_id"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	var users []string
	for _, hit := range result.Hits.Hits {
		users = append(users, hit.Source.UserID)
	}

	return users, nil
}
