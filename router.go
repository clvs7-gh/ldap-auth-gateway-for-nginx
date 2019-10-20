package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/gobuffalo/packr"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/login", Login)
	router.GET("/api/v1/cookie/loggedin", CookieLoginTest)
	router.ServeFiles("/assets/*filepath", packr.NewBox("./assets"))
	router.NotFound = http.HandlerFunc(NotFound)

	return router
}
