import { useMutation } from '@tanstack/react-query'

const login = async () => {
	const res = await fetch(`/api/auth/login`, {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ }),
	});
	const data = await res.json();
	if (res.ok) {
		return data;
	} else {
		throw new Error(data.message);
	}
};

export const useLogin = () => useMutation({
    mutationFn: login,
})