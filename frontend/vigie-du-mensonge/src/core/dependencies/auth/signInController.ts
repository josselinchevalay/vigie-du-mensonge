import {toast} from "@/core/utils/toast.ts";
import {navigate} from "@/core/utils/router.ts";
import {HTTPError} from "ky";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {authClient} from "@/core/dependencies/auth/authClient.ts";

export class SignInController {

    async onSignIn(email: string, password: string): Promise<boolean> {
        try {
            const auth = await authClient.signIn({email, password});
            auth.saveToStorage();
            authManager.authStore.setState(() => auth);
            void navigate({to: '/', replace: true});
            return true;
        } catch (e) {
            let status: number | undefined;
            if (e instanceof HTTPError) {
                status = e.response.status;
            }

            const message = status === 401
                ? "Identifiants invalides."
                : "Une erreur est survenue. Veuillez réessayer.";

            toast.error(message);
            return false;
        }
    }
}