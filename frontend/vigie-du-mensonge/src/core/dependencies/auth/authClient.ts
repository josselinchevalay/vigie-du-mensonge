import {Auth, type AuthJson} from "@/core/models/auth";
import type {KyInstance} from "ky";
import {api} from "@/core/dependencies/api.ts";

export class AuthClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async signIn(creds: { email: string, password: string }): Promise<Auth> {
        const res = await this.api
            .post("sign-in", {
                json: creds,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }

    async refresh(): Promise<Auth> {
        const res = await this.api
            .post("refresh")
            .json<AuthJson>();

        return Auth.fromJson(res);
    }

    async signOut(): Promise<void> {
        await this.api.post("sign-out");
    }

    async inquireSignUp(email: string): Promise<void> {
        await this.api.post("sign-up/inquire", {
            json: {email},
        });
    }

    async processSignUp(creds: { token: string, password: string }): Promise<Auth> {
        const res = await this.api
            .post("sign-up/process", {
                json: creds,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }
}

export const authClient = new AuthClient(api.extend({prefixUrl: "/auth"}));