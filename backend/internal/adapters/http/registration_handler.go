package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	applicationregistration "institute-platform/backend/internal/application/registration"
	domainregistration "institute-platform/backend/internal/domain/registration"
)

type RegistrationHandler struct {
	service *applicationregistration.Service
}

func NewRegistrationHandler(
	service *applicationregistration.Service,
) *RegistrationHandler {
	return &RegistrationHandler{
		service: service,
	}
}

type CreateRegistrationRequest struct {
	StudentID int64 `json:"student_id"`
	CourseID  int64 `json:"course_id"`
}

type UpdateRegistrationStatusRequest struct {
	Status domainregistration.Status `json:"status"`
}

type RegistrationResponse struct {
	ID        int64  `json:"id"`
	StudentID int64  `json:"student_id"`
	CourseID  int64  `json:"course_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toRegistrationResponse(
	registration domainregistration.Registration,
) RegistrationResponse {
	return RegistrationResponse{
		ID:        registration.ID,
		StudentID: registration.StudentID,
		CourseID:  registration.CourseID,
		Status:    string(registration.Status),
		CreatedAt: registration.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: registration.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (h *RegistrationHandler) CreateRegistration(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request CreateRegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if request.StudentID <= 0 || request.CourseID <= 0 {
		http.Error(
			w,
			"student_id and course_id are required",
			http.StatusBadRequest,
		)
		return
	}

	registration, err := h.service.CreateRegistration(
		r.Context(),
		request.StudentID,
		request.CourseID,
	)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusConflict,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(
		toRegistrationResponse(*registration),
	)
}

func (h *RegistrationHandler) GetRegistration(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.ParseInt(
		chi.URLParam(r, "id"),
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"invalid registration id",
			http.StatusBadRequest,
		)
		return
	}

	registration, err := h.service.GetRegistration(
		r.Context(),
		id,
	)
	if err != nil {
		http.Error(
			w,
			"registration not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(
		toRegistrationResponse(*registration),
	)
}

func (h *RegistrationHandler) GetStudentRegistrations(
	w http.ResponseWriter,
	r *http.Request,
) {
	studentID, err := strconv.ParseInt(
		chi.URLParam(r, "id"),
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"invalid student id",
			http.StatusBadRequest,
		)
		return
	}

	registrations, err := h.service.GetStudentRegistrations(
		r.Context(),
		studentID,
	)
	if err != nil {
		http.Error(
			w,
			"failed to get registrations",
			http.StatusInternalServerError,
		)
		return
	}

	response := make(
		[]RegistrationResponse,
		0,
		len(registrations),
	)

	for _, registration := range registrations {
		response = append(
			response,
			toRegistrationResponse(registration),
		)
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}

func (h *RegistrationHandler) GetCourseRegistrations(
	w http.ResponseWriter,
	r *http.Request,
) {
	courseID, err := strconv.ParseInt(
		chi.URLParam(r, "id"),
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"invalid course id",
			http.StatusBadRequest,
		)
		return
	}

	registrations, err := h.service.GetCourseRegistrations(
		r.Context(),
		courseID,
	)
	if err != nil {
		http.Error(
			w,
			"failed to get registrations",
			http.StatusInternalServerError,
		)
		return
	}

	response := make(
		[]RegistrationResponse,
		0,
		len(registrations),
	)

	for _, registration := range registrations {
		response = append(
			response,
			toRegistrationResponse(registration),
		)
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}

func (h *RegistrationHandler) UpdateStatus(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := strconv.ParseInt(
		chi.URLParam(r, "id"),
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"invalid registration id",
			http.StatusBadRequest,
		)
		return
	}

	var request UpdateRegistrationStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	switch request.Status {
	case domainregistration.StatusPending,
		domainregistration.StatusConfirmed,
		domainregistration.StatusCancelled,
		domainregistration.StatusCompleted:
	default:
		http.Error(
			w,
			"invalid registration status",
			http.StatusBadRequest,
		)
		return
	}

	if err := h.service.UpdateStatus(
		r.Context(),
		id,
		request.Status,
	); err != nil {
		http.Error(
			w,
			"failed to update registration status",
			http.StatusInternalServerError,
		)
		return
	}

	registration, err := h.service.GetRegistration(
		r.Context(),
		id,
	)
	if err != nil {
		http.Error(
			w,
			"registration not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(
		toRegistrationResponse(*registration),
	)
}