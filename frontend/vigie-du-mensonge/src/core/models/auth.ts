export type AuthJson = {
    accessTokenExpiry: string;
    refreshTokenExpiry: string;
    roles?: string[];
    tag: string;
};

export class Auth {
    public accessTokenExpiry: Date;
    public refreshTokenExpiry: Date;
    public roles: string[];
    public tag: string;

    public static readonly STORAGE_KEY = 'vdm_auth';

    constructor(
        accessTokenExpiry: Date,
        refreshTokenExpiry: Date,
        roles: string[],
        tag: string,
    ) {
        this.accessTokenExpiry = accessTokenExpiry;
        this.refreshTokenExpiry = refreshTokenExpiry;
        this.roles = roles;
        this.tag = tag;
    }

    get isRedactor(): boolean {
        return this.roles.includes('REDACTOR');
    }

    get isModerator(): boolean {
        return this.roles.includes('MODERATOR');
    }

    get accessTokenExpired(): boolean {
        return this.accessTokenExpiry < new Date();
    }

    get refreshTokenExpired(): boolean {
        return this.refreshTokenExpiry < new Date();
    }

    public static fromJson(json: AuthJson): Auth {
        const accessTokenExpiry = new Date(json.accessTokenExpiry);
        accessTokenExpiry.setSeconds(accessTokenExpiry.getSeconds() - 15);

        const refreshTokenExpiry = new Date(json.refreshTokenExpiry);
        refreshTokenExpiry.setSeconds(refreshTokenExpiry.getSeconds() - 15);

        return new Auth(
            accessTokenExpiry,
            refreshTokenExpiry,
            json.roles ?? [],
            json.tag,
        );
    }

    public static fromStorage(): Auth | null {
        try {
            const raw = localStorage.getItem(Auth.STORAGE_KEY);
            if (!raw) return null;
            const parsed = JSON.parse(raw) as Partial<AuthJson>;
            if (!parsed?.accessTokenExpiry || !parsed?.refreshTokenExpiry) return null;
            return Auth.fromJson({
                accessTokenExpiry: parsed.accessTokenExpiry,
                refreshTokenExpiry: parsed.refreshTokenExpiry,
                roles: parsed.roles ?? [],
                tag: parsed.tag ?? '',
            });
        } catch {
            return null;
        }
    }

    public toJson(): AuthJson {
        return {
            accessTokenExpiry: this.accessTokenExpiry.toISOString(),
            refreshTokenExpiry: this.refreshTokenExpiry.toISOString(),
            roles: this.roles,
            tag: this.tag,
        };
    }

    public saveToStorage(): void {
        try {
            localStorage.setItem(Auth.STORAGE_KEY, JSON.stringify(this.toJson()));
        } catch {
            // ignore if storage not available
        }
    }

    public static clearStorage(): void {
        try {
            localStorage.removeItem(Auth.STORAGE_KEY);
        } catch {
            // ignore storage errors
        }
    }
}