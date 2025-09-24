import {createFileRoute, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {SignInController} from "@/core/dependencies/auth/signInController.ts";
import {SignIn} from "@/core/components/auth/SignIn.tsx";
import {authClient} from "@/core/dependencies/auth/authClient.ts";

export const Route = createFileRoute('/sign-in')({
    beforeLoad: () => {
        const isAuthenticated = authManager.authStore.state !== null;
        if (isAuthenticated) {
            throw redirect({to: '/', replace: true});
        }

        const controller = new SignInController(authClient);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: SignInController };
    return <SignIn controller={controller}/>;
}