import { RouterProvider, createRouter } from "@tanstack/react-router";
import { useState } from "react";
import { routeTree } from "./routeTree.gen";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();
const router = createRouter({
  routeTree,
  context: undefined!,
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

export interface GameState {
  token: string | null;
  isHost: boolean;
  nickname: string;
  setToken: (token: string | null) => void;
  setIsHost: (isHost: boolean) => void;
  setNickname: (nickname: string) => void;
}

export default function App() {
  const [token, setToken] = useState<string | null>(null);
  const [isHost, setIsHost] = useState(false);
  const [nickname, setNickname] = useState("");

  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider
        router={router}
        context={{
          queryClient,
          gameState: { token, isHost, nickname, setToken, setIsHost, setNickname },
        }}
      />
    </QueryClientProvider>
  );
}