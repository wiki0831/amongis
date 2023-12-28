package model

import (
	"github.com/google/uuid"
	"github.com/peterstace/simplefeatures/geom"
)

type PlayerModel struct {
	Name     string     `json:"name"`
	UID      uuid.UUID  `json:"uuid"`
	Role     string     `json:"role"`
	Status   string     `json:"status"`
	Location geom.Point `json:"location"`
}

