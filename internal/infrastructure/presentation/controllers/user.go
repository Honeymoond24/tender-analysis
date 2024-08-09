package controllers

import (
	"fmt"
	"net/http"
)

func RootController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hey Am getting called From Controllers", r.RequestURI)

	_, err := w.Write([]byte(`{"message": "working..."}`))
	if err != nil {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI, "user id - ", r.PathValue("id"))

	_, err := w.Write([]byte(`{"message": "working..."}`))
	if err != nil {
		return
	}
}
