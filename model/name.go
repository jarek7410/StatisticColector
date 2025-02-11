package model

import (
	"StatisticColector/database"
	"StatisticColector/dto"
	"gorm.io/gorm"
)

type Name struct {
	gorm.Model
	Name  string
	Stats []Stat `gorm:"foreignKey:NameID" json:"stats,omitempty"`
}

func (N *Name) GetByName() *Name {
	database.Re.DB.
		Where("name = ?", N.Name).Limit(1).First(&N)

	return N
}
func (N *Name) Save() {
	database.Re.DB.Save(&N)
}
func (N *Name) Get() {
	database.Re.DB.Find(&N)
}
func (N *Name) GetStats(limit, offset uint) []Stat {
	database.Re.DB.Preload("Stats", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(int(limit)).Offset(int(offset))
	}).First(&N)
	return N.Stats
}
func GetAllNames(limit, offset uint) []dto.NameDto {
	var N []dto.NameDto
	database.Re.DB.Model(&Name{}).Limit(int(limit)).Offset(int(offset)).Find(&N)
	return N
}
