import type {UserRole} from "@/core/models/userRole.ts";

export type UserJson = {
    tag: string;
    roles: UserRole[];
}

export class User {
    public tag: string;
    public roles: UserRole[];

    constructor(tag: string, roles: UserRole[]) {
        this.tag = tag;
        this.roles = roles;
    }

    static fromJson(json: UserJson): User {
        return new User(json.tag, json.roles);
    }
}