package controllers

import (
	"alexdenkk/liseum/model"
	"alexdenkk/liseum/pkg/token"
	"context"
	"html/template"
	"net/http"
	"os"
)

func (h *Handler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "request", r)

	classes, err := h.Repository.GetAllClasses(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/html/classes.html"))
	tmpl.Execute(w, classes)
}

func (h *Handler) CreateClass(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		ctx := context.WithValue(context.Background(), "request", r)

		name := r.FormValue("name")

		class := &model.Class{
			Name: name,
		}

		err := h.Repository.CreateClass(ctx, class)

		if err != nil {
			http.Redirect(w, r, "", http.StatusMovedPermanently)
			return
		}

		err = os.Mkdir("web/static/img/"+class.Name, 0777)

		if err != nil {
			http.Redirect(w, r, "", http.StatusMovedPermanently)
			return
		}

		http.Redirect(w, r, "/class/"+name+"/", http.StatusMovedPermanently)

	} else if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./web/html/create_class.html"))
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	ctx := context.WithValue(context.Background(), "request", r)

	classes, err := h.Repository.GetAllClasses(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/html/admin.html"))
	tmpl.Execute(w, classes)
}
