package http

import (
	applicationauth "institute-platform/backend/internal/application/auth"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewServer(
	host string,
	port string,
	logger *slog.Logger,
	courseHandler *CourseHandler,
	studentHandler *StudentHandler,
	registrationHandler *RegistrationHandler,
	authHandler *AuthHandler,
	authService *applicationauth.Service,
	studentDashboardHandler *StudentDashboardHandler,
) *Server {

	router := chi.NewRouter()

	router.Use(corsMiddleware)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	authMiddleware := NewAuthMiddleware(authService)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/courses", courseHandler.GetCourses)
		r.Get("/courses/{slug}", courseHandler.GetCourseBySlug)

		r.Post("/students", studentHandler.CreateStudent)
		r.Get("/students", studentHandler.GetStudents)
		r.Get("/students/{id}", studentHandler.GetStudent)
		r.Put("/students/{id}", studentHandler.UpdateStudent)

		r.With(authMiddleware.RequireAuth).Post(
			"/registrations",
			registrationHandler.CreateRegistration,
		)

		r.Get(
			"/registrations/{id}",
			registrationHandler.GetRegistration,
		)

		r.Get(
			"/registrations/{id}",
			registrationHandler.GetRegistration,
		)

		r.Get(
			"/students/{id}/registrations",
			registrationHandler.GetStudentRegistrations,
		)

		r.Get(
			"/courses/{id}/registrations",
			registrationHandler.GetCourseRegistrations,
		)

		r.Put(
			"/registrations/{id}/status",
			registrationHandler.UpdateStatus,
		)

		r.Post("/auth/register", authHandler.Register)
		r.Post("/auth/login", authHandler.Login)
		r.Post("/auth/logout", authHandler.Logout)

		r.With(authMiddleware.RequireAuth).Get(
			"/auth/me",
			authHandler.Me,
		)

		r.With(authMiddleware.RequireAuth).Get(
			"/students/me/dashboard",
			studentDashboardHandler.GetDashboard,
		)
	})

	httpServer := &http.Server{
		Addr:              host + ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
