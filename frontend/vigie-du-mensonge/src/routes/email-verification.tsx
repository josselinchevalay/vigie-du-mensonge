import {createFileRoute, redirect} from '@tanstack/react-router';
import {EmailVerification} from '@/core/components/email_verification/EmailVerification';
import {EmailVerificationController} from '@/core/dependencies/email_verification/emailVerificationController';
import {authManager} from "@/core/dependencies/auth/authManager.ts";

export const Route = createFileRoute('/email-verification')({
    validateSearch: (search: { token?: string }) => ({token: search.token}),
    beforeLoad: ({search}) => {
        const auth = authManager.authStore.state;

        if (!auth || auth.emailVerified) {
            throw redirect({to: '/', replace: true});
        }

        const token = search.token ?? null;
        const controller = new EmailVerificationController(token);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: EmailVerificationController };
    return <EmailVerification controller={controller}/>;
}