// downtown_district.go
package neighborhoods

// func GetDowntownDistrictNeighborhood() *components.Neighborhood {
// 	propertyIDs := make([]int, 0, len(DowntownResidential)+len(DowntownCommercial))
// 	for _, prop := range DowntownResidential {
// 		propertyIDs = append(propertyIDs, prop.ID)
// 	}
// 	for _, prop := range DowntownCommercial {
// 		propertyIDs = append(propertyIDs, prop.ID)
// 	}

// 	return &components.Neighborhood{
// 		ID:                   1,
// 		Name:                 "Downtown District",
// 		PropertyIDs:          propertyIDs,
// 		AveragePropertyValue: 0.0,
// 		RentBoostThreshold:   50.0, // 50% of properties need to be upgraded
// 		RentBoostPercent:      10.0, // 10% rent boost
// 	}
// }

// func GetDowntownProperties() []components.Property {
// 	return append(DowntownResidential, DowntownCommercial...)
// }

// var DowntownResidential = []components.Property{
// 	{
// 		ID:                           1,
// 		Name:                         "Skyline Lofts",
// 		Address:                      "101 Main Street, Downtown District",
// 		Description:                  "Modern lofts with panoramic city views and open-concept layouts.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     3500.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        750000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           2,
// 		Name:                         "Urban Penthouse",
// 		Address:                      "202 Elm Avenue, Downtown District",
// 		Description:                  "Luxurious penthouse with private terrace and state-of-the-art amenities.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Penthouse,
// 		BaseRent:                     5000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        1500000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           3,
// 		Name:                         "Central Park View",
// 		Address:                      "303 Park Boulevard, Downtown District",
// 		Description:                  "Apartments overlooking Central Park with access to exclusive gym facilities.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     3200.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        680000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           4,
// 		Name:                         "City Lights Residency",
// 		Address:                      "404 Light Street, Downtown District",
// 		Description:                  "High-rise apartments featuring smart home technology and concierge services.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     3400.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        720000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           5,
// 		Name:                         "Downtown Duplex",
// 		Address:                      "505 Commerce Street, Downtown District",
// 		Description:                  "Spacious duplex units with modern interiors and rooftop access.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Duplex,
// 		BaseRent:                     4000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        900000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           6,
// 		Name:                         "Metro Heights",
// 		Address:                      "606 Metro Avenue, Downtown District",
// 		Description:                  "Apartments with direct access to subway stations and premium parking.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     3100.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        690000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           7,
// 		Name:                         "Skyview Apartments",
// 		Address:                      "707 Skyview Road, Downtown District",
// 		Description:                  "Modern apartments offering stunning skyline views and luxury amenities.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     3600.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        800000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           8,
// 		Name:                         "City Center Condos",
// 		Address:                      "808 Center Street, Downtown District",
// 		Description:                  "Condominiums with exclusive access to rooftop lounges and fitness centers.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Condo,
// 		BaseRent:                     3800.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        850000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           9,
// 		Name:                         "Downtown Terrace",
// 		Address:                      "909 Terrace Lane, Downtown District",
// 		Description:                  "Luxurious terrace apartments with private balconies and high-end finishes.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     4200.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        950000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           10,
// 		Name:                         "Urban Oasis",
// 		Address:                      "1001 Oasis Street, Downtown District",
// 		Description:                  "Eco-friendly apartments with green spaces and sustainable features.",
// 		Type:                         components.Residential,
// 		Subtype:                      components.Apartment,
// 		BaseRent:                     3700.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownResidentialUpgradePaths(),
// 		Price:                        820000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// }

// var DowntownCommercial = []components.Property{
// 	{
// 		ID:                           11,
// 		Name:                         "Tech Plaza",
// 		Address:                      "10 Innovation Drive, Downtown District",
// 		Description:                  "A modern office space with state-of-the-art facilities and coworking areas.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.OfficeSpace,
// 		BaseRent:                     12000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        3000000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           12,
// 		Name:                         "Luxury Bistro",
// 		Address:                      "20 Gourmet Street, Downtown District",
// 		Description:                  "An upscale dining experience with gourmet dishes and fine wines.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Restaurant,
// 		BaseRent:                     8000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        2000000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           13,
// 		Name:                         "Central Market",
// 		Address:                      "30 Commerce Avenue, Downtown District",
// 		Description:                  "A bustling marketplace offering a variety of goods from local vendors.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.RetailStore,
// 		BaseRent:                     7000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        1800000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           14,
// 		Name:                         "Skyline Fitness Center",
// 		Address:                      "40 Health Boulevard, Downtown District",
// 		Description:                  "A state-of-the-art fitness center with modern equipment and personal trainers.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Gym,
// 		BaseRent:                     9000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        2200000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           15,
// 		Name:                         "Artisan Bakery",
// 		Address:                      "50 Baker Street, Downtown District",
// 		Description:                  "A bakery offering a wide range of artisan breads and pastries.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Bakery,
// 		BaseRent:                     6500.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        1600000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           16,
// 		Name:                         "Elite Data Center",
// 		Address:                      "60 Tech Park, Downtown District",
// 		Description:                  "A secure data center providing cloud services and storage solutions.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.DataCenter,
// 		BaseRent:                     25000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        7000000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           17,
// 		Name:                         "Modern Art Gallery",
// 		Address:                      "70 Gallery Lane, Downtown District",
// 		Description:                  "A contemporary art gallery showcasing local and international artists.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Museum,
// 		BaseRent:                     8000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        2000000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           18,
// 		Name:                         "City Lights Mall",
// 		Address:                      "80 Shopping Avenue, Downtown District",
// 		Description:                  "A large mall featuring a variety of retail stores, eateries, and entertainment options.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.Mall,
// 		BaseRent:                     15000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        5000000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           19,
// 		Name:                         "Nightlife Hub",
// 		Address:                      "90 Party Street, Downtown District",
// 		Description:                  "A vibrant nightclub offering live music, DJ sets, and a dynamic social scene.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.NightClub,
// 		BaseRent:                     12000.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        3000000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// 	{
// 		ID:                           20,
// 		Name:                         "Urban Art Gallery",
// 		Address:                      "100 Gallery Drive, Downtown District",
// 		Description:                  "A contemporary art gallery showcasing local artists with rotating exhibitions.",
// 		Type:                         components.Commercial,
// 		Subtype:                      components.ArtGallery,
// 		BaseRent:                     7500.0,
// 		RentBoost:                    0.0,
// 		Owned:                        false,
// 		Upgrades:                     []components.Upgrade{},
// 		UpgradePaths:                 downtownCommercialUpgradePaths(),
// 		Price:                        1600000.0,
// 		PlayerID:                     0,
// 		OccupancyRate:                0.0,
// 		TenantSatisfaction:           0,
// 		PurchaseDate:                 time.Time{},
// 		NeighborhoodID:               1,
// 	},
// }

// func downtownResidentialUpgradePaths() map[string][]components.Upgrade {
// 	return map[string][]components.Upgrade{
// 		"Modernization": {
// 			{
// 				Name:           "Basic Automation",
// 				ID:             "mod_1",
// 				Level:          1,
// 				Cost:           5000.0,
// 				RentIncrease:   200.0,
// 				DaysToComplete: 7,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Full Automation",
// 				ID:             "mod_2",
// 				Level:          2,
// 				Cost:           10000.0,
// 				RentIncrease:   500.0,
// 				DaysToComplete: 14,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil, // Set prerequisites if any
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Voice-Controlled System",
// 				ID:             "mod_3",
// 				Level:          3,
// 				Cost:           15000.0,
// 				RentIncrease:   800.0,
// 				DaysToComplete: 21,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil, // Set prerequisites if any
// 				Applied:        false,
// 			},
// 		},
// 		"Luxury Interiors": {
// 			{
// 				Name:           "New Furniture",
// 				ID:             "lux_1",
// 				Level:          1,
// 				Cost:           3000.0,
// 				RentIncrease:   100.0,
// 				DaysToComplete: 5,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Designer Finishes",
// 				ID:             "lux_2",
// 				Level:          2,
// 				Cost:           7000.0,
// 				RentIncrease:   300.0,
// 				DaysToComplete: 10,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Custom Artwork",
// 				ID:             "lux_3",
// 				Level:          3,
// 				Cost:           12000.0,
// 				RentIncrease:   500.0,
// 				DaysToComplete: 15,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Community Amenities": {
// 			{
// 				Name:           "Gym Access",
// 				ID:             "comm_1",
// 				Level:          1,
// 				Cost:           4000.0,
// 				RentIncrease:   150.0,
// 				DaysToComplete: 7,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Pool Access",
// 				ID:             "comm_2",
// 				Level:          2,
// 				Cost:           8000.0,
// 				RentIncrease:   350.0,
// 				DaysToComplete: 14,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Rooftop Garden",
// 				ID:             "comm_3",
// 				Level:          3,
// 				Cost:           15000.0,
// 				RentIncrease:   600.0,
// 				DaysToComplete: 21,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 	}
// }

// func downtownCommercialUpgradePaths() map[string][]components.Upgrade {
// 	return map[string][]components.Upgrade{
// 		"Office Design": {
// 			{
// 				Name:           "Open Layout",
// 				ID:             "off_1",
// 				Level:          1,
// 				Cost:           10000.0,
// 				RentIncrease:   500.0,
// 				DaysToComplete: 10,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Quiet Pods",
// 				ID:             "off_2",
// 				Level:          2,
// 				Cost:           20000.0,
// 				RentIncrease:   1000.0,
// 				DaysToComplete: 20,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Smart Desks",
// 				ID:             "off_3",
// 				Level:          3,
// 				Cost:           30000.0,
// 				RentIncrease:   1500.0,
// 				DaysToComplete: 30,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Amenities": {
// 			{
// 				Name:           "Coffee Bar",
// 				ID:             "amen_1",
// 				Level:          1,
// 				Cost:           5000.0,
// 				RentIncrease:   250.0,
// 				DaysToComplete: 5,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Gym",
// 				ID:             "amen_2",
// 				Level:          2,
// 				Cost:           15000.0,
// 				RentIncrease:   750.0,
// 				DaysToComplete: 15,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Relaxation Lounge",
// 				ID:             "amen_3",
// 				Level:          3,
// 				Cost:           25000.0,
// 				RentIncrease:   1250.0,
// 				DaysToComplete: 25,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 		"Technology": {
// 			{
// 				Name:           "Fiber-optic Internet",
// 				ID:             "tech_1",
// 				Level:          1,
// 				Cost:           8000.0,
// 				RentIncrease:   400.0,
// 				DaysToComplete: 8,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "Smart Whiteboards",
// 				ID:             "tech_2",
// 				Level:          2,
// 				Cost:           18000.0,
// 				RentIncrease:   900.0,
// 				DaysToComplete: 18,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 			{
// 				Name:           "VR Meeting Rooms",
// 				ID:             "tech_3",
// 				Level:          3,
// 				Cost:           30000.0,
// 				RentIncrease:   1500.0,
// 				DaysToComplete: 28,
// 				PurchaseDate:   time.Time{},
// 				Prerequisite:   nil,
// 				Applied:        false,
// 			},
// 		},
// 	}
// }
