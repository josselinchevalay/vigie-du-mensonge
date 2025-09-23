import type {SignUpController} from "@/core/dependencies/sign_up/signUpController.ts";
import {ProcessSignUp} from "@/core/components/sign_up/ProcessSignUp.tsx";
import {InquireSignUp} from "@/core/components/sign_up/InquireSignUp.tsx";

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