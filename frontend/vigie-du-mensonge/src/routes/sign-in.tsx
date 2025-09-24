import {createFileRoute, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {SignInController} from "@/core/dependencies/auth/signInController.ts";
import {SignIn} from "@/core/components/auth/SignIn.tsx";

export const Route = createFileRoute('/sign-in')({
    beforeLoad: () => {
        const isAuthenticated = authManager.authStore.state !== null;
        if (isAuthenticated) {
            throw redirect({to: '/', replace: true});
        }

        const controller = new SignInController();
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: SignInController };
    return <SignIn controller={controller}/>;
}