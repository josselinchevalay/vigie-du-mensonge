import {Auth, type AuthJson} from "@/core/models/auth.ts";
import {api} from "@/core/dependencies/api.ts";

export class SignInClient {
    async signIn(creds: { email: string, password: string }): Promise<Auth> {
        const res = await api
            .post("auth/sign-in", {
                json: creds,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }
}