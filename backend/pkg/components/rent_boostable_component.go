package components

type RentBoostable struct {
	GroupID             int     // The group id the rent boost applies to
	ThresholdPercentage float64 // Percentage of upgraded properties in a group required to activate the rent boost
	BoostPercentage     float64 // Boost percentage applied to rents
}
