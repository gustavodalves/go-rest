package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gustavodalves/go-api/internal/application"
)

type UserHandler struct {
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Post(w http.ResponseWriter, r *http.Request) {
	postRequest := &application.RegisterNewUserDTO{}
	err := json.NewDecoder(r.Body).Decode(postRequest)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	h.service.Register(*postRequest)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	users, _ := h.service.GetAll()

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUnique(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	idInt, _ := strconv.ParseInt(idString, 10, 64)
	user, _ := h.service.GetById(uint64(idInt))

	json.NewEncoder(w).Encode(user)
}
