package models

import "time"

type Stats struct {
	Uptime time.Time `json:"uptime"`
}
