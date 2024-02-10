package tools

import (
	"fmt"
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, name, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "http://147.45.108.41/",
		Expires:  time.Now().Add(1000 * time.Hour),
		HttpOnly: false,
		Secure:   false,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	fmt.Println(cookie)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
