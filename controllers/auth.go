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

var register = &models.Register{}

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
		register.Username = username
		register.Password = hashedPassword

		_, err = db.Exec(helpers.Query["register"], register.Username, register.Password)
		if err != nil {
			fmt.Println(err.Error())
			res := helpers.ResponseMsg(false, err.Error())
			statusCode := http.StatusBadRequest
			stdout := helpers.GenerateStdout(register, "application/json", statusCode, res)
			fmt.Println(stdout)
			json.NewEncoder(w).Encode(res)
			return
		}

		body := "Register success"
		statusCode := http.StatusOK
		w.WriteHeader(statusCode)
		res := helpers.ResponseMsg(true, body)
		json.NewEncoder(w).Encode(res)

		stdout := helpers.GenerateStdout(register, "application/json", statusCode, res)
		fmt.Println(stdout)

	} else {
		statusCode := http.StatusBadRequest
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
		stdout := helpers.GenerateStdout(register, "application/json", statusCode, res)
		fmt.Println(stdout)
	}
}

var login = models.Login{}

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

		username, password := r.FormValue("username"), r.FormValue("password")
		err = db.
			QueryRow(helpers.Query["login"], username).
			Scan(&login.Username, &login.Password)

		if err != nil {
			res := helpers.ResponseMsg(false, err.Error())
			statusCode := http.StatusBadRequest
			stdout := helpers.GenerateStdout(login, "application/json", statusCode, res)
			fmt.Println(stdout)
			json.NewEncoder(w).Encode(res)
			return
		}

		match := helpers.CheckPasswordHash(password, login.Password)

		if match {
			sessionToken := uuid.NewV4().String()
			_, err = helpers.InitCache().Do("SETEX", sessionToken, "3600", username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: time.Now().Add(3600 * time.Second),
			})

			statusCode := http.StatusOK
			w.WriteHeader(statusCode)
			res := helpers.ResponseMsg(true, login)
			json.NewEncoder(w).Encode(res)

			stdout := helpers.GenerateStdout(login, "application/json", statusCode, res)
			fmt.Println(stdout)
		} else {
			statusCode := http.StatusBadRequest
			body := "Username or password is wrong"
			res := helpers.ResponseMsg(false, body)
			json.NewEncoder(w).Encode(res)
			stdout := helpers.GenerateStdout(login, "application/json", statusCode, res)
			fmt.Println(stdout)
		}
	} else {
		statusCode := http.StatusBadRequest
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
		stdout := helpers.GenerateStdout(login, "application/json", statusCode, res)
		fmt.Println(stdout)
	}
}
