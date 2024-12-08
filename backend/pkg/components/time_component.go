package components

import "time"

type GameTime struct {
	CurrentDate       time.Time
	IsPaused          bool
	SpeedMultiplier   float64 // 1.0 = normal, 2.0 = fast, etc.
	NewMonth          bool
	LastUpdated       time.Time
	RentCollectionDay int // e.g., 1 for the 1st of the month
}
