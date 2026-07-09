package usercontroller

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
}

type User struct {
	Name     string `json:"name"`
	Password int    `json:"password"`
}

func NewUserController() *UserController {
	return &UserController{}
}

// write index handler
func (user *UserController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(User{Name: "Sam", Password: 123})
}

type Controller interface {
	Index(w http.ResponseWriter, r *http.Request)
}
