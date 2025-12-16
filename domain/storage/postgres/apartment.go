package postgres

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
)

type ApartmentPostgresRepo struct{}

func NewApartmentPostgresRepo() *ApartmentPostgresRepo {
	return &ApartmentPostgresRepo{}
}

func (r *ApartmentPostgresRepo) GetAll(ctx context.Context) ([]apartment.Apartment, error) {
	rows, err := DB.Query(ctx, "SELECT id, title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, address, city FROM apartments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []apartment.Apartment
	for rows.Next() {
		var a apartment.Apartment
		if err := rows.Scan(&a.ID, &a.Title, &a.PricePerMonth, &a.RoomNumbers, &a.BedroomNumbers, &a.BathroomNumbers, &a.Address, &a.City); err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}

func (r *ApartmentPostgresRepo) Create(ctx context.Context, a apartment.Apartment) (apartment.Apartment, error) {
	var newA apartment.Apartment
	err := DB.QueryRow(ctx,
		`INSERT INTO apartments(title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, address, city)
		 VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, address, city`,
		a.Title, a.PricePerMonth, a.RoomNumbers, a.BedroomNumbers, a.BathroomNumbers, a.Address, a.City,
	).Scan(&newA.ID, &newA.Title, &newA.PricePerMonth, &newA.RoomNumbers, &newA.BedroomNumbers, &newA.BathroomNumbers, &newA.Address, &newA.City)

	if err != nil {
		return apartment.Apartment{}, err
	}
	return newA, nil
}
