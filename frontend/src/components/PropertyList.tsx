// src/components/PropertyList.tsx

import React from "react";
import { performAction } from "../api";
import { World, Property } from "../types";

interface PropertyListProps {
  gameState: World;
  onActionComplete: () => void;
}

const PropertyList: React.FC<PropertyListProps> = ({
  gameState,
  onActionComplete,
}) => {
  const entities = Object.values(gameState.Entities);
  const properties = entities.filter(
    (entity) => "Property" in entity.Components
  );

  const playerID = 1; // Assuming the player's entity ID is 1

  const handleBuyProperty = async (propertyID: number) => {
    try {
      await performAction("buy_property", {
        property_id: propertyID,
        player_id: playerID,
      });
      onActionComplete();
    } catch (error) {
      if (error instanceof Error) {
        alert(`Failed to buy property: ${error.message}`);
      } else {
        alert("Failed to buy property: An unexpected error occurred");
      }
    }
  };

  const handleSellProperty = async (propertyID: number) => {
    try {
      await performAction("sell_property", { property_id: propertyID });
      onActionComplete();
    } catch (error) {
      if (error instanceof Error) {
        alert(`Failed to sell property: ${error.message}`);
      } else {
        alert("Failed to sell property: An unexpected error occurred");
      }
    }
  };

  const handleUpgradeProperty = async (
    propertyID: number,
    pathName: string
  ) => {
    try {
      await performAction("upgrade_property", {
        property_id: propertyID,
        path_name: pathName,
      });
      onActionComplete();
    } catch (error) {
      if (error instanceof Error) {
        alert(`Failed to upgrade property: ${error.message}`);
      } else {
        alert("Failed to upgrade property: An unexpected error occurred");
      }
    }
  };

  return (
    <div className="property-list">
      <h2>Properties</h2>
      {properties.map((propertyEntity) => {
        const property = propertyEntity.Components.Property as Property;
        const owned = property.Owned;
        const propertyID = property.ID;

        return (
          <div key={propertyID} className="property-card">
            <h3>{property.Name}</h3>
            <p>
              <strong>Type:</strong> {property.Type}
            </p>
            <p>
              <strong>Subtype:</strong> {property.Subtype}
            </p>
            <p>
              <strong>Price:</strong> ${property.Price}
            </p>
            <p>
              <strong>Base Rent:</strong> ${property.BaseRent}
            </p>
            <p>
              <strong>Owned:</strong> {owned ? "Yes" : "No"}
            </p>
            {owned ? (
              <div className="owned-actions">
                <button onClick={() => handleSellProperty(propertyID)}>
                  Sell Property
                </button>
                <div className="upgrade-paths">
                  <h4>Upgrade Paths</h4>
                  {Object.keys(property.UpgradePaths).map((pathName) => (
                    <button
                      key={pathName}
                      onClick={() =>
                        handleUpgradeProperty(propertyID, pathName)
                      }
                    >
                      Upgrade {pathName}
                    </button>
                  ))}
                </div>
              </div>
            ) : (
              <button onClick={() => handleBuyProperty(propertyID)}>
                Buy Property
              </button>
            )}
          </div>
        );
      })}
    </div>
  );
};

export default PropertyList;
