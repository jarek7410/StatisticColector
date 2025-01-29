package model

import (
	"StatisticColector/database"
	"time"
)

type Stat struct {
	CreatedAt time.Time
	Value     float64 `json:"value"`
	Type      string  `json:"type" gorm:"default:default"`
	NameID    uint    `json:"nameID,omitempty"`
}

func (S *Stat) Save() *Stat {
	database.Re.DB.Create(&S)
	return S
}
