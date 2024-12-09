// src/types/index.ts

export interface World {
  Entities: Record<number, Entity>;
  Systems: System[];
}

export interface Entity {
  ID: {ID: number; EntityType: string};
  Components: Record<string, Component>;
}

export type Component = GameTime | Player | Property | Upgrade | Neighborhood;

export interface GameTime {
  CurrentDate: string; // ISO string
  IsPaused: boolean;
  SpeedMultiplier: number;
  LastUpdated: string; // ISO string
  NewMonth: boolean;
}

export interface Player {
  ID: number;
  Funds: number;
  Properties: Property[];
}

export interface Property {
  Name: string;
  Type: PropertyType;
  ID: number;
  Subtype: PropertySubtype;
  BaseRent: number;
  RentBoost: number;
  Owned: boolean;
  Upgrades: Upgrade[];
  UpgradePaths: Record<string, Upgrade[]>;
  Price: number;
  PlayerID: number;
  OccupancyRate: number;
  TenantSatisfaction: number;
  PurchaseDate: string; // ISO string
  NeighborhoodID: number;
  Description: string;
  Address: string;
}

export type PropertyType = "Residential" | "Commercial";

export type PropertySubtype =
  | "SingleFamily"
  | "Townhome"
  | "Multifamily"
  | "Apartment"
  | "Condo"
  | "OfficeSpace"
  | "RetailStore"
  | "Warehouse"
  | "Restaurant"
  | "Hotel"
  | "Mall"
  | "Industrial"
  | "Clinic"
  | "DataCenter"
  | "Bar"
  | "NightClub"
  | "Museum"
  | "Amusement"
  | "Factory"
  | "DistributionCenter";

export interface Upgrade {
  Name: string;
  ID: string;
  Level?: number; // Optional, as it's used in some contexts
  Cost: number;
  RentIncrease: number;
  DaysToComplete: number;
  Prerequisite?: Upgrade;
  PurchaseDate: string; // ISO string
  Applied: boolean;
}

export interface System {
  // Define if you have any specific properties
}

export interface Neighborhood {
  ID: number;
  Name: string;
  PropertyIDs: number[]; // List of property IDs in the neighborhood
  AveragePropertyValue: number;
  RentBoostThreshold: number; // Percentage of properties that need to be upgraded
  RentBoostPercent: number; // Boost percentage applied to rents
}
