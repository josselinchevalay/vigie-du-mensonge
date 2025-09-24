import {createFileRoute, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {SignUpController} from "@/core/dependencies/auth/signUpController.ts";
import {SignUp} from "@/core/components/auth/SignUp.tsx";

export const Route = createFileRoute('/sign-up')({
    validateSearch: (search: { token?: string }) => ({token: search.token}),
    beforeLoad: ({search}) => {
        const isAuthenticated = authManager.authStore.state !== null;
        if (isAuthenticated) {
            throw redirect({to: '/', replace: true});
        }

        const controller = new SignUpController(search.token ?? null);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: SignUpController };
    return <SignUp controller={controller}/>;
}