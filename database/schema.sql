CREATE
EXTENSION IF NOT EXISTS pgcrypto;


CREATE TABLE politicians
(
    id         UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_politicians PRIMARY KEY (id),

    last_name  TEXT        NOT NULL,
    first_name TEXT        NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);


CREATE TABLE governments
(
    id                UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_governments PRIMARY KEY (id),

    prime_minister_id UUID        NOT NULL,
    CONSTRAINT fk_governments_prime_minister FOREIGN KEY (prime_minister_id) REFERENCES politicians (id),

    reference_id      SMALLINT    NOT NULL,
    CONSTRAINT uq_governments_reference UNIQUE (reference_id),

    start_date        TIMESTAMPTZ NOT NULL,
    end_date          TIMESTAMPTZ,

    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMPTZ
);


CREATE TABLE occupations
(
    id                     UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_occupations PRIMARY KEY (id),

    politician_id          UUID        NOT NULL,
    CONSTRAINT fk_occupations_politician FOREIGN KEY (politician_id) REFERENCES politicians (id),

    government_id          UUID,
    CONSTRAINT fk_occupations_government FOREIGN KEY (government_id) REFERENCES governments (id),

    presidential_reference SMALLINT,
    CONSTRAINT uq_occupations_presidential UNIQUE (presidential_reference),

    CONSTRAINT ck_occupations CHECK ((presidential_reference IS NOT NULL AND government_id IS NULL)
        OR (presidential_reference IS NULL AND government_id IS NOT NULL)),

    code                   TEXT        NOT NULL,
    title                  TEXT        NOT NULL,
    start_date             TIMESTAMPTZ NOT NULL,
    end_date               TIMESTAMPTZ,

    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at             TIMESTAMPTZ
);


CREATE TABLE roles
(
    id         UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_roles PRIMARY KEY (id),

    name       TEXT        NOT NULL,
    CONSTRAINT uq_roles_name UNIQUE (name),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);


CREATE TABLE users
(
    id             UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_users PRIMARY KEY (id),

    email          TEXT        NOT NULL,
    CONSTRAINT uq_users_email UNIQUE (email),

    email_verified BOOLEAN     NOT NULL DEFAULT FALSE,
    password       TEXT        NOT NULL,

    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMPTZ
);


CREATE TABLE user_roles
(
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles (id),
    CONSTRAINT pk_user_roles PRIMARY KEY (user_id, role_id)
);


CREATE TABLE user_tokens
(
    id       UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_user_tokens PRIMARY KEY (id),

    user_id  UUID        NOT NULL,
    CONSTRAINT fk_user_tokens_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,

    category TEXT        NOT NULL,
    CONSTRAINT ck_user_tokens_category CHECK (category IN ('REFRESH', 'PASSWORD', 'EMAIL')),

    hash     TEXT        NOT NULL,
    expiry   TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_tokens_category_hash ON user_tokens (category, hash);
CREATE INDEX IF NOT EXISTS idx_user_tokens_expiry ON user_tokens (expiry);


CREATE TABLE articles
(
    id           UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_articles PRIMARY KEY (id),

    author_id    UUID        NOT NULL,
    CONSTRAINT fk_articles_author FOREIGN KEY (author_id) REFERENCES users (id),

    moderator_id UUID,
    CONSTRAINT fk_articles_moderator FOREIGN KEY (moderator_id) REFERENCES users (id),

    status       TEXT        NOT NULL,
    CONSTRAINT ck_articles_status CHECK (status IN
                                         ('DRAFT', 'PUBLISHED', 'ARCHIVED', 'UNDER_REVIEW', 'CHANGE_REQUESTED')),

    category     TEXT        NOT NULL,
    CONSTRAINT ck_articles_category CHECK (category IN ('LIE', 'FALSEHOOD')),

    title        TEXT        NOT NULL,
    body         TEXT        NOT NULL,
    event_date   TIMESTAMPTZ NOT NULL,

    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);


CREATE TABLE article_sources
(
    article_id UUID NOT NULL,
    CONSTRAINT fk_article_sources_article FOREIGN KEY (article_id) REFERENCES articles (id),

    url        TEXT NOT NULL,

    CONSTRAINT pk_article_sources PRIMARY KEY (article_id, url)
);


CREATE TABLE article_tags
(
    article_id UUID NOT NULL,
    CONSTRAINT fk_article_tags_article FOREIGN KEY (article_id) REFERENCES articles (id),

    tag        TEXT NOT NULL,
    CONSTRAINT pk_article_tags PRIMARY KEY (article_id, tag)
);


CREATE TABLE article_reviews
(
    id           UUID        NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT pk_article_reviews PRIMARY KEY (id),

    article_id   UUID        NOT NULL,
    CONSTRAINT fk_article_reviews_article FOREIGN KEY (article_id) REFERENCES articles (id),

    moderator_id UUID        NOT NULL,
    CONSTRAINT fk_article_reviews_moderator FOREIGN KEY (moderator_id) REFERENCES users (id),

    notes        TEXT        NOT NULL,
    seen         BOOLEAN     NOT NULL DEFAULT FALSE,

    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);


CREATE TABLE article_politicians
(
    article_id    UUID NOT NULL,
    CONSTRAINT fk_article_politicians_article FOREIGN KEY (article_id) REFERENCES articles (id),

    politician_id UUID NOT NULL,
    CONSTRAINT fk_article_politicians_politician FOREIGN KEY (politician_id) REFERENCES politicians (id),

    CONSTRAINT pk_article_politicians PRIMARY KEY (article_id, politician_id)
);