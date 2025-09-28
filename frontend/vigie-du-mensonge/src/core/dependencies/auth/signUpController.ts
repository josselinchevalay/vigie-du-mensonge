import {toast} from "@/core/utils/toast.ts";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {navigate} from "@/core/utils/router.ts";
import {AuthClient} from "@/core/dependencies/auth/authClient.ts";

export class SignUpController {
    private readonly client: AuthClient;

    public readonly token: string | null;

    constructor(client: AuthClient, token: string | null) {
        this.client = client;
        this.token = token;
    }

    async onInquire(email: string): Promise<boolean> {
        try {
            await this.client.inquireSignUp(email);
            toast.success("L'e-mail contenant le lien d'inscription a été envoyé.");
            return true;
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            return false;
        }
    }

    async onProcess(username: string, password: string): Promise<boolean> {
        try {
            const auth = await this.client.processSignUp({token: this.token!, username: username, password: password});
            auth.saveToStorage();
            authManager.authStore.setState(() => auth);
            toast.success('Votre inscription est terminée!');
            void navigate({to: "/", replace: true});
            return true;
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            return false;
        }
    }
}


