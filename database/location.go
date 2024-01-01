package database

import (
	"amongis/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateLocation(location model.Location) error {
	query := `
			INSERT INTO location 
				(name, role, room, status, location, created_at) 
			VALUES 
				(@locationName, @locationRole, @locationRoom, @locationStatus, ST_GeomFromWKB(@locationLocation,4326), @locationTimestamp)`

	args := pgx.NamedArgs{
		"locationName":      location.Name,
		"locationRole":      location.Role,
		"locationRoom":      location.Room,
		"locationStatus":    location.Status,
		"locationLocation":  location.Location.AsBinary(),
		"locationTimestamp": location.CreatedAt,
	}

	_, err := DB.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func GetLocationsNearby(currentPlayer model.Player) ([]model.Location, error) {

	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM location 
		WHERE ST_DWithin(location,ST_GeomFromText('%s', 4326), 0.001)`,
		currentPlayer.Location.AsText(),
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("location doesnt exist")
	}

	var locations []model.Location

	for rows.Next() {
		var nextlocation model.Location
		err := rows.Scan(&nextlocation.Id, &nextlocation.Name, &nextlocation.Role, &nextlocation.Room, &nextlocation.Status, &nextlocation.CreatedAt, &nextlocation.Location)
		if err != nil {
			return nil, err
		}

		locations = append(locations, nextlocation)
	}

	return locations, nil
}
