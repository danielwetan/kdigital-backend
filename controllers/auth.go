package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/danielwetan/kdigital-backend/models"

	"github.com/danielwetan/kdigital-backend/helpers"
)

func Register(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		username, password := r.FormValue("username"), r.FormValue("password")
		hashedPassword, _ := helpers.HashPassword(password)
		register := &models.Register{
			Username: username,
			Password: hashedPassword,
		}

		_, err = db.Exec(helpers.Query["register"], register.Username, register.Password)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body := "Register success"
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		db, err := helpers.Connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		login := models.Login{}
		username, password := r.FormValue("username"), r.FormValue("password")
		err = db.
			QueryRow(helpers.Query["login"], username).
			Scan(&login.Username, &login.Password)
		if err != nil {
			fmt.Println(err.Error())
			res := helpers.ResponseMsg(false, err.Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		// Match password input with password from db
		match := helpers.CheckPasswordHash(password, login.Password)

		if match {
			// Create a new random session token
			sessionToken := uuid.NewV4().String()
			// Set the token in the cache, along with the user whom it represents
			// The token has an expirt time of 3600 seconds (1 hour)
			_, err = helpers.InitCache().Do("SETEX", sessionToken, "3600", username)
			if err != nil {
				// If there is an error in setting the cache, return an internal server error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Finally, we set the client cookie for "session_token" as the session token we just generated
			// We also set an expiry time of 120 seconds, the same as the cache
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: time.Now().Add(3600 * time.Second),
			})

			statusCode := http.StatusOK
			w.WriteHeader(statusCode)
			res := helpers.ResponseMsg(true, login)
			json.NewEncoder(w).Encode(res)
		} else {
			body := "Username or password is wrong"
			res := helpers.ResponseMsg(false, body)
			json.NewEncoder(w).Encode(res)
		}
	} else {
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
	}
}
