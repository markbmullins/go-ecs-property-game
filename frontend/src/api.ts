// src/api.ts

import { World } from "./types";

const API_BASE_URL = "http://localhost:8080";

/**
 * Helper function to handle fetch responses.
 * Throws an error if the response is not ok.
 */
const handleResponse = async <T>(response: Response): Promise<T> => {
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || "Unknown error occurred");
  }
  return response.json();
};

/**
 * Fetches the current game state from the server.
 */
export const fetchGameState = async (): Promise<World> => {
  try {
    const response = await fetch(`${API_BASE_URL}/state`);
    const data = await handleResponse<World>(response);
    return data;
  } catch (error) {
    console.error("Error fetching game state:", error);
    throw error;
  }
};

/**
 * Performs an action by sending a POST request to the server.
 * @param action - The action name.
 * @param payload - The payload for the action.
 */
export interface ActionRequest {
  action: string;
  payload: any;
}

export const performAction = async (
  action: string,
  payload: any
): Promise<any> => {
  try {
    const response = await fetch(`${API_BASE_URL}/actions`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        action,
        payload,
      }),
    });
    const data = await handleResponse<any>(response);
    return data;
  } catch (error) {
    console.error(`Error performing action ${action}:`, error);
    throw error;
  }
};
