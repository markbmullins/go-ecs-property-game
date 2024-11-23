

## Server

GET `/state` - and endpoint to fetch the game state
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
                        "ProrateRent": false,
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
                        "ProrateRent": false,
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


POST `/actions` - An endpoint to execute actions within the game. All actions have a follow a standard shape:

Actions have an action name `buy_property` and a `payload`. This follows the Redux actions model.
```
{
  "action": "buy_property",
  "payload": {
    "property_id": 2,
    "player_id": 1
  }
}
```


Example game state after buying
```
{
    "status": "success",
    "message": "Property purchased successfully",
    "data": {
        "Entities": {
            "0": {
                "ID": 0,
                "Components": {
                    "GameTime": {
                        "Time": {
                            "CurrentDate": "2023-02-12T00:00:00Z",
                            "IsPaused": false,
                            "SpeedMultiplier": 1,
                            "NewMonth": false,
                            "LastUpdated": "2023-02-11T00:00:00Z"
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
                            "Funds": 0,
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
                            "Owned": true,
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
                            "PlayerID": 1,
                            "OccupancyRate": 0,
                            "TenantSatisfaction": 0,
                            "PurchaseDate": "2023-02-12T00:00:00Z",
                            "ProrateRent": true,
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
                            "ProrateRent": false,
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
}
```
