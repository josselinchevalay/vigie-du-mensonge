import { api } from "@/core/dependencies/api";

export class EmailVerificationClient {
    /**
     * Request a verification email to be sent to the currently authenticated user.
     * OpenAPI: POST /email-verification/inquire -> 204
     */
    async inquire(): Promise<void> {
        await api.post("email-verification/inquire");
    }

    /**
     * Validate the email address using the token received by email.
     * OpenAPI: POST /email-verification/process (body: { token: string }) -> 204
     */
    async process(token: string): Promise<void> {
        await api.post("email-verification/process", {
            json: { token },
        });
    }
}