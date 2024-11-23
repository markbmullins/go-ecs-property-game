package models

import "time"

type GameTime struct {
	CurrentDate     time.Time
	IsPaused        bool
	SpeedMultiplier float64 // 1.0 = normal, 2.0 = fast, etc.
	NewMonth        bool
	LastUpdated     time.Time
}
