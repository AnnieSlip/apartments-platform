package postgres

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type FilterPostgresRepo struct{}

func NewFilterPostgresRepo() *FilterPostgresRepo {
	return &FilterPostgresRepo{}
}

func (r *FilterPostgresRepo) SaveFilter(ctx context.Context, userID int, f models.ApartmentFilter) error {
	_, err := DB.Exec(ctx, `INSERT INTO user_filters(user_id, district, room_numbers, bedroom_numbers, bathroom_numbers, price_per_month) 
		VALUES($1,$2,$3,$4,$5,$6) 
		ON CONFLICT (user_id) DO UPDATE 
		SET district=$2, room_numbers=$3, bedroom_numbers=$4, bathroom_numbers=$5, price_per_month=$6`,
		userID, f.District, f.RoomNumbers, f.BedroomNumbers, f.BathroomNumbers, f.MaxPrice,
	)
	return err
}

func (r *FilterPostgresRepo) GetFiltersByUser(ctx context.Context, userID int) ([]models.ApartmentFilter, error) {
	rows, err := DB.Query(ctx, `
        SELECT district, room_numbers, bedroom_numbers, bathroom_numbers, price_per_month
        FROM user_filters
        WHERE user_id = $1
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var filters []models.ApartmentFilter
	for rows.Next() {
		var f models.ApartmentFilter
		if err := rows.Scan(&f.District, &f.RoomNumbers, &f.BedroomNumbers, &f.BathroomNumbers, &f.MaxPrice); err != nil {
			return nil, err
		}
		filters = append(filters, f)
	}

	return filters, nil
}

func (r *FilterPostgresRepo) GetAllFilters(ctx context.Context) ([]models.UserFilter, error) {
	rows, err := DB.Query(ctx, `
		SELECT user_id, district, room_numbers, bedroom_numbers, bathroom_numbers, price_per_month
		FROM user_filters
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.UserFilter

	for rows.Next() {
		var uf models.UserFilter
		err := rows.Scan(
			&uf.UserID,
			&uf.Filter.District,
			&uf.Filter.RoomNumbers,
			&uf.Filter.BedroomNumbers,
			&uf.Filter.BathroomNumbers,
			&uf.Filter.MaxPrice,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, uf)
	}

	return res, nil
}
