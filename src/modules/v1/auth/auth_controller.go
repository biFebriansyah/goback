package auth

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"github.com/biFebriansyah/goback/src/interfaces"
)

type auth_ctrl struct {
	rep interfaces.AuthService
}

func NewCtrk(rep interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{rep}
}

func (re *auth_ctrl) Sigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	json.NewDecoder(r.Body).Decode(&data)
	result, err := re.rep.Login(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)

}
