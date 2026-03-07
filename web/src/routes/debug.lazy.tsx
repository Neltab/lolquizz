import { createLazyFileRoute } from "@tanstack/react-router";
import WsDebug from "../components/ws-debug";


export const Route = createLazyFileRoute("/debug")({
	component: Debug,
});

function Debug() {
	return (
		<div>
			<WsDebug />
		</div>
	)
}
