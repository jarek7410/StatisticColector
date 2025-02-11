package dto

type HalfCheckDto struct {
	Version     string `json:"version"`
	VersionName string `json:"version_name"`
	Database    string `json:"database"`
}
