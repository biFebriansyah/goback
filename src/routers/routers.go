package routers

import (
	"net/http"

	database "github.com/biFebriansyah/goback/src/database/gorm"
	"github.com/biFebriansyah/goback/src/modules/v1/auth"
	"github.com/biFebriansyah/goback/src/modules/v1/products"
	"github.com/biFebriansyah/goback/src/modules/v1/users"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	nRelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("goback"),
		newrelic.ConfigLicense("1c5eebe1b195c3bad62de324f3aeceaddfe9NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		return nil, err
	}

	mainRoute.Use(nrgorilla.Middleware(nRelic))

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc(newrelic.WrapHandleFunc(nRelic, "/relic", relicHandler)).Methods("GET")
	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")

	products.New(mainRoute, db)
	users.New(mainRoute, db)
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}

func relicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello newrelic"))
}
