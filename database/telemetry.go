package database

import (
	"amongis/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateTelemetry(user model.PlayerModel) error {
	query := `
			INSERT INTO telemetry 
				(name, role, room, status, location, created_at) 
			VALUES 
				(@userName, @userRole, @userRoom, @userStatus, ST_GeomFromWKB(@userLocation,4326), @userTimestamp)`

	args := pgx.NamedArgs{
		"userName":      user.Name,
		"userRole":      user.Role,
		"userRoom":      user.Room,
		"userStatus":    user.Status,
		"userLocation":  user.Location.AsBinary(),
		"userTimestamp": user.CreatedAt,
	}

	_, err := DB.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
