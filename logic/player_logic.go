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

	for _, value := range playerlist {
		action := model.Action{
			ActionStatus: value.Status,
			ActionType:   p.Role,
			Target:       value.Name,
			TargetType:   value.Role,
		}
		actions = append(actions, action)
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
	for _, value := range locationlist {
		action := model.Action{
			ActionStatus: value.Status,
			ActionType:   value.Role,
			Target:       value.Name,
			TargetType:   "location",
		}
		actions = append(actions, action)
	}
	//get location around player p
	//return action
	return actions, nil
}

// use map to prevent invalid entry
// actionStausMap := map[string]string{
// 	"kill": "dead",
// 	"key2": "value2",
// 	"key3": "value3",
// }

func PerformActionItem(p *model.Player, avaibleActions []model.Action) {
	for _, value := range avaibleActions {
		if reflect.DeepEqual(p.Action, value) {
			targetPlayer, _ := database.GetPlayerInfo(p.Action.Target)
			targetPlayer.Status = p.Action.ActionType
			targetPlayer.CreatedAt = time.Now()
			database.CreatePlayerTelemetry(*targetPlayer)
		}
	}
}
