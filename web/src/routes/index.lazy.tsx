import { createLazyFileRoute, useNavigate } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { useCreateRoom } from "@/lib/api/room";
import { useLogin } from "@/lib/api/auth";

export const Route = createLazyFileRoute("/")({
	component: Index,
});

function Index() {
    const { gameState } = Route.useRouteContext();
    const [playerName, setplayerName] = useState("");
    const [roomCode, setRoomCode] = useState("");
    const { mutateAsync: login } = useLogin();
    const { mutateAsync: createRoom } = useCreateRoom();
    const navigate = useNavigate()


    const handleCreateRoom = async () => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);

        const { code, is_host } = await createRoom({ token, playerName });
        gameState.setIsHost(is_host);
        console.log(gameState)
        
        await navigate({ to: '/room/$code', params: { code } });
    }

    const handleJoinRoom = async () => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);
        console.log(gameState)
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
	);
}
