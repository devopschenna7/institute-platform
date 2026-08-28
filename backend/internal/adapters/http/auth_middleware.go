package http

import (
	"context"
	"net/http"

	applicationauth "institute-platform/backend/internal/application/auth"
)

type contextKey string

const studentIDContextKey contextKey = "student_id"

type AuthMiddleware struct {
	service *applicationauth.Service
}

func NewAuthMiddleware(
	service *applicationauth.Service,
) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}

func (m *AuthMiddleware) RequireAuth(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(sessionCookieName)
		if err != nil || cookie.Value == "" {
			http.Error(
				w,
				"authentication required",
				http.StatusUnauthorized,
			)
			return
		}

		studentID, err := m.service.GetSessionStudentID(
			r.Context(),
			cookie.Value,
		)
		if err != nil {
			http.Error(
				w,
				"invalid or expired session",
				http.StatusUnauthorized,
			)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			studentIDContextKey,
			studentID,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}

func GetStudentIDFromContext(
	ctx context.Context,
) (int64, bool) {
	studentID, ok := ctx.Value(
		studentIDContextKey,
	).(int64)

	return studentID, ok
}
