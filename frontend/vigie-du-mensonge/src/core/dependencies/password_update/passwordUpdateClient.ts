import type {KyInstance} from "ky";

export class PasswordUpdateClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async inquire(email: string): Promise<void> {
        await this.api.post("password-update/inquire", {
            json: {email},
        });
    }

    async process(token: string, newPassword: string): Promise<void> {
        await this.api.post("password-update/process", {
            json: {token, newPassword},
        });
    }
}