import {api} from "@/core/dependencies/api.ts";

export class PasswordUpdateClient {
    async inquire(email: string): Promise<void> {
        await api.post("password-update/inquire", {
            json: {email},
        });
    }

    async process(token: string, newPassword: string): Promise<void> {
        await api.post("password-update/process", {
            json: {token, newPassword},
        });
    }
}