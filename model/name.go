package model

import (
	"StatisticColector/database"
	"gorm.io/gorm"
)

type Name struct {
	gorm.Model
	Name  string
	Stats []Stat
}

func (N *Name) GetByName() *Name {
	database.Re.DB.
		Where("name = ?", N.Name).Limit(1).First(&N)

	return N
}
func (N *Name) Save() {
	database.Re.DB.Save(&N)
}
func (N *Name) GetStats() []Stat {
	database.Re.DB.Preload("Stats").Find(&N.Stats)
	return N.Stats
}
