package models

type ApartmentFilter struct {
	MinPrice        float64 `json:"min_price"`
	MaxPrice        float64 `json:"max_price"`
	RoomNumbers     []int   `json:"room_numbers"`
	BedroomNumbers  []int   `json:"bedroom_numbers"`
	BathroomNumbers []int   `json:"bathroom_numbers"`
	City            string  `json:"city"`
	District        string  `json:"district"`
}

type UserFilter struct {
	UserID int
	Filter ApartmentFilter
}
