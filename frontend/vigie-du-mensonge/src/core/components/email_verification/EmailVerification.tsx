import {useState} from 'react';
import {EmailVerificationController} from '@/core/dependencies/email_verification/emailVerificationController';
import {toast} from '@/core/utils/toast';
import {useStore} from "@tanstack/react-store";

export type EmailVerificationProps = {
    controller: EmailVerificationController;
};

export function EmailVerification({controller}: EmailVerificationProps) {
    const [submitting, setSubmitting] = useState(false);
    const [emailSent, setEmailSent] = useState(false);
    const hasToken = useStore(controller.tokenStore) !== null;

    if (hasToken) {
        // Nothing to render specifically here yet; we auto-trigger the request.
        return (
            <div className="p-2">
                <h3 className="text-lg font-semibold">Validation en cours...</h3>
            </div>
        );
    }

    if (emailSent) {
        return <>L'email contenant le lien de validation a été envoyé.</>;
    }

    const handleInquire = async () => {
        setSubmitting(true);
        try {
            await controller.inquire();
            toast("L'email de validation a été envoyé");
            setEmailSent(true);
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
        } finally {
            setSubmitting(false);
        }
    };

    return (
        <div className="p-4 space-y-4">
            <p>
                Afin de pouvoir apporter votre contribution à Vigie du mensonge, nous devons vérifier votre adresse
                email. Cliquez ci-dessous pour recevoir votre email de validation.
            </p>
            <button
                type="button"
                className="inline-flex items-center justify-center rounded bg-primary px-4 py-2 text-primary-foreground hover:bg-primary/90 disabled:opacity-50"
                onClick={handleInquire}
                disabled={submitting}
            >
                {submitting ? 'Envoi en cours...' : "Recevoir l'email de validation"}
            </button>
        </div>
    );
}
