CREATE TABLE student_accounts (
    id BIGSERIAL PRIMARY KEY,

    student_id BIGINT NOT NULL,

    email VARCHAR(150) NOT NULL,
    password_hash TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_student_accounts_student
        FOREIGN KEY (student_id)
        REFERENCES students(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_student_accounts_student
        UNIQUE (student_id),

    CONSTRAINT uq_student_accounts_email
        UNIQUE (email)
);