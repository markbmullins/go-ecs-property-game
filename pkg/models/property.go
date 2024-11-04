package models

type Property struct {
    Position    Position
    Level       int
    BaseIncome  float64
    UpgradeCost float64
    Owned       bool
    Name        string
    Address     string
    Type        string // "Residential" or "Commercial"
    Subtype     string // e.g., "Single Family", "Multifamily"
}
