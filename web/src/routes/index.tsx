import { useNavigate, createFileRoute } from "@tanstack/react-router";
import { useCallback } from "react";
import { useCreateRoom } from "@/lib/api/room";
import { useLogin } from "@/lib/api/auth";
import { flushSync } from "react-dom";
import { Card, CardContent } from "../components/ui/card";
import { AvatarPicker } from "../components/ui/avatarPicker";
import { PlayerSetup } from "../components/ui/playerSetup";

export const Route = createFileRoute("/")({
	component: Index,
});

function Index() {
    const { gameState } = Route.useRouteContext();
    const { mutateAsync: login } = useLogin();
    const { mutateAsync: createRoom } = useCreateRoom();
    const navigate = useNavigate()


    const handleCreateRoom = useCallback(async (playerName: string) => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);
        const { code, is_host } = await createRoom({ token, playerName });
        flushSync(() => {
            gameState.setIsHost(is_host);
        })
        
        await navigate({ to: '/room/$code', params: { code } });
    }, [navigate, login])

    const handleJoinRoom = async ( playerName: string, code: string) => {
        const { token } = await login();
        gameState.setToken(token);
        gameState.setPlayerName(playerName);
        await navigate({ to: '/room/$code', params: { code: code } });
    }

	return (
		<div className="w-full h=full flex flex-col items-center justify-center grow">
            <Card>
                <CardContent>
                    <PlayerSetup handleCreateRoom={handleCreateRoom} handleJoinRoom={handleJoinRoom} />
                </CardContent>
            </Card>
		</div>
	)
}
