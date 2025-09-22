import {useStore} from "@tanstack/react-store";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {Link, useLocation} from "@/core/utils/router";

export default function VerifyMyEmailLink() {
    const auth = useStore(authManager.authStore);
    const location = useLocation();

    if (!auth) {
        return <></>;
    }

    if (auth.emailVerified || location.pathname === '/email-verification') {
        return <></>;
    }

    return <>
        <Link
            to="/email-verification"
            className="items-center rounded-md border px-3 py-2 font-medium hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background"
            search={{token: undefined}}>
            Vérifier mon email
        </Link>
    </>;
}