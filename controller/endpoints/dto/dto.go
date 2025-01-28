package dto

type PostStatDto struct {
	Name   string  `json:"name,omitempty"`
	NameID uint    `json:"nameID,omitempty"`
	Value  float64 `json:"value"`
	Type   string  `json:"type"`
}
