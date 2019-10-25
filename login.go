package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/julienschmidt/httprouter"
)

var indexTemplate *template.Template = nil

func ShowLoginForm(message string, parentRedirectTo string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if indexTemplate == nil {
			tplString, err := packr.NewBox("./templates").FindString("index.tpl")
			if err != nil {
				log.Fatal(err)
			}
			_t, err := template.New("index.tpl").Parse(tplString)
			if err != nil {
				log.Fatal(err)
			}
			indexTemplate = _t
		}
		t := indexTemplate

		redirectTo := parentRedirectTo
		if redirectTo == "" {
			redirectTo = r.Header.Get("X-GATEWAY-REDIRECT-TO")
		}

		templateMap := map[string]string{
			"RedirectTo": redirectTo,
			"Message":    message,
		}

		err := t.Execute(w, templateMap)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	config := LoadConfig()

	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = r.Header.Get("X-GATEWAY-REDIRECT-TO")
	}

	l, err := ConnectLDAP()
	if err != nil {
		log.Printf("LDAP Error : %+v\n", err)
		if config.LDAP.IsShowErrorDetails {
			ShowLoginForm(fmt.Sprintf("LDAP Error : %s", err.Error()), redirectTo)(w, r, p)
		} else {
			ShowLoginForm("Failed to connect to LDAP server.", redirectTo)(w, r, p)
		}
		return
	}
	defer l.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		ShowLoginForm("Invalid value!", redirectTo)(w, r, p)
		return
	}

	err = AuthLDAP(l, username, password)
	if err != nil {
		log.Printf("Auth Error (%s) : %+v\n", username, err)
		if config.LDAP.IsShowErrorDetails {
			ShowLoginForm(fmt.Sprintf("Failed to login : %s", err.Error()), redirectTo)(w, r, p)
		} else {
			ShowLoginForm("Failed to login.", redirectTo)(w, r, p)
		}
		return
	}

	// log.Printf("Logged in : %s\n", username)

	// Regenerate SessionID to prevent session fixitation attack
	session := GetSessionManager().Load(r)
	session.PutBool(w, "isLoggedIn", true)
	session.PutString(w, "username", username)
	session.RenewToken(w)

	http.Redirect(w, r, redirectTo, http.StatusSeeOther)
}

func CookieLoginTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := GetSessionManager().Load(r)
	isLoggedIn, _ := session.GetBool("isLoggedIn")
	username, _ := session.GetString("username")
	if !isLoggedIn {
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.Header().Set("X-AUTH-USERNAME", username)
		w.WriteHeader(http.StatusOK)
	}
}
