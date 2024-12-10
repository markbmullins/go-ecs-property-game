package components

type Maintainable struct {
	Condition      float64  // 0-100, where 100 is perfect
	MaintenanceDue float64  // Accumulated maintenance cost not yet addressed
	DamageEvents   []string // e.g., "HVAC_Break", "Roof_Leak"
}
