package web

import (
	"a21hc3NpZ25tZW50/client"
	"embed"
	"fmt"
	"net/http"
	"html/template"
	"path"
)

type AuthWeb interface {
	Login(w http.ResponseWriter, r *http.Request)
	LoginProcess(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	RegisterProcess(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authWeb struct {
	userClient client.UserClient
	embed      embed.FS
}

func NewAuthWeb(userClient client.UserClient, embed embed.FS) *authWeb {
	return &authWeb{userClient, embed}
}

func (a *authWeb) Login(w http.ResponseWriter, r *http.Request) {
	var filelogin = path.Join("views", "auth", "login.html")
	var fileheader = path.Join("views", "general", "header.html")
	webLogin := template.Must(template.ParseFS(a.embed, filelogin, fileheader))

	var data = map[string]interface{}{
		"title": "Login",
	}	

	Alert := webLogin.Execute(w, data)

	if Alert != nil {
		http.Error(w, Alert.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: answer here
}

func (a *authWeb) LoginProcess(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	userId, status, err := a.userClient.Login(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		http.SetCookie(w, &http.Cookie{
			Name:   "user_id",
			Value:  fmt.Sprintf("%d", userId),
			Path:   "/",
			MaxAge: 31536000,
			Domain: "",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (a *authWeb) Register(w http.ResponseWriter, r *http.Request) {
	var filelogin = path.Join("views", "auth", "register.html")
	var fileheader = path.Join("views", "general", "header.html")
	webLogin := template.Must(template.ParseFS(a.embed, filelogin, fileheader))

	var data = map[string]interface{}{
		"title": "Register",
	}	

	Alert := webLogin.Execute(w, data)

	if Alert != nil {
		http.Error(w, Alert.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: answer here
}

func (a *authWeb) RegisterProcess(w http.ResponseWriter, r *http.Request) {
	fullname := r.FormValue("fullname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	userId, status, err := a.userClient.Register(fullname, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		http.SetCookie(w, &http.Cookie{
			Name:   "user_id",
			Value:  fmt.Sprintf("%d", userId),
			Path:   "/",
			MaxAge: 31536000,
			Domain: "",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}

func (a *authWeb) Logout(w http.ResponseWriter, r *http.Request) {
	
	http.SetCookie(w, &http.Cookie{
		Name	:   "user_id",
		Value 	: 	 "",
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// TODO: answer here
}
