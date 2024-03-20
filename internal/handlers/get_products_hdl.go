package handlers

import (
	"net/http"

	"github.com/leonardodelira/first-templ-go/internal/templates"
)

type ProducsHandler struct {
}

func NewProducsHandler() *ProducsHandler {
	return &ProducsHandler{}
}

func (h *ProducsHandler) Process(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("search")

	productsView := templates.Products(input)
	err := templates.Layout(productsView).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
