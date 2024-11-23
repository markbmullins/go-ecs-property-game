// src/types/index.ts

export interface World {
  Entities: Record<number, Entity>;
  Systems: System[];
}

export interface Entity {
  ID: number;
  Components: Record<string, Component>;
}

export type Component = GameTimeComponent | PlayerComponent | PropertyComponent;

export interface GameTime {
  CurrentDate: string; // ISO string
  IsPaused: boolean;
  SpeedMultiplier: number;
  NewMonth: boolean;
}

export interface Player {
  ID: number;
  Funds: number;
  Properties: Property[];
}

export interface PlayerComponent {
  Player: Player;
}

export interface GameTimeComponent {
  Time: GameTime;
}

export interface Property {
  Name: string;
  Type: PropertyType;
  Subtype: PropertySubtype;
  BaseRent: number;
  RentBoost: number;
  Owned: boolean;
  UpgradeLevel: number;
  Upgrades: Upgrade[];
  UpgradePaths: Record<string, Upgrade[]>;
  Price: number;
  PlayerID: number;
  OccupancyRate: number;
  TenantSatisfaction: number;
  PurchaseDate: string; // ISO string
  ProrateRent: boolean;
  NeighborhoodID: number;
  UgradedNeighborhoodRentBoost: number;
}

export interface PropertyComponent {
  Property: Property;
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
}

export interface System {
  // Define if you have any specific properties
}
