package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"KratosNew/internal/biz"
	"KratosNew/internal/data"
)

type HTTPServer struct {
	uc *biz.UserUsecase
}

func NewHTTPServer(uc *biz.UserUsecase) *HTTPServer {
	return &HTTPServer{uc: uc}
}

func (s *HTTPServer) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createUser(w, r)

	case http.MethodGet:
		s.getUsers(w, r)

	case http.MethodPatch:
		s.updateUser(w, r)

	case http.MethodDelete:
		s.deleteUser(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *HTTPServer) createUser(w http.ResponseWriter, r *http.Request) {
	var user data.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := s.uc.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *HTTPServer) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.uc.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (s *HTTPServer) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := s.uc.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user.ID = id

	if err := s.uc.Update(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *HTTPServer) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := s.uc.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getID(r *http.Request) (uint64, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) < 2 {
		return 0, strconv.ErrSyntax
	}

	return strconv.ParseUint(parts[len(parts)-1], 10, 64)
}