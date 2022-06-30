package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FileUpload(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		var pathfile string

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			fmt.Fprint(w, err.Error())
		}

		for _, fheaders := range r.MultipartForm.File {
			for _, hdr := range fheaders {
				infile, err := hdr.Open()
				if err != nil {
					fmt.Fprint(w, err.Error())
				}

				pathfile = "./upload/images/" + hdr.Filename
				outfile, err := os.Create(pathfile)
				if err != nil {
					fmt.Fprint(w, err.Error())
				}

				defer outfile.Close()
				io.Copy(outfile, infile)
			}
		}

		log.Println("Upload Middleware Pass")

		// share context to controller
		ctx := context.WithValue(r.Context(), "file", pathfile)

		// Serve the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
