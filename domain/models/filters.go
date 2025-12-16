package models

type ApartmentFilter struct {
	MinPrice        float64
	MaxPrice        float64
	RoomNumbers     []int
	BedroomNumbers  []int
	BathroomNumbers []int
	City            string
	District        string
}
