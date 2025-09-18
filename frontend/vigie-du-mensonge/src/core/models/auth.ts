export type AuthJson = {
    accessTokenExpiry: string;
    refreshTokenExpiry: string;
    emailVerified: boolean;
};

export class Auth {
    public accessTokenExpiry: Date;
    public refreshTokenExpiry: Date;
    public emailVerified: boolean;

    constructor(
        accessTokenExpiry: Date,
        refreshTokenExpiry: Date,
        emailVerified: boolean,
    ) {
        this.accessTokenExpiry = accessTokenExpiry;
        this.refreshTokenExpiry = refreshTokenExpiry;
        this.emailVerified = emailVerified;
    }

    public static fromJson(json: AuthJson): Auth {
        const accessTokenExpiry = new Date(json.accessTokenExpiry);
        accessTokenExpiry.setSeconds(accessTokenExpiry.getSeconds() - 15);

        const refreshTokenExpiry = new Date(json.refreshTokenExpiry);
        refreshTokenExpiry.setSeconds(refreshTokenExpiry.getSeconds() - 15);

        return new Auth(
            accessTokenExpiry,
            refreshTokenExpiry,
            json.emailVerified,
        );
    }

    get accessTokenExpired(): boolean {
        return this.accessTokenExpiry < new Date();
    }

    get refreshTokenExpired(): boolean {
        return this.refreshTokenExpiry < new Date();
    }
}