// cedar_grove.go
package neighborhoods

import (
	"time"

	"github.com/markbmullins/city-developer/pkg/models"
)
func GetCedarGroveNeighborhood() *models.Neighborhood {
	propertyIDs := make([]int, 0, len(CedarResidential)+len(CedarCommercial))
	for _, prop := range CedarResidential {
		propertyIDs = append(propertyIDs, prop.ID)
	}
	for _, prop := range CedarCommercial {
		propertyIDs = append(propertyIDs, prop.ID)
	}

	return &models.Neighborhood{
		ID:                   4,
		Name:                 "Cedar Grove",
		PropertyIDs:          propertyIDs,
		AveragePropertyValue: 0.0,
		RentBoostThreshold:   50.0, // 50% of properties need to be upgraded
		RentBoostAmount:      10.0, // 10% rent boost
	}
}

// CedarResidential defines the residential properties for Cedar Grove.
var CedarResidential = []models.Property{
	{
		ID:                           61,
		Name:                         "Maplewood Lane House",
		Address:                      "101 Maplewood Lane, Cedar Grove",
		Description:                  "A cozy single-family home with a large backyard and modern amenities.",
		Type:                         models.Residential,
		Subtype:                      models.SingleFamily,
		BaseRent:                     1800.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        300000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           62,
		Name:                         "Sunnybrook Townhome",
		Address:                      "202 Sunnybrook Drive, Cedar Grove",
		Description:                  "A charming townhome with modern finishes and a community garden.",
		Type:                         models.Residential,
		Subtype:                      models.Townhome,
		BaseRent:                     2200.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        350000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           63,
		Name:                         "Oakwood Apartments",
		Address:                      "303 Oakwood Road, Cedar Grove",
		Description:                  "Modern apartments with access to shared recreational facilities and secure parking.",
		Type:                         models.Residential,
		Subtype:                      models.Apartment,
		BaseRent:                     1500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        260000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           64,
		Name:                         "Cedar Grove Condos",
		Address:                      "404 Cedar Boulevard, Cedar Grove",
		Description:                  "Condominiums with private balconies and state-of-the-art home automation systems.",
		Type:                         models.Residential,
		Subtype:                      models.Condo,
		BaseRent:                     2000.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        400000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           65,
		Name:                         "Cedar Grove Estates",
		Address:                      "505 Cedar Lane, Cedar Grove",
		Description:                  "Spacious multifamily residences with modern amenities and landscaped gardens.",
		Type:                         models.Residential,
		Subtype:                      models.Multifamily,
		BaseRent:                     1900.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        380000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           66,
		Name:                         "Cedar Grove Villas",
		Address:                      "606 Villa Avenue, Cedar Grove",
		Description:                  "Luxurious villas featuring private pools and high-end finishes.",
		Type:                         models.Residential,
		Subtype:                      models.SingleFamily,
		BaseRent:                     2200.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        420000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           67,
		Name:                         "Maplewood Condos",
		Address:                      "707 Maplewood Street, Cedar Grove",
		Description:                  "Condominiums with smart home integrations and access to communal lounges.",
		Type:                         models.Residential,
		Subtype:                      models.Condo,
		BaseRent:                     2100.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        400000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           68,
		Name:                         "Sunnybrook Apartments",
		Address:                      "808 Sunnybrook Road, Cedar Grove",
		Description:                  "Apartments with modern designs and access to recreational facilities.",
		Type:                         models.Residential,
		Subtype:                      models.Apartment,
		BaseRent:                     1700.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        340000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           69,
		Name:                         "Cedar Grove Flats",
		Address:                      "909 Cedar Circle, Cedar Grove",
		Description:                  "Modern flats with integrated smart systems and community amenities.",
		Type:                         models.Residential,
		Subtype:                      models.Apartment,
		BaseRent:                     1600.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        330000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           70,
		Name:                         "Oakridge Apartments",
		Address:                      "1001 Oakridge Road, Cedar Grove",
		Description:                  "Spacious apartments with eco-friendly features and access to green spaces.",
		Type:                         models.Residential,
		Subtype:                      models.Apartment,
		BaseRent:                     1500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarResidentialUpgradePaths(),
		Price:                        260000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  true,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
}

// CedarCommercial defines the commercial properties for Cedar Grove.
var CedarCommercial = []models.Property{
	{
		ID:                           71,
		Name:                         "Cozy Corner Café",
		Address:                      "10 Cozy Street, Cedar Grove",
		Description:                  "A friendly neighborhood café serving fresh coffee and pastries.",
		Type:                         models.Commercial,
		Subtype:                      models.Cafe,
		BaseRent:                     4000.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        900000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           72,
		Name:                         "Suburban Shoppe",
		Address:                      "20 Market Lane, Cedar Grove",
		Description:                  "A local retail store offering a variety of household items and essentials.",
		Type:                         models.Commercial,
		Subtype:                      models.RetailStore,
		BaseRent:                     5000.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1100000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           73,
		Name:                         "Cedar Gym",
		Address:                      "30 Fitness Boulevard, Cedar Grove",
		Description:                  "A comprehensive fitness center with modern equipment and personal trainers.",
		Type:                         models.Commercial,
		Subtype:                      models.Gym,
		BaseRent:                     6500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1400000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           74,
		Name:                         "Playtime Arcade",
		Address:                      "40 Fun Avenue, Cedar Grove",
		Description:                  "An arcade offering a variety of games and entertainment for all ages.",
		Type:                         models.Commercial,
		Subtype:                      models.Amusement,
		BaseRent:                     8500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1750000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           75,
		Name:                         "Tech Mart",
		Address:                      "50 Tech Avenue, Cedar Grove",
		Description:                  "A retail store specializing in the latest tech gadgets and accessories.",
		Type:                         models.Commercial,
		Subtype:                      models.RetailStore,
		BaseRent:                     7500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1600000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           76,
		Name:                         "Cedar Pharmacy",
		Address:                      "60 Health Street, Cedar Grove",
		Description:                  "A pharmacy supplying affordable medications and health products.",
		Type:                         models.Commercial,
		Subtype:                      models.Clinic,
		BaseRent:                     8500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1750000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           77,
		Name:                         "Simple Salon",
		Address:                      "70 Beauty Lane, Cedar Grove",
		Description:                  "A basic salon offering essential beauty and grooming services.",
		Type:                         models.Commercial,
		Subtype:                      models.Salon,
		BaseRent:                     3200.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        750000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           78,
		Name:                         "Affordable Arcade",
		Address:                      "80 Fun Boulevard, Cedar Grove",
		Description:                  "An arcade providing a variety of budget-friendly games and entertainment options.",
		Type:                         models.Commercial,
		Subtype:                      models.Amusement,
		BaseRent:                     8500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1750000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           79,
		Name:                         "Cedar Bakery",
		Address:                      "90 Baker Street, Cedar Grove",
		Description:                  "A bakery offering a variety of fresh baked goods and pastries.",
		Type:                         models.Commercial,
		Subtype:                      models.Bakery,
		BaseRent:                     7500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1800000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
	{
		ID:                           80,
		Name:                         "Playtime Arcade",
		Address:                      "100 Arcade Avenue, Cedar Grove",
		Description:                  "An arcade offering a variety of games and fun for families and friends.",
		Type:                         models.Commercial,
		Subtype:                      models.Amusement,
		BaseRent:                     8500.0,
		RentBoost:                    0.0,
		Owned:                        false,
		UpgradeLevel:                 0,
		Upgrades:                     []models.Upgrade{},
		UpgradePaths:                 cedarCommercialUpgradePaths(),
		Price:                        1750000.0,
		PlayerID:                     0,
		OccupancyRate:                0.0,
		TenantSatisfaction:           0,
		PurchaseDate:                 time.Time{},
		ProrateRent:                  false,
		NeighborhoodID:               4,
		UpgradedNeighborhoodRentBoost: 0.0,
	},
}

// Helper functions to define upgrade paths
func cedarResidentialUpgradePaths() map[string][]models.Upgrade {
	return map[string][]models.Upgrade{
		"Cozy Enhancements": {
			{
				Name:           "Insulation Upgrade",
				ID:             "cozy_1",
				Level:          1,
				Cost:           3000.0,
				RentIncrease:   150.0,
				DaysToComplete: 5,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Energy-efficient Appliances",
				ID:             "cozy_2",
				Level:          2,
				Cost:           6000.0,
				RentIncrease:   300.0,
				DaysToComplete: 10,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Smart Thermostat",
				ID:             "cozy_3",
				Level:          3,
				Cost:           9000.0,
				RentIncrease:   450.0,
				DaysToComplete: 15,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
		},
		"Modern Upgrades": {
			{
				Name:           "Open-plan Kitchen",
				ID:             "mod_1",
				Level:          1,
				Cost:           4000.0,
				RentIncrease:   200.0,
				DaysToComplete: 7,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Smart Lighting",
				ID:             "mod_2",
				Level:          2,
				Cost:           8000.0,
				RentIncrease:   400.0,
				DaysToComplete: 14,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Home Automation System",
				ID:             "mod_3",
				Level:          3,
				Cost:           12000.0,
				RentIncrease:   600.0,
				DaysToComplete: 21,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
		},
		"Exterior Enhancements": {
			{
				Name:           "New Patio",
				ID:             "ext_1",
				Level:          1,
				Cost:           5000.0,
				RentIncrease:   250.0,
				DaysToComplete: 5,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Fire Pit Installation",
				ID:             "ext_2",
				Level:          2,
				Cost:           10000.0,
				RentIncrease:   500.0,
				DaysToComplete: 10,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Outdoor Kitchen",
				ID:             "ext_3",
				Level:          3,
				Cost:           15000.0,
				RentIncrease:   750.0,
				DaysToComplete: 15,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
		},
	}
}

func cedarCommercialUpgradePaths() map[string][]models.Upgrade {
	return map[string][]models.Upgrade{
		"Facility Enhancements": {
			{
				Name:           "Extended Operating Hours",
				ID:             "fac_1",
				Level:          1,
				Cost:           5000.0,
				RentIncrease:   250.0,
				DaysToComplete: 5,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Advanced Security Systems",
				ID:             "fac_2",
				Level:          2,
				Cost:           10000.0,
				RentIncrease:   500.0,
				DaysToComplete: 10,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Automated Inventory Management",
				ID:             "fac_3",
				Level:          3,
				Cost:           15000.0,
				RentIncrease:   750.0,
				DaysToComplete: 15,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
		},
		"Technology Upgrades": {
			{
				Name:           "Digital POS System",
				ID:             "op_1",
				Level:          1,
				Cost:           4000.0,
				RentIncrease:   200.0,
				DaysToComplete: 7,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Basic Inventory Software",
				ID:             "op_2",
				Level:          2,
				Cost:           8000.0,
				RentIncrease:   400.0,
				DaysToComplete: 14,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Automated Reporting",
				ID:             "op_3",
				Level:          3,
				Cost:           12000.0,
				RentIncrease:   600.0,
				DaysToComplete: 21,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
		},
		"Marketing Enhancements": {
			{
				Name:           "Local Advertising Campaign",
				ID:             "mar_1",
				Level:          1,
				Cost:           5000.0,
				RentIncrease:   250.0,
				DaysToComplete: 10,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Social Media Marketing",
				ID:             "mar_2",
				Level:          2,
				Cost:           10000.0,
				RentIncrease:   500.0,
				DaysToComplete: 20,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
			{
				Name:           "Community Events Sponsorship",
				ID:             "mar_3",
				Level:          3,
				Cost:           15000.0,
				RentIncrease:   750.0,
				DaysToComplete: 30,
				PurchaseDate:   time.Time{},
				Prerequisite:   nil,
				Applied:        false,
			},
		},
	}
}
