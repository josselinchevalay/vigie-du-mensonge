import { setupServer } from 'msw/node';
import { http, HttpResponse } from 'msw';

export const handlers = [
    // Example handler
    http.get('/api/v1/something', () => {
        return HttpResponse.json({ items: [{ id: 1, name: 'alpha' }] });
    }),

    // Auth: sign-in success (match both relative and absolute URLs)
    http.post('/api/v1/auth/sign-in', async () => {
        const now = Date.now();
        const oneHour = 60 * 60 * 1000;
        const twoHours = 2 * oneHour;
        return HttpResponse.json({
            accessTokenExpiry: new Date(now + oneHour).toISOString(),
            refreshTokenExpiry: new Date(now + twoHours).toISOString(),
            emailVerified: true,
            roles: [],
        });
    }),
    http.post('http://localhost:8080/api/v1/auth/sign-in', async () => {
        const now = Date.now();
        const oneHour = 60 * 60 * 1000;
        const twoHours = 2 * oneHour;
        return HttpResponse.json({
            accessTokenExpiry: new Date(now + oneHour).toISOString(),
            refreshTokenExpiry: new Date(now + twoHours).toISOString(),
            emailVerified: true,
            roles: [],
        });
    }),

    // Auth: sign-up success (match both relative and absolute URLs)
    http.post('/api/v1/auth/sign-up', async () => {
        const now = Date.now();
        const oneHour = 60 * 60 * 1000;
        const twoHours = 2 * oneHour;
        return HttpResponse.json({
            accessTokenExpiry: new Date(now + oneHour).toISOString(),
            refreshTokenExpiry: new Date(now + twoHours).toISOString(),
            emailVerified: true,
            roles: [],
        });
    }),
    http.post('http://localhost:8080/api/v1/auth/sign-up', async () => {
        const now = Date.now();
        const oneHour = 60 * 60 * 1000;
        const twoHours = 2 * oneHour;
        return HttpResponse.json({
            accessTokenExpiry: new Date(now + oneHour).toISOString(),
            refreshTokenExpiry: new Date(now + twoHours).toISOString(),
            emailVerified: true,
            roles: [],
        });
    }),
];

export const server = setupServer(...handlers);