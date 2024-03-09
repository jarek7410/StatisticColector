package dbStats

import "gorm.io/gorm"

type Stat struct {
	gorm.Model
	Name   string  `json:"name" gorm:"default:default"`
	Value  float64 `json:"value"`
	Type   string  `json:"type"gorm:"default:default"`
	UserID uint    `json:"userID,omitempty"`
}

func (Stat) TableName() string {
	return "Stats"
}
