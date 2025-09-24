import type {SignUpController} from "@/core/dependencies/auth/signUpController.ts";
import {ProcessSignUp} from "@/core/components/auth/ProcessSignUp.tsx";
import {InquireSignUp} from "@/core/components/auth/InquireSignUp.tsx";

export type SignUpProps = {
    controller: SignUpController;
}

export function SignUp({controller}: SignUpProps) {
    if (controller.token) {
        return <ProcessSignUp
            submit={({password}) => controller.onProcess(password)}/>;
    }

    return <InquireSignUp submit={({email}) => controller.onInquire(email)}/>;
}