package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type SiteInfo struct {
	Dir string
}

func main() {
	siteInfo := SiteInfo{}
	siteInfo.Dir = "../sites"

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		firstSec := strings.Split(r.Host, ".")[0]

		folder := filepath.Join(siteInfo.Dir, firstSec)
		// router.
		fs := http.FileServer(http.Dir(folder))
		fs.ServeHTTP(w, r)
	})

}
