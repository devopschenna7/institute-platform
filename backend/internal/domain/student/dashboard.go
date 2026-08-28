package student

import "time"

type Dashboard struct {
	Student       Student
	Registrations []DashboardRegistration
}

type DashboardRegistration struct {
	RegistrationID int64
	CourseID       int64
	CourseName     string
	CourseSlug     string
	CourseDuration string
	CourseLevel    string
	Status         string
	RegisteredAt   time.Time
}