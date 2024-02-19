import { useState } from "react";
import { Input } from "@/components/ui/input";

interface SendMessageProps {
  sendMessage: (message: string) => void;
}
export const SendMessage = ({ sendMessage }: SendMessageProps) => {
  const [text, setText] = useState("");

  const handleSendMessage = () => {
    sendMessage(text);
    setText("");
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        handleSendMessage();
      }}
    >
      <Input value={text} onChange={(e) => setText(e.target.value)} />
    </form>
  );
};
