import {describe, it, expect, vi, beforeEach, afterEach} from 'vitest';
import {AuthRefreshScheduler} from '@/core/dependencies/auth/authRefreshScheduler';
import {Auth} from '@/core/models/auth';

// Utility to control document.visibilityState for jsdom
function setDocumentVisibility(state: DocumentVisibilityState) {
  Object.defineProperty(document, 'visibilityState', {
    value: state,
    configurable: true,
  });
}

describe('AuthRefreshScheduler', () => {
  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.runOnlyPendingTimers();
    vi.useRealTimers();
  });

  it('does not schedule when no auth (getAuth returns null)', async () => {
    const onRefresh = vi.fn(async () => {});
    const scheduler = new AuthRefreshScheduler({
      getAuth: () => null,
      onRefresh,
    });

    scheduler.start();

    // Advance time generously; no refresh should happen
    vi.advanceTimersByTime(10_000);
    expect(onRefresh).not.toHaveBeenCalled();

    scheduler.stop();
  });

  it('schedules refresh at minDelay (1000ms) when token already expired', async () => {
    const onRefresh = vi.fn(async () => {});
    const expiredAuth = {
      accessTokenExpiry: new Date(Date.now() - 1_000),
      accessTokenExpired: true,
    } as unknown as Auth;

    const scheduler = new AuthRefreshScheduler({
      getAuth: () => expiredAuth,
      onRefresh,
    });

    scheduler.start();

    // First run happens after minDelay (1000ms)
    vi.advanceTimersByTime(999);
    expect(onRefresh).not.toHaveBeenCalled();

    vi.advanceTimersByTime(1);
    expect(onRefresh).toHaveBeenCalledTimes(1);

    // It reschedules again after .finally(() => reschedule())
    await vi.waitFor(() => expect(onRefresh).toHaveBeenCalledTimes(1));

    // Next call should occur again after at least 1000ms
    vi.advanceTimersByTime(1000);
    expect(onRefresh).toHaveBeenCalledTimes(2);

    scheduler.stop();
  });

  it('schedules using time until expiry (>= minDelay)', async () => {
    const onRefresh = vi.fn(async () => {});
    // Expiry in 1500ms from now
    const soonAuth = {
      accessTokenExpiry: new Date(Date.now() + 1_500),
      accessTokenExpired: false,
    } as unknown as Auth;

    const scheduler = new AuthRefreshScheduler({
      getAuth: () => soonAuth,
      onRefresh,
    });

    scheduler.start();

    // Before 1500ms: should not call refresh
    vi.advanceTimersByTime(1000);
    expect(onRefresh).not.toHaveBeenCalled();

    vi.advanceTimersByTime(500);
    expect(onRefresh).toHaveBeenCalledTimes(1);

    scheduler.stop();
  });

  it('visibilitychange to visible triggers immediate refresh when timeLeft < minDelay', async () => {
    const onRefresh = vi.fn(async () => {});
    // Expiry in 500ms (less than minDelay), not expired flag
    const nearAuth = {
      accessTokenExpiry: new Date(Date.now() + 500),
      accessTokenExpired: false,
    } as unknown as Auth;

    const scheduler = new AuthRefreshScheduler({
      getAuth: () => nearAuth,
      onRefresh,
    });

    scheduler.start();

    // Trigger visibility to visible -> should call onRefresh quickly
    setDocumentVisibility('visible');
    document.dispatchEvent(new Event('visibilitychange'));

    await vi.waitFor(() => expect(onRefresh).toHaveBeenCalledTimes(1));

    scheduler.stop();
  });

  it('online event triggers immediate refresh check (calls onRefresh if timeLeft < minDelay)', async () => {
    const onRefresh = vi.fn(async () => {});
    const nearAuth = {
      accessTokenExpiry: new Date(Date.now() + 500),
      accessTokenExpired: false,
    } as unknown as Auth;

    const scheduler = new AuthRefreshScheduler({
      getAuth: () => nearAuth,
      onRefresh,
    });

    scheduler.start();

    window.dispatchEvent(new Event('online'));

    await vi.waitFor(() => expect(onRefresh).toHaveBeenCalledTimes(1));

    scheduler.stop();
  });

  it('storage event on Auth.STORAGE_KEY reschedules with latest auth value', async () => {
    const onRefresh = vi.fn(async () => {});

    // Start with expiry far in future (5s)
    let currentAuth: Auth = {
      accessTokenExpiry: new Date(Date.now() + 5_000),
      accessTokenExpired: false,
    } as unknown as Auth;

    const getAuth = vi.fn(() => currentAuth);

    const scheduler = new AuthRefreshScheduler({
      getAuth,
      onRefresh,
    });

    scheduler.start();

    // After 1s, not yet refreshed
    vi.advanceTimersByTime(1_000);
    expect(onRefresh).not.toHaveBeenCalled();

    // Change auth to expire soon and dispatch storage to reschedule
    currentAuth = {
      accessTokenExpiry: new Date(Date.now() + 500), // less than minDelay
      accessTokenExpired: false,
    } as unknown as Auth;

    window.dispatchEvent(new StorageEvent('storage', {key: Auth.STORAGE_KEY}));

    // Because timeLeft < minDelay, the rescheduleFrom will set a 1000ms timeout
    vi.advanceTimersByTime(999);
    expect(onRefresh).not.toHaveBeenCalled();

    vi.advanceTimersByTime(1);
    expect(onRefresh).toHaveBeenCalledTimes(1);

    scheduler.stop();
  });
});
