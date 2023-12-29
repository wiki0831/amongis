package logic

import (
	"amongis/database"
	"amongis/model"
	"fmt"
)

func GetActionItem(p *model.PlayerModel) ([]model.ActionModel, error) {
	// init action array
	var actions []model.ActionModel
	//check my current role and room

	//get player around player p
	playerlist, err := database.GetPlayersNearby(*p)
	if err != nil {
		return nil, fmt.Errorf("error reading database")
	}
	for _, value := range playerlist {
		action := model.ActionModel{
			ActionType: "kill",
			Target:     value.Name,
		}
		actions = append(actions, action)
	}
	//get location around player p
	//return action
	return actions, nil
}

