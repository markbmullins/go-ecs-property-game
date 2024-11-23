// src/components/GameState.tsx

import React from "react";
import { World, Entity, PlayerComponent, GameTimeComponent } from "../types";

interface GameStateProps {
  gameState: World;
}

const GameState: React.FC<GameStateProps> = ({ gameState }) => {
  const entities = Object.values(gameState.Entities);

  const playerEntity = entities.find(
    (entity) => "PlayerComponent" in entity.Components
  ) as Entity | undefined;
  const gameTimeEntity = entities.find(
    (entity) => "GameTime" in entity.Components
  ) as Entity | undefined;

  if (!playerEntity || !gameTimeEntity) {
    return <p>Player or Game Time information is missing.</p>;
  }

  const playerComponent = playerEntity.Components
    .PlayerComponent as PlayerComponent;
  const gameTimeComponent = (
    gameTimeEntity.Components.GameTime as GameTimeComponent
  ).Time;

  const playerFunds = playerComponent.Player.Funds.toFixed(2);
  const currentDate = new Date(
    gameTimeComponent.CurrentDate
  ).toLocaleDateString();

  console.log({ gameTimeComponent, currentDate, gameState });
  return (
    <div className="game-state">
      <h2>Game State</h2>
      <p>
        <strong>Player Funds:</strong> ${playerFunds}
      </p>
      <p>
        <strong>Current Date:</strong> {currentDate}
      </p>
    </div>
  );
};

export default GameState;
