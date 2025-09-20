import ky from "ky";

export const api = ky.create({
    // Default to '/api/v1' in tests or when VITE_API_URL is not provided
    prefixUrl: import.meta.env.VITE_API_URL ?? '/api/v1',
    credentials: "include",
    timeout: 10000,
    headers: {
        "Content-Type": "application/json",
    },
});
