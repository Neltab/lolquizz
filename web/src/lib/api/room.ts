import { useMutation } from '@tanstack/react-query'

type CreateRoomParams = {
    token: string
    playerName: string
}

const createRoom = async ({token, playerName} : CreateRoomParams) => {
	const res = await fetch(`/api/rooms`, {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ token, playerName }),
	});
	const data = await res.json();
	if (res.ok) {
		console.log(data);
		return data;
	} else {
		throw new Error(data.message);
	}
};

export const useCreateRoom = () => useMutation({
    mutationFn: createRoom,
})