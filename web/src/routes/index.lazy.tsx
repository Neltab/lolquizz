import { createLazyFileRoute } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import WsDebug from "../../components/ws-debug";

export const Route = createLazyFileRoute("/")({
	component: Index,
});

function Index() {
	const [message, setMessage] = useState("");
    const [socket, setSocket] = useState<WebSocket | null>(null);

    useEffect(() => {
        const socket = new WebSocket("ws://localhost:8080/ws")
        setSocket(socket)

        // Connection opened
        socket.addEventListener("open", () => {
            socket.send("Connection established")
        });
    
        // Listen for messages
        socket.addEventListener("message", event => {
            console.log("Message from server ", event.data)
        });
	}, []);


	useEffect(() => {
		fetch("/api/ping")
			.then((res) => res.text())
			.then((data) => setMessage(data))
			.catch((err) => console.error(err));
	}, []);

    const handleClick = () => {
        socket?.send("Hello from client")
    }

	return (
		<div>
			<WsDebug />
		</div>
	);
}
