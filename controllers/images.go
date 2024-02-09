package controllers

import (
	"alexdenkk/liseum/model"
	"alexdenkk/liseum/pkg/token"
	"context"
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateImage(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	if r.Method == http.MethodPost {
		className := mux.Vars(r)["class"]

		err := r.ParseMultipartForm(10 << 20)

		if err != nil {
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		file, handler, err := r.FormFile("image")

		if err != nil {
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		defer file.Close()

		ctx := context.WithValue(context.Background(), "request", r)

		class, err := h.Repository.GetClassByName(ctx, className)

		if err != nil {
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		if len(r.FormValue("label")) < 10 || len(r.FormValue("name")) < 6 {
			http.Redirect(w, r, r.URL.Path+"?err="+errors.New("Слишком мало текста").Error(), http.StatusMovedPermanently)
			return
		}

		img := &model.Image{
			ClassID: class.ID,
			Label:   r.FormValue("label"),
			Name:    r.FormValue("name"),
		}

		err = h.Repository.CreateImage(ctx, img)

		if err != nil {
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		img.FileName = "/static/img/" + class.Name + "/" + strconv.Itoa(int(img.ID)) +
			"." + strings.Split(handler.Filename, ".")[1]

		err = h.Repository.UpdateImage(ctx, img)

		if err != nil {
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		fileValue, err := io.ReadAll(file)

		if err != nil {
			h.Repository.DeleteImage(ctx, img.ID)
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		err = os.WriteFile("web"+img.FileName, fileValue, 0666)

		if err != nil {
			h.Repository.DeleteImage(ctx, img.ID)
			http.Redirect(w, r, r.URL.Path+"?err="+err.Error(), http.StatusMovedPermanently)
			return
		}

		http.Redirect(w, r, "/img/for/"+className+"/", http.StatusMovedPermanently)

	} else if r.Method == http.MethodGet {
		isErr := r.URL.Query().Get("err")
		tmpl := template.Must(template.ParseFiles("./web/html/create_image.html"))
		tmpl.Execute(w, isErr)
	}
}

func (h *Handler) GetImagesFor(w http.ResponseWriter, r *http.Request) {
	className := mux.Vars(r)["class"]

	ctx := context.WithValue(context.Background(), "request", r)

	class, err := h.Repository.GetClassByName(ctx, className)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	images, err := h.Repository.GetImagesFor(ctx, class.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/html/images.html"))
	tmpl.Execute(w, images)
}
