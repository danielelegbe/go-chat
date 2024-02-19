import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useRef } from "react";
import { Message } from "./useMessages";
import { WS_URL } from "@/lib/contants";

export const useSubscription = () => {
  const queryClient = useQueryClient();
  const websocket = useRef<WebSocket | null>(null);

  useEffect(() => {
    websocket.current = new WebSocket(WS_URL);

    websocket.current.onmessage = (event) => {
      const data = JSON.parse(event.data) as Message;
      queryClient.setQueriesData({ queryKey: ["messages"] }, (oldData) => [
        ...(oldData as Message[]),
        data,
      ]);
    };

    return () => {
      websocket.current?.close();
    };
  }, [queryClient]);

  const sendMessage = (message: string) => {
    // Ensure the WebSocket is open before sending a message
    if (websocket.current?.readyState === WebSocket.OPEN) {
      websocket.current.send(message);
    } else {
      console.error("WebSocket is not open.");
    }
  };

  // Return the sendMessage function so it can be used by the component
  return { sendMessage };
};
