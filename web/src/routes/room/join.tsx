import { createFileRoute } from '@tanstack/react-router'
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { useLogin } from "@/lib/api/auth";
import { useNavigate } from "@tanstack/react-router";

type JoinPageSearch = {
  code: string,
}

export const Route = createFileRoute('/room/join')({
  component: RouteComponent,
  validateSearch: (search: Record<string, unknown>): JoinPageSearch => {
    return {
      code: search.code as string || '',
    }
  },
})

function RouteComponent() {
  const { code } = Route.useSearch()
  const { gameState } = Route.useRouteContext();
  const [nickname, setNickname] = useState("");
  const { mutateAsync: login } = useLogin();
  const navigate = useNavigate()


  const handleJoinRoom = async () => {
      const { token } = await login();
      gameState.setToken(token);
      gameState.setIsHost(false);
      gameState.setNickname(nickname);
      
      navigate({ to: '/room/$code', params: { code } });
  }

  return (
    <div>
            <Input onChange={(e) => setNickname(e.target.value)} placeholder="Nickname" />
      <Button disabled={!nickname} onClick={handleJoinRoom}>Join</Button>
    </div>
  );
}
