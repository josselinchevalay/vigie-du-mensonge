import {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";
import {authClient, AuthClient} from "@/core/dependencies/auth/authClient.ts";
import {toast} from "@/core/utils/toast";
import {queryClient} from "@/core/dependencies/queryClient.ts";

class AuthManager {
    private readonly client: AuthClient;
    private refreshing = false;

    public readonly authStore = new Store<Auth | null>(null);

    private handleStorage = (e: StorageEvent) => {
        if (e.key !== Auth.STORAGE_KEY) return;
        // Pull from storage to avoid trusting e.newValue shape
        const auth = Auth.fromStorage();
        if (!auth || auth.refreshTokenExpired) {
            this.authStore.setState(() => null);
            return;
        }
        this.authStore.setState(() => auth);
    };

    constructor(client: AuthClient) {
        this.client = client;

        const stored = Auth.fromStorage();

        if (!stored || stored.refreshTokenExpired) {
            this.authStore.setState(() => null);
            return;
        } else {
            this.authStore.setState(() => stored);

            if (stored.accessTokenExpired) {
                void this.refresh();
            }
        }

        if (typeof window !== "undefined") {
            window.addEventListener("storage", this.handleStorage);
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

    signOut = () => {
        Auth.clearStorage();
        this.authStore.setState(() => null);
        void this.client.signOut();

        queryClient.removeQueries({
            predicate: (query) => {
                const key = query.queryKey[0];
                return key === "redactor" || key === "moderator" || key === "admin";
            },
        });
    };
}

export const authManager = new AuthManager(authClient);