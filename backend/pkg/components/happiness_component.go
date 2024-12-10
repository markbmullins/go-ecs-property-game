package components

// This can be attached to tenants, customers, or any entity affected by conditions.
type Happiness struct {
	Value float64 // 0-100 scale
	// Could add other fields like a list of modifiers or reason codes
}
