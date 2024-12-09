// src/components/PropertyList.tsx

import React from "react";
import { performAction } from "../api";
import { World, Property, Neighborhood } from "../types";

interface PropertyListProps {
  gameState: World;
  onActionComplete: () => void;
}

const PropertyList: React.FC<PropertyListProps> = ({
  gameState,
  onActionComplete,
}) => {
  const entities = Object.values(gameState.Entities);

  // Extract neighborhoods and properties
  const neighborhoods = entities
    .filter((entity) => "Neighborhood" in entity.Components)
    .map((entity) => entity.Components.Neighborhood as Neighborhood);
  const properties = entities
    .filter((entity) => "Property" in entity.Components)
    .map((entity) => entity.Components.Property as Property);

  // Group properties by neighborhood
  const groupedProperties = neighborhoods.map((neighborhood) => {
    const neighborhoodProperties = properties.filter(
      (property) => property.NeighborhoodID === neighborhood.ID
    );
    return { neighborhood, properties: neighborhoodProperties };
  });

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
      <h2>Properties by Neighborhood</h2>
      {groupedProperties.map(({ neighborhood, properties }) => (
        <div key={neighborhood.ID} className="neighborhood-section">
          <h3>{neighborhood.Name}</h3>
          <p>
            <strong>Average Property Value:</strong> $
            {neighborhood.AveragePropertyValue.toLocaleString()}
          </p>
          <p>
            <strong>Rent Boost Threshold:</strong>{" "}
            {neighborhood.RentBoostThreshold}%
          </p>
          <p>
            <strong>Rent Boost Percent:</strong> {neighborhood.RentBoostPercent}
            %
          </p>
          <div className="property-grid">
            {properties.length > 0 ? (
              properties.map((property) => (
                <div key={property.ID} className="property-card">
                  <h4>{property.Name}</h4>
                  <p>
                    <strong>Type:</strong> {property.Type}
                  </p>
                  <p>
                    <strong>Subtype:</strong> {property.Subtype}
                  </p>
                  <p>
                    <strong>Price:</strong> ${property.Price.toLocaleString()}
                  </p>
                  <p>
                    <strong>Base Rent:</strong> $
                    {property.BaseRent.toLocaleString()}
                  </p>
                  <p>
                    <strong>Owned:</strong> {property.Owned ? "Yes" : "No"}
                  </p>
                  {property.Owned ? (
                    <div className="owned-actions">
                      <button onClick={() => handleSellProperty(property.ID)}>
                        Sell Property
                      </button>
                      <div className="upgrade-paths">
                        <h4>Upgrade Paths</h4>
                        {Object.keys(property.UpgradePaths).map((pathName) => (
                          <button
                            key={pathName}
                            onClick={() =>
                              handleUpgradeProperty(property.ID, pathName)
                            }
                          >
                            Upgrade {pathName}
                          </button>
                        ))}
                      </div>
                    </div>
                  ) : (
                    <button onClick={() => handleBuyProperty(property.ID)}>
                      Buy Property
                    </button>
                  )}
                </div>
              ))
            ) : (
              <p>No properties available in this neighborhood.</p>
            )}
          </div>
        </div>
      ))}
    </div>
  );
};

export default PropertyList;
