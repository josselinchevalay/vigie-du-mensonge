import {Auth, type AuthJson} from "@/core/models/auth.ts";
import {api} from "@/core/dependencies/api.ts";

export class SignUpClient {
    async inquireSignUp(email: string): Promise<void> {
        await api.post("auth/sign-up/inquire", {
            json: {email},
        });
    }

    async processSignUp(creds: { token: string, password: string }): Promise<Auth> {
        const res = await api
            .post("auth/sign-up/process", {
                json: creds,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }
}