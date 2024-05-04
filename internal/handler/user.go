package handler

import (
	"log"
	"net/http"

	"github.com/Michael-Sjogren/gotempl/internal/model"
	"github.com/Michael-Sjogren/gotempl/internal/views/pages"
)

type UserHandler struct {
	UserModel *model.UserRepo
}

func (h *UserHandler) HandleUsersPageView(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserModel.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = pages.UsersPage(users).Render(r.Context(), w)
	if err != nil {
		log.Fatal(err)
	}
}
