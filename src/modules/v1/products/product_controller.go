package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"github.com/biFebriansyah/goback/src/helpers"
	"github.com/gorilla/mux"
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

	helpers.New(data, 200, false).Send(w)
}

func (rep *product_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	theId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data, err := rep.repo.FindById(theId)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	helpers.New(data, 200, false).Send(w)
}

func (rep *product_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uploads := r.Context().Value("file")
	var data models.Product
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data.Image = uploads.(string)
	result, err := rep.repo.Add(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *product_ctrl) TESTPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uploads := r.Context().Value("file")

	w.Write([]byte(uploads.(string)))
}
