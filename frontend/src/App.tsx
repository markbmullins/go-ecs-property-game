// src/App.tsx

import React, { useEffect, useState } from "react";
import { fetchGameState } from "./api";
import { World } from "./types";
import GameState from "./components/GameState";
import PropertyList from "./components/PropertyList";
import TimeControl from "./components/TimeControl";
import "./App.css";

const App: React.FC = () => {
  const [gameState, setGameState] = useState<World | null>(null);
  const [error, setError] = useState<string | null>(null);

  const loadGameState = async () => {
    try {
      const state = await fetchGameState();
      setGameState(state);
      setError(null);
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError("An unexpected error occurred");
      }
    }
  };

  useEffect(() => {
    loadGameState();
    const interval = setInterval(loadGameState, 2000); // Refresh every 2 seconds
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="App">
      <h1>City Developer Game</h1>
      {error && <p className="error">Error: {error}</p>}
      {gameState ? (
        <>
          <GameState gameState={gameState} />
          <TimeControl onActionComplete={loadGameState} />
          <PropertyList
            gameState={gameState}
            onActionComplete={loadGameState}
          />
        </>
      ) : (
        !error && <p>Loading game state...</p>
      )}
    </div>
  );
};

export default App;
