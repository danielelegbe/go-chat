import { Messages } from "@/components/messages";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="space-y-2 p-4">
        <Messages />
      </div>
    </QueryClientProvider>
  );
}

export default App;
