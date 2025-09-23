import '@testing-library/jest-dom';
import { afterAll, afterEach, beforeAll, vi } from 'vitest';
import { server } from './testServer';
import { cleanup } from '@testing-library/react';

// Load global mocks before any test files
import './mocks';

beforeAll(() => server.listen({ onUnhandledRequest: 'error' }));
afterEach(() => {
    server.resetHandlers();
    cleanup();
    vi.clearAllMocks();
});
afterAll(() => server.close());