import {api} from "@/core/dependencies/api";
import {Auth, type AuthJson} from "@/core/models/auth";

export class AuthClient {
    async refresh(): Promise<Auth> {
        const res = await api
            .post("auth/refresh")
            .json<AuthJson>();

        return Auth.fromJson(res);
    }

    async signOut(): Promise<void> {
        await api.post("auth/sign-out");
    }
}
