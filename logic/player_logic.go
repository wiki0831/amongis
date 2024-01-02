package logic

import (
	"amongis/database"
	"amongis/model"
	"fmt"
	"reflect"
	"time"
)

func GetAvailableAction(p *model.Player) ([]model.Action, error) {
	playerActions, err := GetPlayerActions(p)
	if err != nil {
		return nil, fmt.Errorf("unable to compute user actions")
	}

	locationActions, err := GetLocActions(p)
	if err != nil {
		return nil, fmt.Errorf("unable to compute user actions")
	}

	return append(playerActions, locationActions...), nil
}

func GetPlayerActions(p *model.Player) ([]model.Action, error) {
	// init action array
	var actions []model.Action
	//check my current role and room

	//get player around player p
	playerlist, err := database.GetPlayersNearby(*p)
	if err != nil {
		return nil, fmt.Errorf("error reading database")
	}

	if p.Role == "killer" {
		for _, value := range playerlist {
			action := model.Action{
				ActionStatus: value.Status,
				ActionType:   "kill",
				Target:       value.Name,
				TargetType:   value.Role,
			}
			if value.Role == "player" {
				actions = append(actions, action)
			}
		}
	} else {
		for _, value := range playerlist {
			action := model.Action{
				ActionStatus: value.Status,
				ActionType:   p.Role,
				Target:       value.Name,
				TargetType:   value.Role,
			}
			if value.Role == "player" {
				actions = append(actions, action)
			}
		}
	}

	//get location around player p
	//return action
	return actions, nil
}

func GetLocActions(p *model.Player) ([]model.Action, error) {
	// init action array
	var actions []model.Action
	//check my current role and room

	//get location around player p
	locationlist, err := database.GetLocationsNearby(*p)
	if err != nil {
		return nil, fmt.Errorf("error reading database")
	}
	if p.Role == "killer" {
		for _, value := range locationlist {
			action := model.Action{
				ActionStatus: value.Status,
				ActionType:   "stop",
				Target:       value.Name,
				TargetType:   "location",
			}
			if value.Role == "mission" {
				actions = append(actions, action)
			}
		}
	} else {
		for _, value := range locationlist {
				action := model.Action{
					ActionStatus: value.Status,
					Target:       value.Name,
					TargetType:   "location"}
			if value.Role == "mission" && value.Status == "dead" {
				action.ActionType = "fix"
			} else {
				action.ActionType = value.Role
			}

			actions = append(actions, action)
		}
	}

	//get location around player p
	//return action
	return actions, nil
}

func PerformActionItem(p *model.Player, avaibleActions []model.Action) {
	actionStausMap := map[string]string{
		"kill":    "dead",
		"mission": "fixed",
		"exit":    "excape",
		"respawn": "alive",
		"stop":    "dead",
		"fix":     "alive",
	}
	for _, value := range avaibleActions {
		if p.Action.TargetType == "player" {
			if reflect.DeepEqual(p.Action, value) {
				targetPlayer, _ := database.GetPlayerInfo(p.Action.Target)
				targetPlayer.Status = actionStausMap[p.Action.ActionType]
				targetPlayer.CreatedAt = time.Now()
				database.CreatePlayerTelemetry(*targetPlayer)
			}
		}

		if p.Action.TargetType == "location" {
			switch p.Action.ActionType {
			case "respawn":
				fmt.Println("perform respawn")
				players, _ := database.GetPlayersNearby(*p)
				for _, value := range players {
					if value.Status == "dead" && p.Role == "player" && p.Status == "alive" {
						targetPlayer, _ := database.GetPlayerInfo(value.Name)
						targetPlayer.Status = actionStausMap[p.Action.ActionType]
						targetPlayer.CreatedAt = time.Now()
						database.CreatePlayerTelemetry(*targetPlayer)
					}
				}
			case "mission":
				fmt.Println("perform mission")
				locations, locErr1 := database.GetLocationInfo(p.Action.Target, p.Room)
				if locErr1 == nil {
					for _, currentLocation := range locations {
						if currentLocation.Role == "mission" && currentLocation.Status == "alive" && p.Role == "player" && p.Status == "alive" {
							database.UpdateLocationStatus(currentLocation, actionStausMap[p.Action.ActionType])
						}
					}
				}
				database.CheckExit(p.Room)
			case "exit":
				fmt.Println("perform exit")
				p.Status = actionStausMap[p.Action.ActionType]
				p.CreatedAt = time.Now()
				database.CreatePlayerTelemetry(*p)
			case "stop":
				fmt.Println("perform stop")
				locations, locErr1 := database.GetLocationInfo(p.Action.Target, p.Room)
				if locErr1 == nil {
					for _, currentLocation := range locations {
						if currentLocation.Role == "mission" && currentLocation.Status == "alive" && p.Role == "killer" {
							database.UpdateLocationStatus(currentLocation, actionStausMap[p.Action.ActionType])
						}
					}
				}
			case "fix":
				fmt.Println("perform fix")
				locations, locErr1 := database.GetLocationInfo(p.Action.Target, p.Room)
				if locErr1 == nil {
					for _, currentLocation := range locations {
						if currentLocation.Role == "mission" && currentLocation.Status == "dead" && p.Role == "player" && p.Status == "alive" {
							database.UpdateLocationStatus(currentLocation, actionStausMap[p.Action.ActionType])
						}
					}
				}
			default:
				fmt.Println("perform default")
			}

		}

	}
}
