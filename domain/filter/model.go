package filter

type ApartmentFilter struct {
	MinPrice, MaxPrice float64
	RoomNumbers        []int
	BedroomNumbers     []int
	BathroomNumbers    []int
	City               string
	District           string
}
