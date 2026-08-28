CREATE TABLE registrations (
    id BIGSERIAL PRIMARY KEY,

    student_id BIGINT NOT NULL,
    course_id BIGINT NOT NULL,

    status VARCHAR(20) NOT NULL DEFAULT 'pending',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_registrations_student
        FOREIGN KEY (student_id)
        REFERENCES students(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_registrations_course
        FOREIGN KEY (course_id)
        REFERENCES courses(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_registrations_student_course
        UNIQUE (student_id, course_id),

    CONSTRAINT chk_registrations_status
        CHECK (status IN ('pending', 'confirmed', 'cancelled', 'completed'))
);