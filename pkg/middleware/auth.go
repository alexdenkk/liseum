package middleware

import (
	"alexdenkk/liseum/pkg/token"
	"alexdenkk/liseum/pkg/tools"
	"net/http"
)

// HandlerFunc type - type for handler functions using Auth Mmiddleware
type HandlerFunc func(http.ResponseWriter, *http.Request, *token.Claims)

// Auth - function for checking user authorization
func (mw *Middleware) Auth(f HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t, err := tools.GetCookie(r, "token")

		if err != nil {
			print("cant get cookie")
			http.Redirect(w, r, "/user/login/", http.StatusMovedPermanently)
			return
		}

		claims, err := token.ParseJWT(t, mw.SignKey)

		if err != nil {
			print("cant parse cookie")
			http.Redirect(w, r, "/user/login/", http.StatusMovedPermanently)
			return
		}

		f(w, r, claims)
	})
}

// NotAuth - function for checking user non-authorization
func (mw *Middleware) NotAuth(f http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t, err := tools.GetCookie(r, "token")

		if err == nil {
			_, err := token.ParseJWT(t, mw.SignKey)

			if err == nil {
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return
			}
		}

		f(w, r)
	})
}
