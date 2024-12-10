package components

import "time"

type Rentable struct {
	BaseRent               float64
	RentBoost              float64 // Any applied rent boosts e.g. the neighborhood upgrade rent boost
	LastRentCollectionDate time.Time
}
