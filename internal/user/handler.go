package user

import (
	"go-rest/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET("/api/users", h.GetUsers)
	router.GET("/api/users/:userId", h.GetUserById)
	router.POST("/api/users", h.CreateUser)
	router.PUT("/api/users/:userId", h.UpdateUser)
	router.PATCH("/api/users/:userId", h.PartUpdateUser)
	router.DELETE("/api/users/:userId", h.DeleteUser)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("all users"))
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("get one user"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("create user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("update user"))
}

func (h *handler) PartUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("update user"))
}
