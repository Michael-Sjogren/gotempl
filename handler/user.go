package handler

import (
	"log"
	"net/http"

	"github.com/Michael-Sjogren/gotempl/mytypes"
	"github.com/Michael-Sjogren/gotempl/views/pages"
)

type UserHandler struct{}

func (h *UserHandler) HandleUsersPageView(w http.ResponseWriter, r *http.Request) {
	err :=
		pages.UsersPage([]*mytypes.User{
			{Username: "test"},
		}).Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}
