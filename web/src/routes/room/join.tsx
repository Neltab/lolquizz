import { createFileRoute } from '@tanstack/react-router'
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { useLogin } from "@/lib/api/auth";
import { useNavigate } from "@tanstack/react-router";
import { Card, CardContent } from '../../components/ui/card';
import { PlayerSetup } from '../../components/ui/playerSetup';

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
  const { mutateAsync: login } = useLogin();
  const navigate = useNavigate()


  const handleJoinRoom = async (playerName: string, roomCode: string) => {
      const { token } = await login();
      gameState.setToken(token);
      gameState.setIsHost(false);
      gameState.setPlayerName(playerName);
      
      navigate({ to: '/room/$code', params: { code: roomCode } });
  }

  return (
    <div className="w-full h=full flex flex-col items-center justify-center grow">
      <Card>
        <CardContent>
          <PlayerSetup code={code} joinOnly={true} handleJoinRoom={handleJoinRoom} />
        </CardContent>
      </Card>
    </div>
  );
}
