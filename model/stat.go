package model

import (
	"StatisticColector/database"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	Value  float64 `json:"value"`
	Type   string  `json:"type" gorm:"default:default"`
	NameID uint    `json:"nameID,omitempty"`
}

func (S *Stat) Save() *Stat {
	database.Re.DB.Save(&S)
	return S
}
