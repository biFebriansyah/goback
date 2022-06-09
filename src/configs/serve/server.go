package serve

import (
	"log"
	"net/http"
	"os"

	"github.com/biFebriansyah/goback/src/routers"
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

		if pr := os.Getenv("APP_PORT"); pr != "" {
			addrs = "0.0.0.0:" + pr
		}

		log.Println("App running on " + addrs)

		if err := http.ListenAndServe(addrs, mainRoute); err != nil {
			return err
		}

		return nil

	} else {
		return err
	}
}
