import { useNavigate, createFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState, useCallback } from "react";
import { useCreateRoom } from "@/lib/api/room";
import { useLogin } from "@/lib/api/auth";
import { flushSync } from "react-dom";

export const Route = createFileRoute("/")({
	component: Index,
});

function Index() {
    const { gameState } = Route.useRouteContext();
    const [playerName, setplayerName] = useState("");
    const [roomCode, setRoomCode] = useState("");
    const { mutateAsync: login } = useLogin();
    const { mutateAsync: createRoom } = useCreateRoom();
    const navigate = useNavigate()


    const handleCreateRoom = useCallback(async () => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);
        const { code, is_host } = await createRoom({ token, playerName });
        flushSync(() => {
            gameState.setIsHost(is_host);
        })
        
        await navigate({ to: '/room/$code', params: { code } });
    }, [playerName, navigate, login])

    const handleJoinRoom = async () => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);
        await navigate({ to: '/room/$code', params: { code: roomCode } });
    }

    const roomCodeValid = roomCode.length === 6;

	return (
		<div className="w-full h=full flex flex-col items-center justify-center grow">
            <div className="flex flex-col gap-4">
                <Input onChange={(e) => setplayerName(e.target.value)} placeholder="Pseudo" />
                <div className="flex items-center justify-center gap-4">
                    <Input onChange={(e) => setRoomCode(e.target.value)} placeholder="Code de la partie" />
                    <Button disabled={!roomCodeValid} onClick={handleJoinRoom}>Rejoindre</Button>
                </div>
                <Button className="w-full" disabled={!playerName} onClick={handleCreateRoom}>Créer une partie privée</Button>
            </div>
		</div>
	)
}
