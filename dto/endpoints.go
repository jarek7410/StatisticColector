package dto

import "time"

type PostStatDto struct {
	Name   string  `json:"name,omitempty"`
	NameID uint    `json:"nameID,omitempty"`
	Value  float64 `json:"value"`
	Type   string  `json:"type"`
}

type GetStatsDto struct {
	Name   string    `json:"name,omitempty"`
	NameID uint      `json:"nameID,omitempty"`
	Offset int       `json:"offset"`
	Limit  int       `json:"limit"`
	Stats  []StatDto `json:"stats,omitempty"`
}

type StatDto struct {
	Value     float64   `json:"value"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

type NameDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
