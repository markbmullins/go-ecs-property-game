package models

import "time"

type Upgrade struct {
	Name           string
	ID             string
	Level          int
	Cost           float64
	RentIncrease   float64
	DaysToComplete int
	PurchaseDate   time.Time
	Prerequisite   *Upgrade
	Applied        bool
}
