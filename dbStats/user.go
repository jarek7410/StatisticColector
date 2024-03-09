package dbStats

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Stats    []Stat `json:"stat,omitempty"`
}

func (User) TableName() string {
	return "Users"
}
