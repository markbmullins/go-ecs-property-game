package models

import "time"

type PropertyType string
type PropertySubtype string

const (
	Residential PropertyType = "Residential"
	Commercial  PropertyType = "Commercial"
)

// Subtypes for residential properties
const (
	SingleFamily PropertySubtype = "SingleFamily"
	Townhome     PropertySubtype = "Townhome"
	Multifamily  PropertySubtype = "Multifamily"
	Apartment    PropertySubtype = "Apartment"
	Condo        PropertySubtype = "Condo"
)

// Subtypes for commercial properties
const (
	OfficeSpace        PropertySubtype = "OfficeSpace"
	RetailStore        PropertySubtype = "RetailStore"
	Warehouse          PropertySubtype = "Warehouse"
	Restaurant         PropertySubtype = "Restaurant"
	Hotel              PropertySubtype = "Hotel"
	Mall               PropertySubtype = "Mall"
	Industrial         PropertySubtype = "Industrial"
	Clinic             PropertySubtype = "Clinic"
	DataCenter         PropertySubtype = "DataCenter"
	Bar                PropertySubtype = "Bar"
	NightClub          PropertySubtype = "NightClub"
	Museum             PropertySubtype = "Museum"
	Amusement          PropertySubtype = "Amusement"
	Factory            PropertySubtype = "Factory"
	DistributionCenter PropertySubtype = "DistributionCenter"
)

type Property struct {
	Name                         string
	Type                         PropertyType
	Subtype                      PropertySubtype
	BaseRent                     float64
	RentBoost                    float64
	Owned                        bool
	UpgradeLevel                 int
	Upgrades                     []Upgrade
	UpgradePaths                 map[string][]Upgrade
	Price                        float64
	PlayerID                     int // ID of the owning player
	OccupancyRate                float64
	TenantSatisfaction           int
	PurchaseDate                 time.Time
	ProrateRent                  bool
	NeighborhoodID               int
	UgradedNeighborhoodRentBoost float64
}
