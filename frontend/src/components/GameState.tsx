// src/components/GameState.tsx

import React from "react";
import { World, Entity, Player, GameTime } from "../types";

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

  const playerComponent = playerEntity.Components.Player as Player;
  const gameTimeComponent = gameTimeEntity.Components.GameTime as GameTime;

  const playerFunds = playerComponent.Funds.toFixed(2);
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
