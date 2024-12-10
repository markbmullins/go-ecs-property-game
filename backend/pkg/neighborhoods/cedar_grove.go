// cedar_grove.go
package neighborhoods

import (
	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/entities"
)

func GetCedarGroveProperties() []*ecs.Entity {
	return append(CedarResidential, CedarCommercial...)
}

func InitializeCedarGroveUpgrades() {
	for _, property := range CedarResidential {
		entities.AddUpgradesToProperty(property, cedarResidentialUpgradePaths())
	}

	for _, property := range CedarCommercial {
		entities.AddUpgradesToProperty(property, cedarCommercialUpgradePaths())
	}
}

// Residential Properties in Cedar Grove
var CedarResidential = []*ecs.Entity{
	entities.CreateProperty(
		"Maplewood Lane House",
		"101 Maplewood Lane, Cedar Grove",
		"A cozy single-family home with a large backyard and modern amenities.",
		components.Residential,
		components.SingleFamily,
		1800.0,
		300000.0,
		4,
	),
	entities.CreateProperty(
		"Sunnybrook Townhome",
		"202 Sunnybrook Drive, Cedar Grove",
		"A charming townhome with modern finishes and a community garden.",
		components.Residential,
		components.Townhome,
		2200.0,
		350000.0,
		4,
	),
	entities.CreateProperty(
		"Oakwood Apartments",
		"303 Oakwood Road, Cedar Grove",
		"Modern apartments with access to shared recreational facilities and secure parking.",
		components.Residential,
		components.Apartment,
		1500.0,
		260000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Grove Condos",
		"404 Cedar Boulevard, Cedar Grove",
		"Condominiums with private balconies and state-of-the-art home automation systems.",
		components.Residential,
		components.Condo,
		2000.0,
		400000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Grove Estates",
		"505 Cedar Lane, Cedar Grove",
		"Spacious multifamily residences with modern amenities and landscaped gardens.",
		components.Residential,
		components.Multifamily,
		1900.0,
		380000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Grove Villas",
		"606 Villa Avenue, Cedar Grove",
		"Luxurious villas featuring private pools and high-end finishes.",
		components.Residential,
		components.SingleFamily,
		2200.0,
		420000.0,
		4,
	),
	entities.CreateProperty(
		"Maplewood Condos",
		"707 Maplewood Street, Cedar Grove",
		"Condominiums with smart home integrations and access to communal lounges.",
		components.Residential,
		components.Condo,
		2100.0,
		400000.0,
		4,
	),
	entities.CreateProperty(
		"Sunnybrook Apartments",
		"808 Sunnybrook Road, Cedar Grove",
		"Apartments with modern designs and access to recreational facilities.",
		components.Residential,
		components.Apartment,
		1700.0,
		340000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Grove Flats",
		"909 Cedar Circle, Cedar Grove",
		"Modern flats with integrated smart systems and community amenities.",
		components.Residential,
		components.Apartment,
		1600.0,
		330000.0,
		4,
	),
	entities.CreateProperty(
		"Oakridge Apartments",
		"1001 Oakridge Road, Cedar Grove",
		"Spacious apartments with eco-friendly features and access to green spaces.",
		components.Residential,
		components.Apartment,
		1500.0,
		260000.0,
		4,
	),
}

// Commercial Properties in Cedar Grove
var CedarCommercial = []*ecs.Entity{
	entities.CreateProperty(
		"Cozy Corner Café",
		"10 Cozy Street, Cedar Grove",
		"A friendly neighborhood café serving fresh coffee and pastries.",
		components.Commercial,
		components.Cafe,
		4000.0,
		900000.0,
		4,
	),
	entities.CreateProperty(
		"Suburban Shoppe",
		"20 Market Lane, Cedar Grove",
		"A local retail store offering a variety of household items and essentials.",
		components.Commercial,
		components.FurnitureStore,
		5000.0,
		1100000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Gym",
		"30 Fitness Boulevard, Cedar Grove",
		"A comprehensive fitness center with modern equipment and personal trainers.",
		components.Commercial,
		components.Gym,
		6500.0,
		1400000.0,
		4,
	),
	entities.CreateProperty(
		"Playtime Arcade",
		"40 Fun Avenue, Cedar Grove",
		"An arcade offering a variety of games and entertainment for all ages.",
		components.Commercial,
		components.Arcade,
		8500.0,
		1750000.0,
		4,
	),
	entities.CreateProperty(
		"Tech Mart",
		"50 Tech Avenue, Cedar Grove",
		"A retail store specializing in the latest tech gadgets and accessories.",
		components.Commercial,
		components.ElectronicsStore,
		7500.0,
		1600000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Pharmacy",
		"60 Health Street, Cedar Grove",
		"A pharmacy supplying affordable medications and health products.",
		components.Commercial,
		components.Clinic,
		8500.0,
		1750000.0,
		4,
	),
	entities.CreateProperty(
		"Simple Salon",
		"70 Beauty Lane, Cedar Grove",
		"A basic salon offering essential beauty and grooming services.",
		components.Commercial,
		components.Salon,
		3200.0,
		750000.0,
		4,
	),
	entities.CreateProperty(
		"Affordable Arcade",
		"80 Fun Boulevard, Cedar Grove",
		"An arcade providing a variety of budget-friendly games and entertainment options.",
		components.Commercial,
		components.Arcade,
		8500.0,
		1750000.0,
		4,
	),
	entities.CreateProperty(
		"Cedar Bakery",
		"90 Baker Street, Cedar Grove",
		"A bakery offering a variety of fresh baked goods and pastries.",
		components.Commercial,
		components.Bakery,
		7500.0,
		1800000.0,
		4,
	),
	entities.CreateProperty(
		"Playtime Arcade",
		"100 Arcade Avenue, Cedar Grove",
		"An arcade offering a variety of games and fun for families and friends.",
		components.Commercial,
		components.Arcade,
		8500.0,
		1750000.0,
		4,
	),
}

func cedarResidentialUpgradePaths() map[string][]*components.Upgrade {
	return map[string][]*components.Upgrade{
		"Cozy Enhancements": {
			entities.CreateUpgrade("Insulation Upgrade", 1, 3000.0, 150.0, 5, nil),
			entities.CreateUpgrade("Energy-efficient Appliances", 2, 6000.0, 300.0, 10, nil),
			entities.CreateUpgrade("Smart Thermostat", 3, 9000.0, 450.0, 15, nil),
		},
		"Modern Upgrades": {
			entities.CreateUpgrade("Open-plan Kitchen", 1, 4000.0, 200.0, 7, nil),
			entities.CreateUpgrade("Smart Lighting", 2, 8000.0, 400.0, 14, nil),
			entities.CreateUpgrade("Home Automation System", 3, 12000.0, 600.0, 21, nil),
		},
		"Exterior Enhancements": {
			entities.CreateUpgrade("New Patio", 1, 5000.0, 250.0, 5, nil),
			entities.CreateUpgrade("Fire Pit Installation", 2, 10000.0, 500.0, 10, nil),
			entities.CreateUpgrade("Outdoor Kitchen", 3, 15000.0, 750.0, 15, nil),
		},
	}
}

func cedarCommercialUpgradePaths() map[string][]*components.Upgrade {
	return map[string][]*components.Upgrade{
		"Facility Enhancements": {
			entities.CreateUpgrade("Extended Operating Hours", 1, 5000.0, 250.0, 5, nil),
			entities.CreateUpgrade("Advanced Security Systems", 2, 10000.0, 500.0, 10, nil),
			entities.CreateUpgrade("Automated Inventory Management", 3, 15000.0, 750.0, 15, nil),
		},
		"Technology Upgrades": {
			entities.CreateUpgrade("Digital POS System", 1, 4000.0, 200.0, 7, nil),
			entities.CreateUpgrade("Basic Inventory Software", 2, 8000.0, 400.0, 14, nil),
			entities.CreateUpgrade("Automated Reporting", 3, 12000.0, 600.0, 21, nil),
		},
		"Marketing Enhancements": {
			entities.CreateUpgrade("Local Advertising Campaign", 1, 5000.0, 250.0, 10, nil),
			entities.CreateUpgrade("Social Media Marketing", 2, 10000.0, 500.0, 20, nil),
			entities.CreateUpgrade("Community Events Sponsorship", 3, 15000.0, 750.0, 30, nil),
		},
	}
}
