import { useMutation } from '@tanstack/react-query'

const createRoom = async (nickname: string) => {
	const res = await fetch(`/api/rooms`, {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ nickname }),
	});
	const data = await res.json();
	if (res.ok) {
		return data;
	} else {
		throw new Error(data.message);
	}
};

export const useCreateRoom = () => useMutation({
    mutationFn: createRoom,
})