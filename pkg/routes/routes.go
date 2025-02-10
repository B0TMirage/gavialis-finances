package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	handlers "github.com/B0TMirage/gavialis-finances/pkg/handlers/login"
	models "github.com/B0TMirage/gavialis-finances/pkg/models/users"
)

func Regrouter() *http.ServeMux {
	mux := &http.ServeMux{}

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Канадская береговая")
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		var user *models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		loggeduser, err := handlers.Login(user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authCookie := &http.Cookie{
			Name:  "id",
			Value: strconv.Itoa(loggeduser.ID),
		}

		http.SetCookie(w, authCookie)
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /logout", func(w http.ResponseWriter, r *http.Request) {
		logoutCookie := &http.Cookie{
			Name:   "id",
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(w, logoutCookie)
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = handlers.Register(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	return mux
}
