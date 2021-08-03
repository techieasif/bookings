package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)


//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: true	,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionLoad loads and saves sessions on every request.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}