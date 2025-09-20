import {EmailVerificationClient} from "@/core/dependencies/email_verification/emailVerificationClient.ts";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {router} from "@/main.tsx";
import {toast} from "@/core/utils/toast";
import {Store} from "@tanstack/react-store";

export class EmailVerificationController {
    private readonly client = new EmailVerificationClient();
    public readonly tokenStore!: Store<string | null>;

    constructor(token: string | null) {
        this.tokenStore = new Store<string | null>(token);

        if (!token) {
            return;
        }

        if (authManager.refreshing) {
            const unsub = authManager.authStore.subscribe((listener) => {
                const auth = listener?.currentVal;
                if (!auth || auth.accessTokenExpired) return;
                unsub();
                void this.process();
            });
        } else {
            void this.process();
        }
    }

    async inquire() {
        await this.client.inquire();
    }

    async process() {
        const token = this.tokenStore.state;
        if (!token) {
            return;
        }

        try {
            await this.client.process(token);
            const auth = authManager.authStore.state!;
            if (auth) {
                auth.emailVerified = true;
                authManager.authStore.setState(() => auth);
                auth.saveToStorage();
                toast(
                    'Votre adresse email a été vérifiée, vous pouvez désormais contribuer à Vigie du mensonge! Merci!'
                );
                void router.navigate({to: '/', replace: true});
            }
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            this.tokenStore.setState(() => null);
        }
    }
}