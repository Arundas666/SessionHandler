package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var Tpl *template.Template
var store = sessions.NewCookieStore([]byte("Very-Secret-Key"))

func SetSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["role"] = "admin"
	session.Options.MaxAge = 300
	session.Save(r, w)
	fmt.Println("Session created")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["role"] = "user"
	session.Options.MaxAge = 300
	session.Save(r, w)
	fmt.Println("Session's Value Changed")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["role"] = ""
	session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Println("Session deleted")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Home(w http.ResponseWriter, r *http.Request) {
	Tpl, _ = template.ParseGlob("template/*.html")
	session, _ := store.Get(r, "cookie-name")
	if role, ok := session.Values["role"].(string); ok && role == "admin" {
		Tpl.ExecuteTemplate(w, "admin.html", nil)

		return
	} else if role, ok := session.Values["role"].(string); ok && role == "user" {
		Tpl.ExecuteTemplate(w, "user.html", nil)

		return
	}
	Tpl.ExecuteTemplate(w, "login.html", nil)

	// if role:= session.Values["role"].(string); role == "" {

}
