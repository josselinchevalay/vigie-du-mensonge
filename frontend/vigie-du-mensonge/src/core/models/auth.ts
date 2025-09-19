export type AuthJson = {
    accessTokenExpiry: string;
    refreshTokenExpiry: string;
    emailVerified?: boolean;
    roles?: string[];
};

export class Auth {
    public accessTokenExpiry: Date;
    public refreshTokenExpiry: Date;
    public emailVerified: boolean;
    public roles: string[];

    private static readonly STORAGE_KEY = 'vdm_auth';

    constructor(
        accessTokenExpiry: Date,
        refreshTokenExpiry: Date,
        emailVerified: boolean,
        roles: string[],
    ) {
        this.accessTokenExpiry = accessTokenExpiry;
        this.refreshTokenExpiry = refreshTokenExpiry;
        this.emailVerified = emailVerified;
        this.roles = roles;
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
            json.emailVerified ?? false,
            json.roles ?? [],
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
                emailVerified: parsed.emailVerified ?? false,
                roles: parsed.roles ?? [],
            });
        } catch {
            return null;
        }
    }

    public toJson(): AuthJson {
        return {
            accessTokenExpiry: this.accessTokenExpiry.toISOString(),
            refreshTokenExpiry: this.refreshTokenExpiry.toISOString(),
            emailVerified: this.emailVerified,
            roles: this.roles,
        };
    }

    public saveToStorage(): void {
        try {
            localStorage.setItem(Auth.STORAGE_KEY, JSON.stringify(this.toJson()));
        } catch {
            // ignore if storage not available
        }
    }
}