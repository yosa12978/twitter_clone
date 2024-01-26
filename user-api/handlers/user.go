package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosa12978/twitter/user-api/logging"
	"github.com/yosa12978/twitter/user-api/repos"
	"github.com/yosa12978/twitter/user-api/types"
)

type User interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	repo   repos.User
	logger logging.Logger
}

func NewUser(ctx context.Context) User {
	return &userHandler{
		repo:   repos.NewUser(ctx),
		logger: logging.New("user handler"),
	}
}

func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.repo.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := writeJSON(w, 200, types.NewArrRes(users)); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (u *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := u.repo.FindById(r.Context(), vars["id"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := writeJSON(w, 200, user); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var dto types.LoginDto
	if err := readJSON(r, &dto); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user, err := u.repo.FindByCredentialsEmail(r.Context(), dto.Email, dto.Password)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	// do jwt generation here
	if err := writeJSON(w, 200, user.ToDto()); err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func (u *userHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signupDto types.SignupDto
	if err := readJSON(r, &signupDto); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user := types.FromDto(signupDto)
	id, err := u.repo.Create(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := writeJSON(w, 201, map[string]interface{}{"id": id}); err != nil {
		http.Error(w, err.Error(), 400)
	}
}
