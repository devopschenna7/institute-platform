package registration

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusCancelled Status = "cancelled"
	StatusCompleted Status = "completed"
)

type Registration struct {
	ID        int64
	StudentID int64
	CourseID  int64
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}
