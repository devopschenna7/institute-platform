package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	applicationcourse "institute-platform/backend/internal/application/course"
	domaincourse "institute-platform/backend/internal/domain/course"
)

type CourseHandler struct {
	service *applicationcourse.Service
}

func NewCourseHandler(
	service *applicationcourse.Service,
) *CourseHandler {
	return &CourseHandler{
		service: service,
	}
}

type CourseResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Level       string `json:"level"`
}

func toCourseResponse(c domaincourse.Course) CourseResponse {
	return CourseResponse{
		ID:          c.ID,
		Name:        c.Name,
		Slug:        c.Slug,
		Description: c.Description,
		Duration:    c.Duration,
		Level:       c.Level,
	}
}

func (h *CourseHandler) GetCourses(
	w http.ResponseWriter,
	r *http.Request,
) {
	courses, err := h.service.GetCourses(r.Context())
	if err != nil {
		http.Error(
			w,
			"failed to get courses",
			http.StatusInternalServerError,
		)
		return
	}

	response := make(
		[]CourseResponse,
		0,
		len(courses),
	)

	for _, course := range courses {
		response = append(
			response,
			toCourseResponse(course),
		)
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}

func (h *CourseHandler) GetCourseBySlug(
	w http.ResponseWriter,
	r *http.Request,
) {
	slug := chi.URLParam(r, "slug")

	course, err := h.service.GetCourseBySlug(
		r.Context(),
		slug,
	)

	if err != nil {
		if errors.Is(err, domaincourse.ErrNotFound) {
			http.Error(
				w,
				"course not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"failed to get course",
			http.StatusInternalServerError,
		)
		return
	}

	response := toCourseResponse(*course)

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}