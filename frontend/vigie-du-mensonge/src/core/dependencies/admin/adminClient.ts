import type {KyInstance} from "ky";
import {User, type UserJson} from "@/core/models/user.ts";

export class AdminClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async findUserByTag(tag: string): Promise<User> {
        const res = await this.api
            .get(`admin/users/${tag}`)
            .json<UserJson>();

        return User.fromJson(res);
    }
}