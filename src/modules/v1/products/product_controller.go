package products

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product_ctrl struct {
	repo *product_repo
}

func NewCtrl(rep *product_repo) *product_ctrl {
	return &product_ctrl{rep}
}

func (rep *product_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *product_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data Product
	json.NewDecoder(r.Body).Decode(&data)

	result, err := rep.repo.Add(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}
