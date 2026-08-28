package http

import (
	"encoding/json"
	"net/http"

	applicationstudent "institute-platform/backend/internal/application/student"
)

type StudentDashboardHandler struct {
	service *applicationstudent.Service
}

func NewStudentDashboardHandler(
	service *applicationstudent.Service,
) *StudentDashboardHandler {
	return &StudentDashboardHandler{
		service: service,
	}
}

func (h *StudentDashboardHandler) GetDashboard(
	w http.ResponseWriter,
	r *http.Request,
) {
	studentID, ok := GetStudentIDFromContext(
		r.Context(),
	)

	if !ok {
		http.Error(
			w,
			"authentication required",
			http.StatusUnauthorized,
		)
		return
	}

	dashboard, err := h.service.GetDashboard(
		r.Context(),
		studentID,
	)
	if err != nil {
		http.Error(
			w,
			"failed to get dashboard",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(dashboard)
}
