# City Developer Backend

City Developer is a backend system that powers a city management game. It is built using Go with an Entity-Component-System (ECS) architecture to enable modularity, scalability, and flexibility.

---

## Features

### Core Gameplay Mechanics
- **Property Management**
  - Buy, sell, and upgrade properties following their upgrade tree.
  - Handles property ownership, rent collection, and rent boosts based on upgrades and neighborhood conditions.

- **Time Management**
  - Supports variable time advancement speeds (e.g., real-time, fast-forward).
  - Accurate rent proration for purchases and upgrades with edge case handling for different calendar months and leap years.

- **Player Management**
  - Track player funds, property ownership, and transactions.
  - Distribute rent and upgrade costs automatically.

### Systems
- **Income System**
  - Calculates rent based on ownership duration and upgrades.
  - Handles prorated rent for partial months and upgrades completed mid-month.

- **Neighborhood System**
  - Boosts property rents based on neighborhood upgrades.

- **Property Management System**
  - Tracks property ownership and interactions.

- **Time System**
  - Advances game time and synchronizes actions with real-time or accelerated gameplay.

---

## Project Structure

```
city-developer/
├── actions/              # Handles game actions (buy, sell, upgrade, control time)
├── components/           # Core ECS components (Player, Property, GameTime, etc.)
├── ecs/                  # ECS core (World, Entity, Component)
├── game/                 # Game initialization logic
├── models/               # Data models (Player, Property, Upgrade, Neighborhood)
├── server/               # HTTP server with CORS support
├── systems/              # ECS systems (Income, Neighborhood, Time)
└── utils/                # Utility functions
```

---

## API Endpoints

### Actions
All game actions are handled via the `/actions` endpoint. Supported actions include:
- **`buy_property`**
- **`sell_property`**
- **`upgrade_property`**
- **`control_time`**

Example Request:
```json
POST /actions
{
  "action": "buy_property",
  "payload": {
    "property_id": 1,
    "player_id": 2
  }
}
```

### State
The `/state` endpoint retrieves the current state of the game, including entities and components.

```
{
    "Entities": {
        "0": {
            "ID": 0,
            "Components": {
                "GameTime": {
                    "Time": {
                        "CurrentDate": "2023-01-04T00:00:00Z",
                        "IsPaused": false,
                        "SpeedMultiplier": 1,
                        "NewMonth": false,
                        "LastUpdated": "2023-01-03T00:00:00Z"
                    }
                }
            }
        },
        "1": {
            "ID": 1,
            "Components": {
                "PlayerComponent": {
                    "Player": {
                        "ID": 1,
                        "Funds": 10000,
                        "Properties": null
                    }
                }
            }
        },
        "2": {
            "ID": 2,
            "Components": {
                "PropertyComponent": {
                    "Property": {
                        "Name": "Residential 1",
                        "Type": "Residential",
                        "Subtype": "SingleFamily",
                        "BaseRent": 1000,
                        "RentBoost": 0,
                        "Owned": false,
                        "UpgradeLevel": 0,
                        "Upgrades": null,
                        "UpgradePaths": {
                            "Efficiency": [
                                {
                                    "Name": "Solar Panels",
                                    "ID": "efficiency_1",
                                    "Level": 0,
                                    "Cost": 8000,
                                    "RentIncrease": 300,
                                    "DaysToComplete": 10,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": null,
                                    "Applied": false
                                },
                                {
                                    "Name": "Energy-efficient Windows",
                                    "ID": "efficiency_2",
                                    "Level": 0,
                                    "Cost": 12000,
                                    "RentIncrease": 500,
                                    "DaysToComplete": 15,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Solar Panels",
                                        "ID": "efficiency_1",
                                        "Level": 0,
                                        "Cost": 8000,
                                        "RentIncrease": 300,
                                        "DaysToComplete": 10,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": null,
                                        "Applied": false
                                    },
                                    "Applied": false
                                },
                                {
                                    "Name": "High-efficiency HVAC",
                                    "ID": "efficiency_3",
                                    "Level": 0,
                                    "Cost": 20000,
                                    "RentIncrease": 800,
                                    "DaysToComplete": 20,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Energy-efficient Windows",
                                        "ID": "efficiency_2",
                                        "Level": 0,
                                        "Cost": 12000,
                                        "RentIncrease": 500,
                                        "DaysToComplete": 15,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": {
                                            "Name": "Solar Panels",
                                            "ID": "efficiency_1",
                                            "Level": 0,
                                            "Cost": 8000,
                                            "RentIncrease": 300,
                                            "DaysToComplete": 10,
                                            "PurchaseDate": "0001-01-01T00:00:00Z",
                                            "Prerequisite": null,
                                            "Applied": false
                                        },
                                        "Applied": false
                                    },
                                    "Applied": false
                                }
                            ],
                            "Luxury": [
                                {
                                    "Name": "Renovated Interior",
                                    "ID": "luxury_1",
                                    "Level": 0,
                                    "Cost": 10000,
                                    "RentIncrease": 500,
                                    "DaysToComplete": 7,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": null,
                                    "Applied": false
                                },
                                {
                                    "Name": "Smart Home Automation",
                                    "ID": "luxury_2",
                                    "Level": 0,
                                    "Cost": 20000,
                                    "RentIncrease": 1000,
                                    "DaysToComplete": 14,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Renovated Interior",
                                        "ID": "luxury_1",
                                        "Level": 0,
                                        "Cost": 10000,
                                        "RentIncrease": 500,
                                        "DaysToComplete": 7,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": null,
                                        "Applied": false
                                    },
                                    "Applied": false
                                },
                                {
                                    "Name": "Premium Fixtures",
                                    "ID": "luxury_3",
                                    "Level": 0,
                                    "Cost": 30000,
                                    "RentIncrease": 1500,
                                    "DaysToComplete": 21,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Smart Home Automation",
                                        "ID": "luxury_2",
                                        "Level": 0,
                                        "Cost": 20000,
                                        "RentIncrease": 1000,
                                        "DaysToComplete": 14,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": {
                                            "Name": "Renovated Interior",
                                            "ID": "luxury_1",
                                            "Level": 0,
                                            "Cost": 10000,
                                            "RentIncrease": 500,
                                            "DaysToComplete": 7,
                                            "PurchaseDate": "0001-01-01T00:00:00Z",
                                            "Prerequisite": null,
                                            "Applied": false
                                        },
                                        "Applied": false
                                    },
                                    "Applied": false
                                }
                            ]
                        },
                        "Price": 10000,
                        "PlayerID": 0,
                        "OccupancyRate": 0,
                        "TenantSatisfaction": 0,
                        "PurchaseDate": "0001-01-01T00:00:00Z",
                        "NeighborhoodID": 1,
                        "UgradedNeighborhoodRentBoost": 0
                    }
                }
            }
        },
        "3": {
            "ID": 3,
            "Components": {
                "PropertyComponent": {
                    "Property": {
                        "Name": "Downtown Restaurant",
                        "Type": "Commercial",
                        "Subtype": "Restaurant",
                        "BaseRent": 5000,
                        "RentBoost": 0,
                        "Owned": false,
                        "UpgradeLevel": 0,
                        "Upgrades": null,
                        "UpgradePaths": {
                            "Capacity": [
                                {
                                    "Name": "Expand Seating Area",
                                    "ID": "capacity_1",
                                    "Level": 0,
                                    "Cost": 15000,
                                    "RentIncrease": 700,
                                    "DaysToComplete": 12,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": null,
                                    "Applied": false
                                },
                                {
                                    "Name": "Add Outdoor Seating",
                                    "ID": "capacity_2",
                                    "Level": 0,
                                    "Cost": 25000,
                                    "RentIncrease": 1200,
                                    "DaysToComplete": 18,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Expand Seating Area",
                                        "ID": "capacity_1",
                                        "Level": 0,
                                        "Cost": 15000,
                                        "RentIncrease": 700,
                                        "DaysToComplete": 12,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": null,
                                        "Applied": false
                                    },
                                    "Applied": false
                                }
                            ],
                            "Technology": [
                                {
                                    "Name": "Install POS System",
                                    "ID": "tech_1",
                                    "Level": 0,
                                    "Cost": 5000,
                                    "RentIncrease": 200,
                                    "DaysToComplete": 5,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": null,
                                    "Applied": false
                                },
                                {
                                    "Name": "Automated Inventory Management",
                                    "ID": "tech_2",
                                    "Level": 0,
                                    "Cost": 10000,
                                    "RentIncrease": 400,
                                    "DaysToComplete": 10,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Install POS System",
                                        "ID": "tech_1",
                                        "Level": 0,
                                        "Cost": 5000,
                                        "RentIncrease": 200,
                                        "DaysToComplete": 5,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": null,
                                        "Applied": false
                                    },
                                    "Applied": false
                                },
                                {
                                    "Name": "Customer Loyalty App",
                                    "ID": "tech_3",
                                    "Level": 0,
                                    "Cost": 15000,
                                    "RentIncrease": 600,
                                    "DaysToComplete": 14,
                                    "PurchaseDate": "0001-01-01T00:00:00Z",
                                    "Prerequisite": {
                                        "Name": "Automated Inventory Management",
                                        "ID": "tech_2",
                                        "Level": 0,
                                        "Cost": 10000,
                                        "RentIncrease": 400,
                                        "DaysToComplete": 10,
                                        "PurchaseDate": "0001-01-01T00:00:00Z",
                                        "Prerequisite": {
                                            "Name": "Install POS System",
                                            "ID": "tech_1",
                                            "Level": 0,
                                            "Cost": 5000,
                                            "RentIncrease": 200,
                                            "DaysToComplete": 5,
                                            "PurchaseDate": "0001-01-01T00:00:00Z",
                                            "Prerequisite": null,
                                            "Applied": false
                                        },
                                        "Applied": false
                                    },
                                    "Applied": false
                                }
                            ]
                        },
                        "Price": 50000,
                        "PlayerID": 0,
                        "OccupancyRate": 0,
                        "TenantSatisfaction": 0,
                        "PurchaseDate": "0001-01-01T00:00:00Z",
                        "NeighborhoodID": 1,
                        "UgradedNeighborhoodRentBoost": 0
                    }
                }
            }
        }
    },
    "Systems": [
        {},
        {
            "Neighborhoods": {
                "1": {
                    "ID": 1,
                    "Name": "Downtown",
                    "PropertyIDs": [
                        2,
                        3,
                        2,
                        3
                    ],
                    "AveragePropertyValue": 30000,
                    "RentBoostThreshold": 50,
                    "RentBoostAmount": 10
                }
            }
        },
        {},
        {}
    ]
}
```

---

## Proration Rules for Rent and Upgrades

### Properties
- Rent collection begins the day after a property is purchased.
- Prorated rent is calculated based on the number of days the property is owned in a month.

### Upgrades
- Rent increases from upgrades take effect the day after the upgrade is completed.
- Proration applies to rent increases for upgrades completed mid-month.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/markbmullins/city-developer.git
   cd city-developer
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. The server will be available at `http://localhost:8080`.

---

## Development

### Testing
Run the test suite with:
```bash
go test ./...
```

### Key Concepts
- **ECS Architecture**: Entities, Components, and Systems form the core game logic.
- **Modularity**: Each component and system is isolated, making it easy to extend.

---