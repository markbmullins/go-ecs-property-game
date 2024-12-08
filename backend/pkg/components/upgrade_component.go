package components

import "time"

type Upgrade struct {
	Name           string
	ID             string
	Level          int // TODO: Do I need this field?
	Cost           float64
	RentIncrease   float64 // The amount of rent increase the upgrade provides to the property
	DaysToComplete int
	PurchaseDate   time.Time
	Prerequisite   *Upgrade
	Applied        bool // Tracks if upgrade has been applied to property
}
