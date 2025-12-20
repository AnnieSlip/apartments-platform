package postgres

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type ApartmentPostgresRepo struct{}

func NewApartmentPostgresRepo() *ApartmentPostgresRepo {
	return &ApartmentPostgresRepo{}
}

func (r *ApartmentPostgresRepo) GetAll(ctx context.Context) ([]models.Apartment, error) {
	rows, err := DB.Query(ctx, "SELECT id, title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, district, city FROM apartments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Apartment
	for rows.Next() {
		var a models.Apartment
		if err := rows.Scan(&a.ID, &a.Title, &a.PricePerMonth, &a.RoomNumbers, &a.BedroomNumbers, &a.BathroomNumbers, &a.District, &a.City); err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}

func (r *ApartmentPostgresRepo) Create(ctx context.Context, a models.Apartment) (models.Apartment, error) {
	var newA models.Apartment
	err := DB.QueryRow(ctx,
		`INSERT INTO apartments(title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, district, city)
		 VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, district, city`,
		a.Title, a.PricePerMonth, a.RoomNumbers, a.BedroomNumbers, a.BathroomNumbers, a.District, a.City,
	).Scan(&newA.ID, &newA.Title, &newA.PricePerMonth, &newA.RoomNumbers, &newA.BedroomNumbers, &newA.BathroomNumbers, &newA.District, &newA.City)

	if err != nil {
		return models.Apartment{}, err
	}
	return newA, nil
}

func (r *ApartmentPostgresRepo) GetApartmentsByFilter(ctx context.Context, f models.ApartmentFilter) ([]apartment.Apartment, error) {
	rows, err := DB.Query(ctx, `
		SELECT id, title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, city, district
		FROM apartments
		WHERE price_per_month >= $1 AND price_per_month <= $2
		  AND room_numbers = ANY($3)
		  AND bedroom_numbers = ANY($4)
		  AND bathroom_numbers = ANY($5)
		  AND city = $6
		  AND (district = $7 OR $7 IS NULL)
	`, f.MinPrice, f.MaxPrice, f.RoomNumbers, f.BedroomNumbers, f.BathroomNumbers, f.City, f.District)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []apartment.Apartment
	for rows.Next() {
		var a apartment.Apartment
		if err := rows.Scan(&a.ID, &a.Title, &a.PricePerMonth, &a.RoomNumbers, &a.BedroomNumbers, &a.BathroomNumbers, &a.City, &a.District); err != nil {
			return nil, err
		}
		res = append(res, a)
	}

	return res, nil
}
