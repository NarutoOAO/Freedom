import React, { createContext, useContext, useReducer, useEffect } from "react";
import WebSocketService from "../WebSocketService";

const WebSocketContext = createContext();

const initialState = {
  messages: [],
};

const reducer = (state, action) => {
  switch (action.type) {
    case "UPDATE_MESSAGES":
      return {
        ...state,
        messages: action.payload,
      };
    case "CLEAR_MESSAGES":
      return {
        ...state,
        messages: [],
      };
    default:
      return state;
  }
};

export const WebSocketProvider = ({ children }) => {
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    console.log("I'm connecting!");
    // Initialize WebSocket connection when the component is mounted
    WebSocketService.init((message) => {
      // Handle incoming WebSocket messages here
      dispatch({ type: "UPDATE_MESSAGES", payload: message });
    });

    return () => {
      // Cleanup function when component unmounts or disconnects
      WebSocketService.close();
      dispatch({ type: "CLEAR_MESSAGES" });
    };
  }, []);

  return (
    <WebSocketContext.Provider value={{ state, dispatch }}>
      {children}
    </WebSocketContext.Provider>
  );
};

export const useWebSocket = () => {
  return useContext(WebSocketContext);
};
