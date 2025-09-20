import {createFileRoute} from '@tanstack/react-router';

export const Route = createFileRoute('/email-verification')({
    component: EmailVerification,
});

function EmailVerification() {
    return (
        <div className="p-2">
            <h3>Email verification</h3>
        </div>
    );
}