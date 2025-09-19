import {api} from "@/core/dependencies/api";
import {Auth, type AuthJson} from "@/core/models/auth";

export type AuthCredentials = {
    email: string;
    password: string;
};

export class AuthClient {
    // OpenAPI: POST /auth/sign-up
// Request: AuthCredentials { email, password }
// Response 201: SignUpResponse { accessTokenExpiry, refreshTokenExpiry }
    async signUp(credentials: AuthCredentials): Promise<Auth> {
        const res = await api
            .post("auth/sign-up", {
                json: credentials,
            })
            .json<AuthJson>();

        return Auth.fromJson(res);
    }
}
