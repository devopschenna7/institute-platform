package http

import (
	"encoding/json"
	"errors"
	"net/http"

	applicationauth "institute-platform/backend/internal/application/auth"
)

const sessionCookieName = "session_id"

type AuthHandler struct {
	service *applicationauth.Service
}

func NewAuthHandler(
	service *applicationauth.Service,
) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

type RegisterRequest struct {
	StudentCode string `json:"student_code"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if request.StudentCode == "" ||
		request.Email == "" ||
		request.Password == "" {

		http.Error(
			w,
			"student_code, email and password are required",
			http.StatusBadRequest,
		)
		return
	}

	err := h.service.Register(
		r.Context(),
		request.StudentCode,
		request.Email,
		request.Password,
	)

	if err != nil {
		if errors.Is(
			err,
			applicationauth.ErrStudentNotFound,
		) {
			http.Error(
				w,
				"account already exists",
				http.StatusConflict,
			)
			return
		}

		if errors.Is(
			err,
			applicationauth.ErrAccountExists,
		) {
			http.Error(
				w,
				"student ID not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(
		map[string]string{
			"message": "account created successfully",
		},
	)
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if request.Email == "" ||
		request.Password == "" {
		http.Error(
			w,
			"email and password are required",
			http.StatusBadRequest,
		)
		return
	}

	sessionID, err := h.service.Login(
		r.Context(),
		request.Email,
		request.Password,
	)

	if err != nil {
		if errors.Is(
			err,
			applicationauth.ErrInvalidCredentials,
		) {
			http.Error(
				w,
				"invalid email or password",
				http.StatusUnauthorized,
			)
			return
		}

		http.Error(
			w,
			"failed to login",
			http.StatusInternalServerError,
		)
		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     sessionCookieName,
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   24 * 60 * 60,
		},
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(
		map[string]string{
			"message": "login successful",
		},
	)
}

func (h *AuthHandler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, err := r.Cookie(sessionCookieName)

	if err == nil && cookie.Value != "" {
		_ = h.service.Logout(
			r.Context(),
			cookie.Value,
		)
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     sessionCookieName,
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
			SameSite: http.SameSiteLaxMode,
		},
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(
		map[string]string{
			"message": "logout successful",
		},
	)
}

func (h *AuthHandler) Me(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(
		map[string]int64{
			"student_id": studentID,
		},
	)
}
