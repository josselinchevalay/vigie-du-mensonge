import {api} from "@/core/dependencies/api";
import {Auth, type AuthJson} from "@/core/models/auth";

export type AuthCredentials = {
    email: string;
    password: string;
};

export class AuthClient {

    async signUp(credentials: AuthCredentials): Promise<Auth> {
        const res = await api
            .post("auth/sign-up", {
                json: credentials,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }

    async refresh(): Promise<Auth> {
        const res = await api
            .post("auth/refresh")
            .json<AuthJson>();

        return Auth.fromJson(res);
    }

    async signIn(credentials: AuthCredentials): Promise<Auth> {
        const res = await api
            .post("auth/sign-in", {
                json: credentials,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }
}
