// willow_flats.go
package neighborhoods

// func GetWillowFlatsNeighborhood() *components.Neighborhood {
// 	propertyIDs := make([]int, 0, len(WillowResidential)+len(WillowCommercial))
// 	for _, prop := range WillowResidential {
// 		propertyIDs = append(propertyIDs, prop.ID)
// 	}
// 	for _, prop := range WillowCommercial {
// 		propertyIDs = append(propertyIDs, prop.ID)
// 	}

// 	return &components.Neighborhood{
// 		ID:                   5,
// 		Name:                 "Willow Flats",
// 		PropertyIDs:          propertyIDs,
// 		AveragePropertyValue: 0.0,
// 		RentBoostThreshold:   65.0, // 65% of properties need to be upgraded
// 		RentBoostPercent:      20.0, // 20% rent boost
// 	}
// }

// func GetWillowFlatsProperties() []components.Property {
// 	return append(WillowResidential, WillowCommercial...)
// }

// var WillowResidential = []components.Property{
// 	{
// 		ID:                           81,
// 		Name:                         "Willow Flats 1A",
// 		Address:                      "101 Willow Street, Willow Flats",
// 		Description:                  "Affordable apartment with basic amenities and convenient location.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     900.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        150000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           82,
// 		Name:                         "Willow Flats 2B",
// 		Address:                      "202 Willow Avenue, Willow Flats",
// 		Description:                  "Spacious apartment with modern kitchen and comfortable living spaces.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     850.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        140000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           83,
// 		Name:                         "Willow Flats 3C",
// 		Address:                      "303 Willow Road, Willow Flats",
// 		Description:                  "Modern apartment with access to community facilities and secure parking.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     920.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        155000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           84,
// 		Name:                         "Willow Flats 4D",
// 		Address:                      "404 Willow Boulevard, Willow Flats",
// 		Description:                  "Affordable apartment with convenient access to public transportation and local amenities.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     880.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        145000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           85,
// 		Name:                         "Willow Flats 5E",
// 		Address:                      "505 Willow Lane, Willow Flats",
// 		Description:                  "Comfortable apartment with modern appliances and community lounge access.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     910.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        150000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           86,
// 		Name:                         "Willow Flats 6F",
// 		Address:                      "606 Willow Circle, Willow Flats",
// 		Description:                  "Economical apartment with essential amenities and nearby shopping centers.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     870.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        142000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           87,
// 		Name:                         "Willow Flats 7G",
// 		Address:                      "707 Willow Drive, Willow Flats",
// 		Description:                  "Spacious apartment with upgraded interiors and access to community park.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     930.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        160000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           88,
// 		Name:                         "Willow Flats 8H",
// 		Address:                      "808 Willow Terrace, Willow Flats",
// 		Description:                  "Affordable apartment with modern design and access to recreational facilities.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     890.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        148000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           89,
// 		Name:                         "Willow Flats 9I",
// 		Address:                      "909 Willow Court, Willow Flats",
// 		Description:                  "Modern apartment with energy-efficient appliances and community garden access.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     940.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        165000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           90,
// 		Name:                         "Willow Flats 10J",
// 		Address:                      "1001 Willow Way, Willow Flats",
// 		Description:                  "Economical apartment with essential amenities and easy access to local businesses.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     860.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowResidentialUpgradePaths(),
// 		Price:                        143000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// }

// var WillowCommercial = []components.Property{
// 	{
// 		ID:                           91,
// 		Name:                         "Budget Bistro",
// 		Address:                      "10 Budget Street, Willow Flats",
// 		Description:                  "A cost-effective eatery offering affordable meals for the community.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Restaurant,
// 		BaseRent:                     3000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        700000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           92,
// 		Name:                         "Local Market",
// 		Address:                      "20 Market Lane, Willow Flats",
// 		Description:                  "A small retail store providing essential goods and daily necessities.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.RetailStore,
// 		BaseRent:                     3500.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        800000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           93,
// 		Name:                         "Economy Gym",
// 		Address:                      "30 Fitness Avenue, Willow Flats",
// 		Description:                  "A basic fitness center offering essential equipment and classes.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Gym,
// 		BaseRent:                     4000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        900000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           94,
// 		Name:                         "Willow Pharmacy",
// 		Address:                      "40 Health Boulevard, Willow Flats",
// 		Description:                  "A pharmacy supplying affordable medications and health products.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Clinic,
// 		BaseRent:                     3800.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        850000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           95,
// 		Name:                         "Simple Salon",
// 		Address:                      "50 Beauty Lane, Willow Flats",
// 		Description:                  "A basic salon offering essential beauty and grooming services.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Salon,
// 		BaseRent:                     3200.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        750000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           96,
// 		Name:                         "Affordable Arcade",
// 		Address:                      "60 Fun Boulevard, Willow Flats",
// 		Description:                  "An arcade providing a variety of budget-friendly games and entertainment options.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Amusement,
// 		BaseRent:                     4500.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        950000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           97,
// 		Name:                         "Willow Bakery",
// 		Address:                      "70 Baker Street, Willow Flats",
// 		Description:                  "A bakery offering a variety of fresh baked goods and pastries.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Bakery,
// 		BaseRent:                     3600.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        820000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           98,
// 		Name:                         "Neighborhood Gym",
// 		Address:                      "80 Fitness Avenue, Willow Flats",
// 		Description:                  "A community gym providing basic fitness equipment and group classes.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Gym,
// 		BaseRent:                     4200.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        880000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           99,
// 		Name:                         "Community Clinic",
// 		Address:                      "90 Health Boulevard, Willow Flats",
// 		Description:                  "A clinic offering affordable medical services and health products to local residents.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Clinic,
// 		BaseRent:                     3900.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        840000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// 	{
// 		ID:                           100,
// 		Name:                         "Basic Beauty Salon",
// 		Address:                      "100 Beauty Lane, Willow Flats",
// 		Description:                  "A salon providing essential beauty services at affordable prices.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Salon,
// 		BaseRent:                     3300.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 willowCommercialUpgradePaths(),
// 		Price:                        770000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               5,
// 	},
// }

// func willowResidentialUpgradePaths() map[string][]components.Upgrade {
// 	return map[string][]components.Upgrade{
// 		"Basic Enhancements": {
// 			{
// 				Name:           "Insulation Upgrade",
// 				ID:             "basic_1",
// 				Level:          1,
// 				Cost:           2000.0,
// 				RentIncrease:   100.0,
// 				DaysToComplete: 5,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Energy-efficient Appliances",
// 				ID:             "basic_2",
// 				Level:          2,
// 				Cost:           4000.0,
// 				RentIncrease:   200.0,
// 				DaysToComplete: 10,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Smart Thermostat",
// 				ID:             "basic_3",
// 				Level:          3,
// 				Cost:           6000.0,
// 				RentIncrease:   300.0,
// 				DaysToComplete: 15,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Modern Upgrades": {
// 			{
// 				Name:           "Open-plan Kitchen",
// 				ID:             "mod_1",
// 				Level:          1,
// 				Cost:           3000.0,
// 				RentIncrease:   150.0,
// 				DaysToComplete: 7,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Smart Lighting",
// 				ID:             "mod_2",
// 				Level:          2,
// 				Cost:           6000.0,
// 				RentIncrease:   300.0,
// 				DaysToComplete: 14,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Home Automation System",
// 				ID:             "mod_3",
// 				Level:          3,
// 				Cost:           9000.0,
// 				RentIncrease:   450.0,
// 				DaysToComplete: 21,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Exterior Enhancements": {
// 			{
// 				Name:           "New Patio",
// 				ID:             "ext_1",
// 				Level:          1,
// 				Cost:           2500.0,
// 				RentIncrease:   125.0,
// 				DaysToComplete: 5,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Fire Pit Installation",
// 				ID:             "ext_2",
// 				Level:          2,
// 				Cost:           5000.0,
// 				RentIncrease:   250.0,
// 				DaysToComplete: 10,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Outdoor Kitchen",
// 				ID:             "ext_3",
// 				Level:          3,
// 				Cost:           7500.0,
// 				RentIncrease:   375.0,
// 				DaysToComplete: 15,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 	}
// }

// func willowCommercialUpgradePaths() map[string][]components.Upgrade {
// 	return map[string][]components.Upgrade{
// 		"Facility Enhancements": {
// 			{
// 				Name:           "Extended Operating Hours",
// 				ID:             "fac_1",
// 				Level:          1,
// 				Cost:           3000.0,
// 				RentIncrease:   150.0,
// 				DaysToComplete: 5,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Basic Security Systems",
// 				ID:             "fac_2",
// 				Level:          2,
// 				Cost:           6000.0,
// 				RentIncrease:   300.0,
// 				DaysToComplete: 10,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Automated Inventory Management",
// 				ID:             "fac_3",
// 				Level:          3,
// 				Cost:           9000.0,
// 				RentIncrease:   450.0,
// 				DaysToComplete: 15,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Operational Upgrades": {
// 			{
// 				Name:           "Digital POS System",
// 				ID:             "op_1",
// 				Level:          1,
// 				Cost:           4000.0,
// 				RentIncrease:   200.0,
// 				DaysToComplete: 7,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Basic Inventory Software",
// 				ID:             "op_2",
// 				Level:          2,
// 				Cost:           8000.0,
// 				RentIncrease:   400.0,
// 				DaysToComplete: 14,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Automated Reporting",
// 				ID:             "op_3",
// 				Level:          3,
// 				Cost:           12000.0,
// 				RentIncrease:   600.0,
// 				DaysToComplete: 21,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Marketing Enhancements": {
// 			{
// 				Name:           "Local Advertising Campaign",
// 				ID:             "mar_1",
// 				Level:          1,
// 				Cost:           5000.0,
// 				RentIncrease:   250.0,
// 				DaysToComplete: 10,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Social Media Marketing",
// 				ID:             "mar_2",
// 				Level:          2,
// 				Cost:           10000.0,
// 				RentIncrease:   500.0,
// 				DaysToComplete: 20,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Community Events Sponsorship",
// 				ID:             "mar_3",
// 				Level:          3,
// 				Cost:           15000.0,
// 				RentIncrease:   750.0,
// 				DaysToComplete: 30,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 	}
// }
