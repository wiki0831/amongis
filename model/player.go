package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/peterstace/simplefeatures/geom"
)

type Player struct {
	Name      string       `json:"name" required:"true"`
	Id        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	Role      string       `json:"role"`
	Room      string       `json:"room"`
	Status    string       `json:"status"`
	Location  geom.Point   `json:"location" required:"true"`
	Action    PlayerAction `json:"action" required:"true"`
}

type PlayerAction struct {
	ActionType string `json:"action_type" required:"true"`
	Target     string `json:"target" required:"true"`
}

func (p *Player) Validate() error {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	if p.Name == "" {
		return fmt.Errorf("empty user name")
	}
	if p.Role != "killer" && p.Role != "player" {
		return fmt.Errorf("invalid role")
	}
	//Todo Check current status in DB
	return nil
}
