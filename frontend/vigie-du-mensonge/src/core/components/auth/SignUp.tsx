import type {SignUpController} from "@/core/dependencies/auth/signUpController.ts";
import {ProcessSignUp} from "@/core/components/auth/ProcessSignUp.tsx";
import {InquireSignUp} from "@/core/components/auth/InquireSignUp.tsx";

export type SignUpProps = {
    controller: SignUpController;
}

export function SignUp({controller}: SignUpProps) {
    if (controller.token) {
        return <ProcessSignUp
            submit={({username, password}) => controller.onProcess(username, password)}/>;
    }

    return <InquireSignUp submitForm={({email}) => controller.onInquire(email)}/>;
}