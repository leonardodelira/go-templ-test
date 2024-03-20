package handlers

import (
	"net/http"

	"github.com/leonardodelira/first-templ-go/internal/templates"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Process(w http.ResponseWriter, r *http.Request) {
	home := templates.Home()
	err := templates.Layout(home).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
