import {PasswordUpdateClient} from "@/core/dependencies/password_update/passwordUpdateClient.ts";
import {Store} from "@tanstack/react-store";
import {toast} from "@/core/utils/toast.ts";

export class PasswordUpdateController {
    private readonly client = new PasswordUpdateClient();
    public readonly tokenStore!: Store<string | null>;

    constructor(token: string | null) {
        this.tokenStore = new Store(token);
    }

    public async onInquire(email: string): Promise<boolean> {
        try {
            await this.inquire(email);
            toast("L'email de modification a été envoyé");
            return true;
        } catch {
            toast("Une erreur est survenue. Veuillez réessayer.");
            return false;
        }
    }

    private async inquire(email: string) {
        await this.client.inquire(email);
    }

    public async onProcess(newPassword: string): Promise<boolean> {
        try {
            await this.process(newPassword);
            toast('Votre mot de passe a été mis à jour.');
            return true;
        } catch {
            toast("Une erreur est survenue. Veuillez réessayer.");
            return false;
        }
    }

    private async process(newPassword: string) {
        await this.client.process(this.tokenStore.state!, newPassword);
    }
}