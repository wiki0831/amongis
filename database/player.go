package database

import (
	"amongis/model"
	"context"
	"fmt"
)

func GetPlayerInfo(name string) (*model.PlayerModel, error) {
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

	var players []model.PlayerModel

	for rows.Next() {
		var player model.PlayerModel
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

func GetAllPlayerInfo() ([]model.PlayerModel, error) {
	ctx := context.Background()
	queryString := `SELECT 
		id, name, role, room, status, created_at, ST_AsBinary(location) 
		FROM latest_player_data`

	rows, err := DB.Query(ctx, queryString)
	if err != nil {
		return nil, fmt.Errorf("player doesnt exist")
	}

	var players []model.PlayerModel

	for rows.Next() {
		var player model.PlayerModel
		err := rows.Scan(&player.Id, &player.Name, &player.Role, &player.Room, &player.Status, &player.CreatedAt, &player.Location)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}
