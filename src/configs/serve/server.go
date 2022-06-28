package serve

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/biFebriansyah/goback/src/routers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "0.0.0.0:8080"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = ":" + pr
		}

		log.Println("Running With" + runtime.GOOS)
		log.Println("App running on " + addrs)

		// t := cors.New(cors.Options{
		// 	AllowedOrigins: []string{"*"},                              // All origins
		// 	AllowedMethods: []string{"GET", "POST", "HEAD", "OPTIONS"}, // Allowing only get, just an example
		// })

		t := cors.AllowAll()

		if err := http.ListenAndServe(addrs, t.Handler(mainRoute)); err != nil {
			return err
		}

		return nil

	} else {
		return err
	}
}
