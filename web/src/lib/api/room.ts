import { useMutation } from '@tanstack/react-query'

type CreateRoomParams = {
    token: string
    nickname: string
}

const createRoom = async ({token, nickname} : CreateRoomParams) => {
	const res = await fetch(`/api/rooms`, {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ token, nickname }),
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