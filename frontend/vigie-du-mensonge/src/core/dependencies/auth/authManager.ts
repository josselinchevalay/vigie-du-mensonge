import {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";
import {AuthClient} from "@/core/dependencies/auth/authClient.ts";
import {toast} from "@/core/utils/toast";

class AuthManager {
    private readonly client = new AuthClient();
    public readonly authStore = new Store<Auth | null>(null);
    private refreshing = false;

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
        void this.client.signOut();
    }
}

export const authManager = new AuthManager();