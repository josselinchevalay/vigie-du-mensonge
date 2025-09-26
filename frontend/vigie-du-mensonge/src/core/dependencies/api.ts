import ky from "ky";

const isTest = Boolean(import.meta.env?.MODE === "test" || import.meta?.env?.VITEST);

export const api = ky.create({
    prefixUrl: import.meta.env.VITE_API_URL ?? 'http://localhost:8080/api/v1',
    credentials: "include",
    timeout: 10000,
    headers: {
        "Content-Type": "application/json",
    },
    retry: 0, // let tanstack query handle retries
    hooks: {
        beforeRequest: isTest ? [] : [
            async (request) => {
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
        .get("csrf-token", {retry: 2}) // do not use ky retries except for csrf token
        .json<{ csrfToken: string }>();
    return response.csrfToken;
};



