package bootstrap

import (
	"context"
	"log/slog"

	httpadapter "institute-platform/backend/internal/adapters/http"
	postgresadapter "institute-platform/backend/internal/adapters/postgres"

	applicationauth "institute-platform/backend/internal/application/auth"
	applicationcourse "institute-platform/backend/internal/application/course"
	applicationregistration "institute-platform/backend/internal/application/registration"
	applicationstudent "institute-platform/backend/internal/application/student"
	"institute-platform/backend/internal/config"
)

type Application struct {
	Server *httpadapter.Server
}

func New(
	ctx context.Context,
	cfg config.Config,
	logger *slog.Logger,
) (*Application, error) {

	db, err := postgresadapter.NewConnectionPool(
		ctx,
		cfg.Database,
	)
	if err != nil {
		return nil, err
	}

	// Infrastructure
	courseRepository := postgresadapter.NewCourseRepository(db)
	studentRepository := postgresadapter.NewStudentRepository(db)
	registrationRepository := postgresadapter.NewRegistrationRepository(db)
	authRepository := postgresadapter.NewAuthRepository(db)

	// Application
	courseService := applicationcourse.NewService(
		courseRepository,
	)

	studentService := applicationstudent.NewService(
		studentRepository,
	)
	authService := applicationauth.NewService(
		authRepository,
		studentRepository,
	)

	authHandler := httpadapter.NewAuthHandler(
		authService,
	)

	registrationService := applicationregistration.NewService(
		registrationRepository,
	)
	registrationHandler := httpadapter.NewRegistrationHandler(
		registrationService,
	)

	// HTTP adapter
	courseHandler := httpadapter.NewCourseHandler(
		courseService,
	)
	studentHandler := httpadapter.NewStudentHandler(
		studentService,
	)

	studentDashboardHandler := httpadapter.NewStudentDashboardHandler(
		studentService,
	)

	server := httpadapter.NewServer(
		cfg.Server.Host,
		cfg.Server.Port,
		logger,
		courseHandler,
		studentHandler,
		registrationHandler,
		authHandler,
		authService,
		studentDashboardHandler,
	)

	return &Application{
		Server: server,
	}, nil
}
