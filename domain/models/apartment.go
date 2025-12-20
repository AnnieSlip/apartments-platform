package models

type Apartment struct {
	ID              int     `json:"id"`
	Title           string  `json:"title"`
	PricePerMonth   float64 `json:"price_per_month"`
	RoomNumbers     int     `json:"room_numbers"`
	BedroomNumbers  int     `json:"bedroom_numbers"`
	BathroomNumbers int     `json:"bathroom_numbers"`
	District        string  `json:"district"`
	City            string  `json:"city"`
}
