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
      "GameTime-0": {
        "Key": {
          "EntityType": "GameTime",
          "ID": 0
        },
        "Components": {
          "GameTime": {
            "CurrentDate": "2023-01-04T00:00:00Z",
            "IsPaused": false,
            "SpeedMultiplier": 1,
            "NewMonth": false,
            "LastUpdated": "2023-01-03T00:00:00Z",
            "RentCollectionDay": 1
          }
        }
      },
      "Player-1": {
        "Key": {
          "EntityType": "Player",
          "ID": 1
        },
        "Components": {
          "Funds": {
            "Amount": 100000000
          },
          "Nameable": {
            "Name": "Mark"
          }
        }
      },
      "Property-61": {
        "Key": {
          "EntityType": "Property",
          "ID": 61
        },
        "Components": {
          "Addressable": {
            "Address": "101 Maplewood Lane, Cedar Grove"
          },
          "Classifiable": {
            "Type": "Residential",
            "Subtype": "SingleFamily"
          },
          "Describable": {
            "Description": "A cozy single-family home with a large backyard and modern amenities."
          },
          "Groupable": {
            "GroupID": 4
          },
          "Nameable": {
            "Name": "Maplewood Lane House"
          },
          "Ownable": {
            "OwnerID": 0,
            "Owned": false
          },
          "Purchaseable": {
            "Cost": 300000,
            "PurchaseDate": "0001-01-01T00:00:00Z"
          },
          "Rentable": {
            "BaseRent": 1800,
            "RentBoost": 0,
            "LastRentCollectionDate": "0001-01-01T00:00:00Z"
          },
          "Upgradable": {
            "PossibleUpgrades": {
              "Cozy Enhancements": [
                {
                  "Name": "Insulation Upgrade",
                  "Level": 1,
                  "Cost": 3000,
                  "RentIncrease": 150,
                  "DaysToComplete": 5,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Energy-efficient Appliances",
                  "Level": 2,
                  "Cost": 6000,
                  "RentIncrease": 300,
                  "DaysToComplete": 10,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Smart Thermostat",
                  "Level": 3,
                  "Cost": 9000,
                  "RentIncrease": 450,
                  "DaysToComplete": 15,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                }
              ],
              "Exterior Enhancements": [
                {
                  "Name": "New Patio",
                  "Level": 1,
                  "Cost": 5000,
                  "RentIncrease": 250,
                  "DaysToComplete": 5,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Fire Pit Installation",
                  "Level": 2,
                  "Cost": 10000,
                  "RentIncrease": 500,
                  "DaysToComplete": 10,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Outdoor Kitchen",
                  "Level": 3,
                  "Cost": 15000,
                  "RentIncrease": 750,
                  "DaysToComplete": 15,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                }
              ],
              "Modern Upgrades": [
                {
                  "Name": "Open-plan Kitchen",
                  "Level": 1,
                  "Cost": 4000,
                  "RentIncrease": 200,
                  "DaysToComplete": 7,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Smart Lighting",
                  "Level": 2,
                  "Cost": 8000,
                  "RentIncrease": 400,
                  "DaysToComplete": 14,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Home Automation System",
                  "Level": 3,
                  "Cost": 12000,
                  "RentIncrease": 600,
                  "DaysToComplete": 21,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                }
              ]
            },
            "AppliedUpgrades": []
          }
        }
      },
      "Property-62": {
        "Key": {
          "EntityType": "Property",
          "ID": 62
        },
        "Components": {
          "Addressable": {
            "Address": "202 Sunnybrook Drive, Cedar Grove"
          },
          "Classifiable": {
            "Type": "Residential",
            "Subtype": "Townhome"
          },
          "Describable": {
            "Description": "A charming townhome with modern finishes and a community garden."
          },
          "Groupable": {
            "GroupID": 4
          },
          "Nameable": {
            "Name": "Sunnybrook Townhome"
          },
          "Ownable": {
            "OwnerID": 0,
            "Owned": false
          },
          "Purchaseable": {
            "Cost": 350000,
            "PurchaseDate": "0001-01-01T00:00:00Z"
          },
          "Rentable": {
            "BaseRent": 2200,
            "RentBoost": 0,
            "LastRentCollectionDate": "0001-01-01T00:00:00Z"
          },
          "Upgradable": {
            "PossibleUpgrades": {
              "Cozy Enhancements": [
                {
                  "Name": "Insulation Upgrade",
                  "Level": 1,
                  "Cost": 3000,
                  "RentIncrease": 150,
                  "DaysToComplete": 5,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Energy-efficient Appliances",
                  "Level": 2,
                  "Cost": 6000,
                  "RentIncrease": 300,
                  "DaysToComplete": 10,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Smart Thermostat",
                  "Level": 3,
                  "Cost": 9000,
                  "RentIncrease": 450,
                  "DaysToComplete": 15,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                }
              ],
              "Exterior Enhancements": [
                {
                  "Name": "New Patio",
                  "Level": 1,
                  "Cost": 5000,
                  "RentIncrease": 250,
                  "DaysToComplete": 5,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Fire Pit Installation",
                  "Level": 2,
                  "Cost": 10000,
                  "RentIncrease": 500,
                  "DaysToComplete": 10,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Outdoor Kitchen",
                  "Level": 3,
                  "Cost": 15000,
                  "RentIncrease": 750,
                  "DaysToComplete": 15,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                }
              ],
              "Modern Upgrades": [
                {
                  "Name": "Open-plan Kitchen",
                  "Level": 1,
                  "Cost": 4000,
                  "RentIncrease": 200,
                  "DaysToComplete": 7,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Smart Lighting",
                  "Level": 2,
                  "Cost": 8000,
                  "RentIncrease": 400,
                  "DaysToComplete": 14,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                },
                {
                  "Name": "Home Automation System",
                  "Level": 3,
                  "Cost": 12000,
                  "RentIncrease": 600,
                  "DaysToComplete": 21,
                  "PurchaseDate": "0001-01-01T00:00:00Z",
                  "Prerequisite": null,
                  "Applied": false
                }
              ]
            },
            "AppliedUpgrades": []
          }
        }
      }
    },
    "Systems": [
      {},
      {},
      {}
    ],
    "Indexes": {
      "components.Addressable": {
        "Property-61": {
          "Key": {
            "EntityType": "Property",
            "ID": 61
          },
          "Components": {
            "Addressable": {
              "Address": "101 Maplewood Lane, Cedar Grove"
            },
            "Classifiable": {
              "Type": "Residential",
              "Subtype": "SingleFamily"
            },
            "Describable": {
              "Description": "A cozy single-family home with a large backyard and modern amenities."
            },
            "Groupable": {
              "GroupID": 4
            },
            "Nameable": {
              "Name": "Maplewood Lane House"
            },
            "Ownable": {
              "OwnerID": 0,
              "Owned": false
            },
            "Purchaseable": {
              "Cost": 300000,
              "PurchaseDate": "0001-01-01T00:00:00Z"
            },
            "Rentable": {
              "BaseRent": 1800,
              "RentBoost": 0,
              "LastRentCollectionDate": "0001-01-01T00:00:00Z"
            },
            "Upgradable": {
              "PossibleUpgrades": {
                "Cozy Enhancements": [
                  {
                    "Name": "Insulation Upgrade",
                    "Level": 1,
                    "Cost": 3000,
                    "RentIncrease": 150,
                    "DaysToComplete": 5,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  },
                  {
                    "Name": "Energy-efficient Appliances",
                    "Level": 2,
                    "Cost": 6000,
                    "RentIncrease": 300,
                    "DaysToComplete": 10,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  },
                  {
                    "Name": "Smart Thermostat",
                    "Level": 3,
                    "Cost": 9000,
                    "RentIncrease": 450,
                    "DaysToComplete": 15,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  }
                ],
                "Exterior Enhancements": [
                  {
                    "Name": "New Patio",
                    "Level": 1,
                    "Cost": 5000,
                    "RentIncrease": 250,
                    "DaysToComplete": 5,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  },
                  {
                    "Name": "Fire Pit Installation",
                    "Level": 2,
                    "Cost": 10000,
                    "RentIncrease": 500,
                    "DaysToComplete": 10,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  },
                  {
                    "Name": "Outdoor Kitchen",
                    "Level": 3,
                    "Cost": 15000,
                    "RentIncrease": 750,
                    "DaysToComplete": 15,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  }
                ],
                "Modern Upgrades": [
                  {
                    "Name": "Open-plan Kitchen",
                    "Level": 1,
                    "Cost": 4000,
                    "RentIncrease": 200,
                    "DaysToComplete": 7,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  },
                  {
                    "Name": "Smart Lighting",
                    "Level": 2,
                    "Cost": 8000,
                    "RentIncrease": 400,
                    "DaysToComplete": 14,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  },
                  {
                    "Name": "Home Automation System",
                    "Level": 3,
                    "Cost": 12000,
                    "RentIncrease": 600,
                    "DaysToComplete": 21,
                    "PurchaseDate": "0001-01-01T00:00:00Z",
                    "Prerequisite": null,
                    "Applied": false
                  }
                ]
              },
              "AppliedUpgrades": []
            }
          }
        }
      }
      // Other component indexes
    },
    "OwnedPropertiesIndex": {},
    // Index of neighborhood id -> property ids
    "GroupPropertiesIndex": {
      "4": [
        61,
        62,
        63,
        64,
        65,
        66,
        67,
        68,
        69,
        70,
        71,
        72,
        73,
        74,
        75,
        76,
        77,
        78,
        79,
        80
      ]
    },
    "GroupUpgradedPercentages": {},
    "GroupUpgradedCounts": {},
    "Players": [
      {
        "Key": {
          "EntityType": "Player",
          "ID": 1
        },
        "Components": {
          "Funds": {
            "Amount": 100000000
          },
          "Nameable": {
            "Name": "Mark"
          }
        }
      }
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