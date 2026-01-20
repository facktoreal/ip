package models

type Health struct {
	Healthy bool   `json:"healthy"`
	Uptime  string `json:"uptime"`
}
