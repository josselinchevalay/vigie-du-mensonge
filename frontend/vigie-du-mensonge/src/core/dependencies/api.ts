import ky from "ky";

const isTest = Boolean(import.meta.env?.MODE === "test" || import.meta?.env?.VITEST);

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
                if (isTest) return;

                if (request.method !== "POST" &&
                    request.method !== "PUT" &&
                    request.method !== "DELETE" &&
                    request.method !== "PATCH") {
                    return;
                }

                const csrfToken = await fetchCSRF();
                request.headers.set("X-Csrf-Token", csrfToken);
            }
        ]
    }
});

const fetchCSRF = async (): Promise<string> => {
    const response = await api
        .get("csrf-token", {retry: 2})
        .json<{ csrfToken: string }>();
    return response.csrfToken;
};