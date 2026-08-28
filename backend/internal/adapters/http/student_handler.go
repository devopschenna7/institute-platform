package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	applicationstudent "institute-platform/backend/internal/application/student"
	domainstudent "institute-platform/backend/internal/domain/student"
)

type StudentHandler struct {
	service *applicationstudent.Service
}

func NewStudentHandler(
	service *applicationstudent.Service,
) *StudentHandler {
	return &StudentHandler{
		service: service,
	}
}

type StudentResponse struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateStudentRequest struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth"`
}

type UpdateStudentRequest struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth"`
}

func toStudentResponse(
	s domainstudent.Student,
) StudentResponse {
	return StudentResponse{
		ID:          s.ID,
		Name:        s.Name,
		Email:       s.Email,
		Phone:       s.Phone,
		DateOfBirth: s.DateOfBirth,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func (h *StudentHandler) CreateStudent(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request CreateStudentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	student := &domainstudent.Student{
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		DateOfBirth: request.DateOfBirth,
	}

	if err := h.service.CreateStudent(
		r.Context(),
		student,
	); err != nil {
		http.Error(
			w,
			"failed to create student",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(
		toStudentResponse(*student),
	)
}

func (h *StudentHandler) GetStudents(
	w http.ResponseWriter,
	r *http.Request,
) {
	students, err := h.service.GetStudents(
		r.Context(),
	)
	if err != nil {
		http.Error(
			w,
			"failed to get students",
			http.StatusInternalServerError,
		)
		return
	}

	response := make(
		[]StudentResponse,
		0,
		len(students),
	)

	for _, student := range students {
		response = append(
			response,
			toStudentResponse(student),
		)
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}

func (h *StudentHandler) GetStudent(
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
			"invalid student id",
			http.StatusBadRequest,
		)
		return
	}

	student, err := h.service.GetStudent(
		r.Context(),
		id,
	)
	if err != nil {
		http.Error(
			w,
			"student not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(
		toStudentResponse(*student),
	)
}

func (h *StudentHandler) UpdateStudent(
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
			"invalid student id",
			http.StatusBadRequest,
		)
		return
	}

	var request UpdateStudentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	student := &domainstudent.Student{
		ID:          id,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		DateOfBirth: request.DateOfBirth,
	}

	if err := h.service.UpdateStudent(
		r.Context(),
		student,
	); err != nil {
		http.Error(
			w,
			"failed to update student",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(
		toStudentResponse(*student),
	)
}
