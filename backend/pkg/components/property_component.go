package components

import "time"

type PropertyType string
type PropertySubtype string

const (
	Residential PropertyType = "Residential"
	Commercial  PropertyType = "Commercial"
)

// Subtypes for residential properties
const (
	Apartment    PropertySubtype = "Apartment"
	Condo        PropertySubtype = "Condo"
	Duplex       PropertySubtype = "Duplex"
	Multifamily  PropertySubtype = "Multifamily"
	Penthouse    PropertySubtype = "Penthouse"
	SingleFamily PropertySubtype = "SingleFamily"
	Townhome     PropertySubtype = "Townhome"
)

// Subtypes for commercial properties
const (
	Amusement          PropertySubtype = "Amusement"
	ArtGallery         PropertySubtype = "ArtGallery"
	Bakery             PropertySubtype = "Bakery"
	Bar                PropertySubtype = "Bar"
	Cafe               PropertySubtype = "Cafe"
	Clinic             PropertySubtype = "Clinic"
	DataCenter         PropertySubtype = "DataCenter"
	DistributionCenter PropertySubtype = "DistributionCenter"
	Factory            PropertySubtype = "Factory"
	Gym                PropertySubtype = "Gym"
	Hotel              PropertySubtype = "Hotel"
	Industrial         PropertySubtype = "Industrial"
	Mall               PropertySubtype = "Mall"
	Museum             PropertySubtype = "Museum"
	NightClub          PropertySubtype = "NightClub"
	OfficeSpace        PropertySubtype = "OfficeSpace"
	Restaurant         PropertySubtype = "Restaurant"
	RetailStore        PropertySubtype = "RetailStore"
	Salon              PropertySubtype = "Salon"
	Warehouse          PropertySubtype = "Warehouse"
)

type Property struct {
  Name                         string
	Type                         PropertyType
	ID                           int
	Subtype                      PropertySubtype
	BaseRent                     float64
	RentBoost                    float64 // Any applied rent boosts e.g. the neighborhood upgrade rent boost
	Owned                        bool
	Upgrades                     []Upgrade
	UpgradePaths                 map[string][]Upgrade
	Price                        float64
	PlayerID                     int // ID of the owning player
	OccupancyRate                float64
	TenantSatisfaction           int
	PurchaseDate                 time.Time
	NeighborhoodID               int
	Description                  string
	Address                      string
}
