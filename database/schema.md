# Database Schema (MPD)

The following Mermaid ER diagram is generated from `db/schema.sql` and is compatible with GitHub rendering.

```mermaid
erDiagram
    POLITICIANS {
        UUID id PK
        TEXT last_name
        TEXT first_name
        TEXT image_url
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    GOVERNMENTS {
        UUID id PK
        UUID prime_minister_id FK
        SMALLINT reference_id UK
        TIMESTAMPTZ start_date
        TIMESTAMPTZ end_date
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    OCCUPATIONS {
        UUID id PK
        UUID politician_id FK
        UUID government_id FK
        SMALLINT presidential_reference UK
        TEXT code
        TEXT title
        TIMESTAMPTZ start_date
        TIMESTAMPTZ end_date
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    ROLES {
        UUID id PK
        TEXT name UK
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    USERS {
        UUID id PK
        TEXT email UK
        BOOLEAN email_verified
        TEXT password
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    USER_ROLES {
        UUID user_id PK
        UUID role_id PK
    }
    
    REFRESH_TOKENS {
        UUID id PK
        UUID user_id FK
        TIMESTAMPTZ expiry
    }

    ARTICLES {
        UUID id PK
        UUID author_id FK
        UUID moderator_id FK
        TEXT status
        TEXT category
        TEXT title
        TEXT body
        TIMESTAMPTZ event_date
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    ARTICLE_SOURCES {
        UUID article_id PK
        TEXT url PK
    }

    ARTICLE_TAGS {
        UUID article_id PK
        TEXT tag PK
    }

    ARTICLE_REVIEWS {
        UUID id PK
        UUID article_id FK
        UUID moderator_id FK
        TEXT notes
        BOOLEAN seen
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
        TIMESTAMPTZ deleted_at
    }

    ARTICLE_POLITICIANS {
        UUID article_id PK
        UUID politician_id PK
    }

    %% Relationships
    POLITICIANS ||--o{ GOVERNMENTS : "prime_minister"
    POLITICIANS ||--o{ OCCUPATIONS : "holds"
    GOVERNMENTS ||--o{ OCCUPATIONS : "contains"

    USERS ||--o{ ARTICLES : "authors"
    USERS ||--o{ ARTICLES : "moderates"

    ARTICLES ||--o{ ARTICLE_SOURCES : "has"
    ARTICLES ||--o{ ARTICLE_TAGS : "tagged"
    ARTICLES ||--o{ ARTICLE_REVIEWS : "reviewed by"

    USERS ||--o{ ARTICLE_REVIEWS : "moderates"

    ARTICLES ||--o{ ARTICLE_POLITICIANS : "mentions"
    POLITICIANS ||--o{ ARTICLE_POLITICIANS : "mentioned in"

    USERS ||--o{ USER_ROLES : "has"
    ROLES ||--o{ USER_ROLES : "assigned to"
```

Notes:
- Check constraints (like status/category enums and exclusivity on OCCUPATIONS) are documented in SQL but not visualized in Mermaid.
- Nullable foreign keys (e.g., moderator_id, government_id) are shown as FK attributes without cardinality change; semantics are documented in SQL.
