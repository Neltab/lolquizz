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
    const { mutateAsync: login } = useLogin();
    const { mutateAsync: createRoom } = useCreateRoom();
    const navigate = useNavigate()


    const handleCreateRoom = async () => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);

        const { code, is_host } = await createRoom({ token, playerName });
        gameState.setIsHost(is_host);
        
        navigate({ to: '/room/$code', params: { code } });
    }

	return (
		<div>
            <Input onChange={(e) => setplayerName(e.target.value)} placeholder="playerName" />
			<Button disabled={!playerName} onClick={handleCreateRoom}>Create</Button>
		</div>
	);
}
