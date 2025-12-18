package elasticsearch

// its a blueprint for Elasticsearch index.
// This tells Elasticsearch:
// What fields documents will have/what type each field is.

// amenities is a field in apartment documents (like ["pool", "parking", "gym"]).
// than I can do exact matches or filtering, like amenities contains "pool" or amenities contains "parking".

// "type": "keyword" means: this field is stored as-is and is not analyzed for full-text search -keyword → for exact matching

// index names
const (
	ApartmentsIndex = "apartments"
	FiltersIndex    = "filters"
)

// ApartmentsMapping is the mapping for the apartments index
const ApartmentsMapping = `
{
  "mappings": {
    "properties": {
      "id":              { "type": "integer" },
      "title":           { 
        "type": "text",
        "fields": {
          "keyword": { "type": "keyword" }
        }
      },
      "price_per_month": { "type": "float" },
      "room_numbers":    { "type": "integer" },
      "bedroom_numbers": { "type": "integer" },
      "bathroom_numbers":{ "type": "integer" },
      "district":        { "type": "keyword" },
      "city":            { "type": "keyword" },
      "amenities":       { "type": "keyword" },
      "created_at":      { "type": "date" }
    }
  }
}`

// FiltersMapping is the mapping for the filters percolator index

// stores user queries (percolator queries)
const FiltersMapping = `
{
  "mappings": {
    "properties": {
      "query":      { "type": "percolator" },
      "user_id":    { "type": "keyword" },
      "created_at": { "type": "date" }
    }
  }
}`
