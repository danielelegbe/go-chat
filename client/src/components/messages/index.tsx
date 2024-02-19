import { useMessages } from "@/hooks/useMessages";
import { useSubscription } from "@/hooks/useSubscription";
import { SendMessage } from "./send-message";

export const Messages = () => {
  const { data: messages, error, isPending } = useMessages();
  const { sendMessage } = useSubscription();

  if (isPending) {
    return <div>Loading...</div>;
  }
  return (
    <>
      <h1 className="text-2xl font-semibold">Messages</h1>
      <div>
        {error && <div>Error: {error.message}</div>}
        <ul className="space-y-2 mb-4">
          {messages?.map((message) => (
            <li key={message.id} className="bg-gray-100 p-2 rounded">
              {message.content}
            </li>
          ))}
        </ul>
        <SendMessage sendMessage={sendMessage} />
      </div>
    </>
  );
};
