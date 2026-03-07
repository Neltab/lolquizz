import { createLazyFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { useCreateRoom } from "@/lib/api/room";

export const Route = createLazyFileRoute("/")({
	component: Index,
});

function Index() {

    const [nickname, setNickname] = useState("");
    const { mutateAsync: createRoom } = useCreateRoom();

    const handleCreateRoom = async () => {
        const data = await createRoom(nickname);
        console.log(data);
    }

	return (
		<div>
            <Input onChange={(e) => setNickname(e.target.value)} placeholder="Nickname" />
			<Button disabled={!nickname} onClick={handleCreateRoom}>Hello</Button>
		</div>
	);
}
