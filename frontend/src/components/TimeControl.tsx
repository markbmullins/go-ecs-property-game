// src/components/TimeControl.tsx

import React, { useState } from "react";
import { performAction } from "../api";

interface TimeControlProps {
  onActionComplete: () => void;
}

const TimeControl: React.FC<TimeControlProps> = ({ onActionComplete }) => {
  const [speedMultiplier, setSpeedMultiplier] = useState<string>("1.0");

  const handlePause = async () => {
    try {
      await performAction("control_time", { action: "pause" });
      onActionComplete();
    } catch (error) {
      if (error instanceof Error) {
        alert(`Failed to pause time: ${error.message}`);
      } else {
        alert("Failed to pause time: An unexpected error occurred");
      }
    }
  };

  const handleStart = async () => {
    try {
      await performAction("control_time", { action: "start" });
      onActionComplete();
    } catch (error) {
      if (error instanceof Error) {
        alert(`Failed to start time: ${error.message}`);
      } else {
        alert("Failed to start time: An unexpected error occurred");
      }
    }
  };

  const handleSetSpeed = async () => {
    try {
      const speed = parseFloat(speedMultiplier);
      if (isNaN(speed) || speed <= 0) {
        alert("Please enter a valid speed multiplier greater than 0.");
        return;
      }
      await performAction("control_time", {
        action: "set_speed",
        speed_multiplier: speed,
      });
      onActionComplete();
    } catch (error) {
      if (error instanceof Error) {
        alert(`Failed to set speed: ${error.message}`);
      } else {
        alert("Failed to set speed: An unexpected error occurred");
      }
    }
  };

  return (
    <div className="time-control">
      <h2>Time Control</h2>
      <button onClick={handlePause}>Pause</button>
      <button onClick={handleStart}>Start</button>
      <div className="speed-control">
        <label>
          Speed Multiplier:
          <input
            type="number"
            step="0.1"
            min="0.1"
            value={speedMultiplier}
            onChange={(e) => setSpeedMultiplier(e.target.value)}
          />
        </label>
        <button onClick={handleSetSpeed}>Set Speed</button>
      </div>
    </div>
  );
};

export default TimeControl;
