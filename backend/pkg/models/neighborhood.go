package models

type Neighborhood struct {
	ID                   int
	Name                 string
	PropertyIDs          []int // List of property IDs in the neighborhood
	AveragePropertyValue float64
	RentBoostThreshold   float64 // Percentage of properties that need to be upgraded
	RentBoostAmount      float64 // Boost percentage applied to rents
}
