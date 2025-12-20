package elasticsearch

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type EsRepo struct {
	Client *elasticsearch.Client
}

func NewEsRepo(client *elasticsearch.Client) *EsRepo {
	return &EsRepo{Client: client}
}

func CreateIndices(es *elasticsearch.Client) error {
	indices := map[string]string{
		"apartments": ApartmentsMapping,
		"filters":    FiltersMapping,
	}

	for name, mapping := range indices {
		exists, err := es.Indices.Exists([]string{name})
		if err != nil {
			return err
		}
		if exists.StatusCode == 404 {
			_, err = es.Indices.Create(name, es.Indices.Create.WithBody(strings.NewReader(mapping)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
