package main

import (
	"time"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/memstore"
)

var sessionManager *scs.Manager

func InitSession() {
	sessionManager = scs.NewManager(memstore.New(10 * time.Minute))
	sessionManager.HttpOnly(true)
	sessionManager.Secure(true)
	sessionManager.Lifetime(6 * time.Hour)
	sessionManager.Name("GATEWAY_SESSIONID")
	sessionManager.Path("/")
	sessionManager.SameSite("Lax")
}

func GetSessionManager() *scs.Manager {
	return sessionManager
}
