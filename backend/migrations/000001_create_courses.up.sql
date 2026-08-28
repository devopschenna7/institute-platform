CREATE TABLE courses (
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(255) NOT NULL,

    slug VARCHAR(255) NOT NULL UNIQUE,

    description TEXT NOT NULL,

    duration VARCHAR(100),

    level VARCHAR(100),

    status VARCHAR(50) NOT NULL DEFAULT 'active',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_courses_status
    ON courses(status);