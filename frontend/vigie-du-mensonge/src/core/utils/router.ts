import type {NavigateOptions} from "@tanstack/react-router";
import {router} from '@/main';

export {
    Link,
    useNavigate,
    useLocation,
    redirect,
    type NavigateOptions,
} from '@tanstack/react-router';

export const navigate = (opts: NavigateOptions) => router.navigate(opts);