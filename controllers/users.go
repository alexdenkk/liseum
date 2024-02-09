package controllers

import (
	"alexdenkk/liseum/pkg/hash"
	"alexdenkk/liseum/pkg/token"
	"alexdenkk/liseum/pkg/tools"
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Login - Gateway layer function for user authorization
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		ctx := context.WithValue(context.Background(), "request", r)

		user, err := h.Repository.GetUserByLogin(ctx, r.FormValue("login"))

		if err != nil {
			http.Redirect(w, r, "/user/login/", http.StatusMovedPermanently)
			return
		}

		// check password
		if user.Password != hash.Hash(r.FormValue("password")) {
			http.Redirect(w, r, "/user/login/", http.StatusMovedPermanently)
			return
		}

		// generating token
		claims := token.Claims{
			ID:    user.ID,
			Login: user.Login,

			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
			},
		}

		token, err := token.GenerateJWT(claims, h.SignKey)

		if err != nil {
			http.Redirect(w, r, "/user/login/", http.StatusMovedPermanently)
			return
		}

		tools.SetCookie(w, "token", token)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	} else if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("./web/html/login.html")
		tmpl.Execute(w, nil)
	}
}
