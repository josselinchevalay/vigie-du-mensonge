import {Auth} from "@/core/models/auth";

type Deps = {
    getAuth: () => Auth | null;
    onRefresh: () => Promise<Auth | null> | Promise<void>;
};

export class AuthRefreshScheduler {
    private timer: ReturnType<typeof setTimeout> | null = null;
    private readonly deps: Deps;
    private readonly minDelayMs = 1_000;

    constructor(deps: Deps) {
        this.deps = deps;
        this.handleVisibility = this.handleVisibility.bind(this);
        this.handleOnline = this.handleOnline.bind(this);
        this.handleStorage = this.handleStorage.bind(this);
    }

    start(): void {
        this.reschedule();
        if (typeof document !== "undefined") {
            document.addEventListener("visibilitychange", this.handleVisibility);
        }
        if (typeof window !== "undefined") {
            window.addEventListener("online", this.handleOnline);
            window.addEventListener("storage", this.handleStorage);
        }
    }

    stop(): void {
        this.clearTimer();
        if (typeof document !== "undefined") {
            document.removeEventListener("visibilitychange", this.handleVisibility);
        }
        if (typeof window !== "undefined") {
            window.removeEventListener("online", this.handleOnline);
            window.removeEventListener("storage", this.handleStorage);
        }
    }

    private handleVisibility(): void {
        if (typeof document !== "undefined" && document.visibilityState === "visible") {
            this.checkAndMaybeRefresh();
        }
    }

    private handleOnline(): void {
        this.checkAndMaybeRefresh();
    }

    private handleStorage(e: StorageEvent): void {
        if (e.key === Auth.STORAGE_KEY) {
            this.reschedule();
        }
    }

    private timeUntil(date: Date): number {
        return date.getTime() - Date.now();
    }

    private checkAndMaybeRefresh(): void {
        const auth = this.deps.getAuth();
        if (!auth) {
            this.clearTimer();
            return;
        }

        const timeLeft = this.timeUntil(auth.accessTokenExpiry);
        if (auth.accessTokenExpired || timeLeft < this.minDelayMs) {
            void this.deps.onRefresh().finally(() => this.reschedule());
        } else {
            this.rescheduleFrom(auth);
        }
    }

    reschedule(): void {
        const auth = this.deps.getAuth();
        if (!auth) {
            this.clearTimer();
            return;
        }
        this.rescheduleFrom(auth);
    }

    private rescheduleFrom(auth: Auth): void {
        this.clearTimer();

        let delay = this.timeUntil(auth.accessTokenExpiry);

        if (delay <= 0) {
            this.timer = setTimeout(() => {
                void this.deps.onRefresh().finally(() => this.reschedule());
            }, this.minDelayMs);
            return;
        }

        delay = Math.max(delay, this.minDelayMs);

        this.timer = setTimeout(() => {
            void this.deps.onRefresh().finally(() => this.reschedule());
        }, delay);
    }

    private clearTimer(): void {
        if (this.timer) {
            clearTimeout(this.timer);
            this.timer = null;
        }
    }
}