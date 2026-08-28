CREATE TABLE sessions (
    id UUID PRIMARY KEY,

    student_id BIGINT NOT NULL,

    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_sessions_student
        FOREIGN KEY (student_id)
        REFERENCES students(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_sessions_student_id
ON sessions(student_id);

CREATE INDEX idx_sessions_expires_at
ON sessions(expires_at);