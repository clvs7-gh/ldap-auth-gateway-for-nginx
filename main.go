package main

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/ldap.v2"
)

var (
	Version  string
	Revision string
)

func main() {
	log.Println("Gateway has been started.")
	if Version != "" && Revision != "" {
		log.Printf("Version : %s , Revision : %s\n", Version, Revision)
	}

	// Load config in order to caching & check error
	_ = LoadConfig()
	// Set default timeout for LDAP
	ldap.DefaultTimeout = 10 * time.Second
	// Initialize session
	InitSession()

	http.ListenAndServe(":18080", GetRouter())
}
