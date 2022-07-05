// https://www.gvoidy.cn/index.php/archives/284/

package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello world!"))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		return
	}

}
