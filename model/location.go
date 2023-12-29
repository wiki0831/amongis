package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/peterstace/simplefeatures/geom"
)

type Location struct {
	Name      string     `json:"name" required:"true"`
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Role      string     `json:"role"`
	Room      string     `json:"room"`
	Status    string     `json:"status"`
	Location  geom.Point `json:"location" required:"true"`
}

func (l *Location) Validate() error {
	if l.CreatedAt.IsZero() {
		l.CreatedAt = time.Now()
	}
	if l.Name == "" {
		return fmt.Errorf("empty user name")
	}
	if l.Role != "exit" && l.Role != "mission" && l.Role != "respawn" {
		return fmt.Errorf("invalid role")
	}

	return nil
}
