package database

import (
	"amongis/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateLocation(location model.LocationModel) error {
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


