import {createFileRoute, redirect} from "@tanstack/react-router";
import {RedactorIndex} from "@/core/components/redactor/RedactorIndex.tsx";
import {authManager} from "@/core/dependencies/auth/authManager.ts";

export const Route = createFileRoute('/redactor/articles')({
    beforeLoad: () => {
        const auth = authManager.authStore.state;
        if (!auth || !auth.isRedactor) {
            throw redirect({to: '/', replace: true});
        }
    },
    component: RedactorIndex,
});
