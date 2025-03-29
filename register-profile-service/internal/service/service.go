package service

import "net/http"

type User interface {
	SingUp(w http.ResponseWriter, r *http.Request)
	SingIn(w http.ResponseWriter, r *http.Request)
}

type Service struct {
	User
}
