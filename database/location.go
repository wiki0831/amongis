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

func UpdateLocationStatus(location model.Location, status string) error {
	query := fmt.Sprintf(`UPDATE Location
	SET Status='%s'WHERE name='%s' AND room = '%s' AND role = '%s'`,
		status, location.Name, location.Room, location.Role)

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to update Location: %w", err)
	}
	return nil
}

func GetLocationInfo(name string, room string) ([]model.Location, error) {
	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM location
		WHERE name = '%s' AND room = '%s'`,
		name, room,
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("location doesnt exist")
	}

	var locations []model.Location

	for rows.Next() {
		var location model.Location
		err := rows.Scan(&location.Id, &location.Name, &location.Role, &location.Room, &location.Status, &location.CreatedAt, &location.Location)
		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	if len(locations) != 1 {
		return nil, fmt.Errorf("db failed at checkout")
	}

	return locations, nil
}

func GetLocationsNearby(currentPlayer model.Player) ([]model.Location, error) {

	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM location 
		WHERE ST_DWithin(location,ST_GeomFromText('%s', 4326), 0.001) AND room = '%s'`,
		currentPlayer.Location.AsText(), currentPlayer.Room,
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("location doesnt exist")
	}

	var locations []model.Location

	for rows.Next() {
		var nextPlayer model.Location
		err := rows.Scan(&nextPlayer.Id, &nextPlayer.Name, &nextPlayer.Role, &nextPlayer.Room, &nextPlayer.Status, &nextPlayer.CreatedAt, &nextPlayer.Location)
		if err != nil {
			return nil, err
		}

		locations = append(locations, nextPlayer)
	}

	return locations, nil
}

func GetPlayersNearLocaiton(currentLocation model.Location) ([]model.Player, error) {

	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM latest_player_data 
		WHERE ST_DWithin(location,ST_GeomFromText('%s', 4326), 0.001) AND room = '%s'`,
		currentLocation.Location.AsText(), currentLocation.Room,
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("location doesnt exist")
	}

	var Players []model.Player

	for rows.Next() {
		var nextPlayer model.Player
		err := rows.Scan(&nextPlayer.Id, &nextPlayer.Name, &nextPlayer.Role, &nextPlayer.Room, &nextPlayer.Status, &nextPlayer.CreatedAt, &nextPlayer.Location)
		if err != nil {
			return nil, err
		}

		Players = append(Players, nextPlayer)
	}

	return Players, nil
}

func CheckExit(room string) bool {
	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM location
		WHERE role = 'mission' AND room = '%s'`,
		room,
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return false
	}

	for rows.Next() {
		var location model.Location
		err := rows.Scan(&location.Id, &location.Name, &location.Role, &location.Room, &location.Status, &location.CreatedAt, &location.Location)
		if err != nil {
			return false
		}
		if location.Status != "fixed" {
			return false
		}
	}
	err = OpenExit(room)
	return err == nil

}

func OpenExit(room string) error {
	query := fmt.Sprintf(`
	UPDATE Location
	SET Status='alive' 
	WHERE room = '%s' AND role = 'exit'`,
		room)

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to update Location: %w", err)
	}
	return nil
}
