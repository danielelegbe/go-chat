import { API_URL } from "@/lib/contants";
import { useQuery } from "@tanstack/react-query";

export interface Message {
  id: string;
  content: string;
}

export const useMessages = () => {
  const { isPending, error, data } = useQuery<Message[]>({
    queryKey: ["messages"],
    queryFn: () => fetch(`${API_URL}/messages`).then((res) => res.json()),
  });

  return { isPending, error, data };
};
