package users

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"github.com/biFebriansyah/goback/src/interfaces"
)

type users_ctrl struct {
	repo interfaces.UserService
}

func NewCtrl(rep interfaces.UserService) *users_ctrl {
	return &users_ctrl{rep}
}

func (rep *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data.Send(w)

}

func (rep *users_ctrl) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claim_user := r.Context().Value("username")
	data, err := rep.repo.FindByUsername(claim_user.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data.Send(w)

}

func (rep *users_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	json.NewDecoder(r.Body).Decode(&data)
	result, err := rep.repo.Save(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)

}
