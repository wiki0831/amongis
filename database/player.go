package database

import (
	"amongis/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetPlayerInfo(name string) (*model.Player, error) {
	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM latest_player_data 
		WHERE name = '%s'`,
		name,
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("player doesnt exist")
	}

	var players []model.Player

	for rows.Next() {
		var player model.Player
		err := rows.Scan(&player.Id, &player.Name, &player.Role, &player.Room, &player.Status, &player.CreatedAt, &player.Location)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	if len(players) != 1 {
		return nil, fmt.Errorf("db failed at checkout")
	}

	return &players[0], nil
}

func GetAllPlayerInfo() ([]model.Player, error) {
	ctx := context.Background()
	queryString := `SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM latest_player_data`

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("player doesnt exist")
	}

	var players []model.Player

	for rows.Next() {
		var player model.Player
		err := rows.Scan(&player.Id, &player.Name, &player.Role, &player.Room, &player.Status, &player.CreatedAt, &player.Location)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

func CreatePlayerTelemetry(user model.Player) error {
	query := `
			INSERT INTO player 
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

func GetPlayersNearby(currentPlayer model.Player) ([]model.Player, error) {

	ctx := context.Background()
	queryString := fmt.Sprintf(
		`SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM latest_player_data 
		where  name != '%s'
		and status = 'alive'`,
		currentPlayer.Name,
	)

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("player doesnt exist")
	}

	var players []model.Player

	for rows.Next() {
		var player model.Player
		err := rows.Scan(&player.Id, &player.Name, &player.Role, &player.Room, &player.Status, &player.CreatedAt, &player.Location)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}
