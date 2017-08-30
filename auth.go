package main

import (
	"log"
	"net/http"
	//"encoding/json"
)

func isAuth(req *http.Request) bool {
	// Get a session. Get() always returns a session, even if empty.
	session, err := sessionStore.Get(req, "session-auth")
	if err != nil {
		log.Println("Failed to read session. Reset.", err.Error(), http.StatusInternalServerError)
		return false
	}
	return (session.Values["logged"] == "true")
}

func LoginRoute(res http.ResponseWriter, req *http.Request) {
	res.Write(pageBuff["login"])
}
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form["username"][0]
	password := req.Form["password"][0]
	if username == "admin" && password == "123456" {
		SetSecureSessionLogged(res, req, "true")
		res.Write([]byte("true"))
	} else {
		res.Write([]byte("false"))
	}
}
func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	SetSecureSessionLogged(res, req, "false")
}

func SetSecureSessionLogged(w http.ResponseWriter, r *http.Request, value string) {
	// Get a session. Get() always returns a session, even if empty.
	session, _ := sessionStore.Get(r, "session-auth")
	// Set some session values.
	session.Values["logged"] = value
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}
