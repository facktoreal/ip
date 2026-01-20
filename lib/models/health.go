package models

type Health struct {
	Healthy  bool   `json:"healthy"`
	Uptime   string `json:"uptime"`
	Version  string `json:"version"`
	Instance string `json:"instance"`
}
