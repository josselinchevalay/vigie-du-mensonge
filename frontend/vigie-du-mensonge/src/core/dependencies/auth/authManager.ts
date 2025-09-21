import {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";
import {AuthClient} from "@/core/dependencies/auth/authClient.ts";
import {toast} from "@/core/utils/toast";

const EMAIL_STORAGE_KEY = 'vdm_email';

class AuthManager {
    private readonly client = new AuthClient();
    private _refreshing = false;

    public get email(): string | null {
        return localStorage.getItem(EMAIL_STORAGE_KEY);
    }

    public get refreshing(): boolean {
        return this._refreshing;
    }

    public readonly authStore = new Store<Auth | null>(null);

    constructor() {
        const stored = Auth.fromStorage();

        if (!stored || stored.refreshTokenExpired) {
            this.authStore.setState(() => null);
            return;
        }

        this.authStore.setState(() => stored);

        if (stored.accessTokenExpired) {
            void this.refresh();
        }
    }

    async signIn(credentials: { email: string; password: string }): Promise<Auth> {
        const auth = await this.client.signIn(credentials);

        auth.saveToStorage();
        localStorage.setItem(EMAIL_STORAGE_KEY, credentials.email);

        this.authStore.setState(() => auth);
        return auth;
    }

    async signUp(credentials: { email: string; password: string }): Promise<Auth> {
        const auth = await this.client.signUp(credentials);

        auth.saveToStorage();

        this.authStore.setState(() => auth);
        return auth;
    }

    async refresh(): Promise<Auth | null> {
        if (this._refreshing) {
            return this.authStore.state;
        }
        this._refreshing = true;

        try {
            const freshAuth = await this.client.refresh();
            freshAuth.saveToStorage();

            this.authStore.setState(() => freshAuth);
            return freshAuth;
        } catch {
            this.authStore.setState(() => null);
            Auth.clearStorage();
            toast('Votre session a expiré.');
            return null;
        } finally {
            this._refreshing = false;
        }
    }

    signOut() {
        Auth.clearStorage();
        this.authStore.setState(() => null);
        void this.client.signOut();
    }
}

export const authManager = new AuthManager();