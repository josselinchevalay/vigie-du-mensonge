import type {PasswordUpdateController} from "@/core/dependencies/password_update/passwordUpdateController.ts";
import {useStore} from "@tanstack/react-store";
import {InquirePasswordUpdate} from "@/core/components/password_update/InquirePasswordUpdate.tsx";
import {ProcessPasswordUpdate} from "@/core/components/password_update/ProcessPasswordUpdate.tsx";

export type PasswordUpdateProps = {
    controller: PasswordUpdateController;
};

export function PasswordUpdate({controller}: PasswordUpdateProps) {
    const hasToken = useStore(controller.tokenStore) !== null;

    if (!hasToken) {
        return <InquirePasswordUpdate submitForm={({email}) => controller.onInquire(email)}/>;
    }

    return (
        <ProcessPasswordUpdate submitForm={({password}) => controller.onProcess(password)}/>
    );
}