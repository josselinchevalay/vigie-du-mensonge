import type {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";
import {AuthClient} from "@/core/dependencies/auth/auth-client.ts";

class AuthManager {
    private authClient = new AuthClient();

    authStore = new Store<Auth | null>(null);
    
    async signUp(credentials: { email: string; password: string }): Promise<Auth> {
        const auth = await this.authClient.signUp(credentials);
        // Persist to localStorage
        auth.saveToStorage();
        // Update store using the recommended API
        this.authStore.setState(() => auth);
        return auth;
    }
}

export const authManager = new AuthManager();