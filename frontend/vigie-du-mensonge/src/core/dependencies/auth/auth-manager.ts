import {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";
import {AuthClient} from "@/core/dependencies/auth/auth-client.ts";
import {toast} from "sonner";

class AuthManager {
    private readonly authClient = new AuthClient();
    private refreshing = false;

    authStore = new Store<Auth | null>(null);

    init() {
        const stored = Auth.fromStorage();
        this.authStore.setState(() => stored);
    }

    async signIn(credentials: { email: string; password: string }): Promise<Auth> {
        const auth = await this.authClient.signIn(credentials);

        auth.saveToStorage();

        this.authStore.setState(() => auth);
        return auth;
    }

    async signUp(credentials: { email: string; password: string }): Promise<Auth> {
        const auth = await this.authClient.signUp(credentials);

        auth.saveToStorage();

        this.authStore.setState(() => auth);
        return auth;
    }

    async refresh(): Promise<Auth | null> {
        if (this.refreshing) {
            return this.authStore.state;
        }
        this.refreshing = true;

        try {
            const freshAuth = await this.authClient.refresh();
            freshAuth.saveToStorage();

            this.authStore.setState(() => freshAuth);
            return freshAuth;
        } catch {
            this.authStore.setState(() => null);
            try {
                localStorage.removeItem(Auth.STORAGE_KEY);
            } catch {
                /* ignore storage removal errors */
            }
            toast('Votre session a expiré.');
            return null;
        } finally {
            this.refreshing = false;
        }
    }
}

export const authManager = new AuthManager();