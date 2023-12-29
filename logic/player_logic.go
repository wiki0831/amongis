package logic

import (
	"amongis/database"
	"amongis/model"
	"fmt"
	"reflect"
	"time"
)

func GetActionItem(p *model.Player) ([]*model.PlayerAction, error) {
	// init action array
	var actions []*model.PlayerAction
	//check my current role and room

	//get player around player p
	playerlist, err := database.GetPlayersNearby(*p)
	if err != nil {
		return nil, fmt.Errorf("error reading database")
	}
	for _, value := range playerlist {
		action := model.PlayerAction{
			ActionType: "kill",
			Target:     value.Name,
		}
		actions = append(actions, &action)
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

func PerformActionItem(p *model.Player, avaibleActions []*model.PlayerAction) {
	for _, value := range avaibleActions {
		if reflect.DeepEqual(p.Action, *value) {
			targetPlayer, _ := database.GetPlayerInfo(p.Action.Target)
			targetPlayer.Status = p.Action.ActionType
			targetPlayer.CreatedAt = time.Now()
			database.CreatePlayerTelemetry(*targetPlayer)
		}
	}
}
