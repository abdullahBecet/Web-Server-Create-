package home

import (
	"net/http"

	"github.com/go-chi/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"hello": "World",
	})
}
