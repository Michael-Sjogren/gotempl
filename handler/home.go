package handler

import (
	"log"
	"net/http"

	"github.com/Michael-Sjogren/gotempl/views/pages"
)

type HomeHandler struct{}

func (h *HomeHandler) HandlerHomePageView(w http.ResponseWriter, r *http.Request) {
	if err := pages.HomePage().Render(r.Context(), w); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
