import {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";
import {AuthClient} from "@/core/dependencies/auth/authClient.ts";
import {toast} from "sonner";

class AuthManager {
    private readonly client = new AuthClient();
    private refreshing = false;

    authStore = new Store<Auth | null>(null);

    init() {
        const stored = Auth.fromStorage();
        this.authStore.setState(() => stored);
    }

    async signIn(credentials: { email: string; password: string }): Promise<Auth> {
        const auth = await this.client.signIn(credentials);

        auth.saveToStorage();

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
        if (this.refreshing) {
            return this.authStore.state;
        }
        this.refreshing = true;

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
            this.refreshing = false;
        }
    }

    signOut() {
        Auth.clearStorage();
        this.authStore.setState(() => null);
        this.client.signOut();
    }
}

export const authManager = new AuthManager();