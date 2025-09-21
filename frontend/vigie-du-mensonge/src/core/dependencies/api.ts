import ky from "ky";

export const api = ky.create({
    // Default to '/api/v1' in tests or when VITE_API_URL is not provided
    prefixUrl: import.meta.env.VITE_API_URL ?? '/api/v1',
    credentials: "include",
    timeout: 10000,
    headers: {
        "Content-Type": "application/json",
    },
    hooks: {
        beforeRequest: [
            async (request) => {
                if (request.method !== "POST" &&
                    request.method !== "PUT" &&
                    request.method !== "DELETE" &&
                    request.method !== "PATCH") {
                    return;
                }

                const csrfToken = await fetchCSRF();
                console.log("CSRF Token:", csrfToken);
                request.headers.set("X-Csrf-Token", csrfToken);
            }
        ]
    }
});

const fetchCSRF = async (): Promise<string> => {
    const response = await api.get("csrf-token").json<{ csrfToken: string }>();
    return response.csrfToken;
};