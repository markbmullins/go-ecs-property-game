import React from "react";
import { Neighborhood, Property } from "../types";

interface NeighborhoodListProps {
  neighborhoods: Record<number, Neighborhood>;
  properties: Record<number, Property>;
}

const NeighborhoodList: React.FC<NeighborhoodListProps> = ({
  neighborhoods,
  properties,
}) => {
  return (
    <div className="neighborhood-list">
      {Object.values(neighborhoods).map((neighborhood) => (
        <div key={neighborhood.ID} className="neighborhood-card">
          <h3>{neighborhood.Name}</h3>
          <ul>
            {neighborhood.PropertyIDs.map((propertyID) => {
              const property = properties[propertyID];
              return (
                <li key={propertyID}>
                  {property.Name} - Base Rent: ${property.BaseRent}
                </li>
              );
            })}
          </ul>
        </div>
      ))}
    </div>
  );
};

export default NeighborhoodList;
