export const UserRoles = {
    ADMIN: "ADMIN",
    MODERATOR: "MODERATOR",
    REDACTOR: "REDACTOR",
} as const;

export type UserRole = keyof typeof UserRoles;

export const UserRoleLabels: Record<UserRole, string> = {
    ADMIN: 'Administrateur',
    MODERATOR: 'Modérateur',
    REDACTOR: 'Redacteur',
};