import type {PasswordUpdateController} from "@/core/dependencies/password_update/passwordUpdateController.ts";
import {useStore} from "@tanstack/react-store";
import {InquirePasswordUpdateForm} from "@/core/components/password_update/InquirePasswordUpdateForm.tsx";
import {ProcessPasswordUpdateForm} from "@/core/components/password_update/ProcessPasswordUpdateForm.tsx";

export type PasswordUpdateProps = {
    controller: PasswordUpdateController;
};

export function PasswordUpdate({controller}: PasswordUpdateProps) {
    const hasToken = useStore(controller.tokenStore) !== null;

    if (!hasToken) {
        return <InquirePasswordUpdateForm submitForm={({email}) => controller.onInquire(email)}/>;
    }

    return (
        <ProcessPasswordUpdateForm submitForm={({password}) => controller.onProcess(password)}/>
    );
}