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
    const [nickname, setNickname] = useState("");
    const { mutateAsync: login } = useLogin();
    const { mutateAsync: createRoom } = useCreateRoom();
    const navigate = useNavigate()


    const handleCreateRoom = async () => {
        const { token } = await login();
        gameState.setToken(token);

        const { code, is_host } = await createRoom({ token, nickname });
        gameState.setIsHost(is_host);
        
        navigate({ to: '/room/$code', params: { code } });
    }

	return (
		<div>
            <Input onChange={(e) => setNickname(e.target.value)} placeholder="Nickname" />
			<Button disabled={!nickname} onClick={handleCreateRoom}>Create</Button>
		</div>
	);
}
