import type {Auth} from "@/core/models/auth.ts";
import {Store} from "@tanstack/react-store";

class AuthManager {
    authStore = new Store<Auth | null>(null);
}

export const authManager = new AuthManager();