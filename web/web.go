package web

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

func StartRouter() *chi.Mux {
	siteInfo := SiteInfo{}
	siteInfo.Dir = "./sites"

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		firstSec := strings.Split(r.Host, ".")[0]
		// w.Write([]byte(firstSec))
		folder := filepath.Join(siteInfo.Dir, firstSec, "public")
		// router.
		fs := http.FileServer(http.Dir(folder))
		fs.ServeHTTP(w, r)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		firstSec := strings.Split(r.Host, ".")[0]
		// w.Write([]byte(firstSec))
		folder := filepath.Join(siteInfo.Dir, firstSec, "public")
		// router.
		fs := http.FileServer(http.Dir(folder))
		fs.ServeHTTP(w, r)
	})

	return router
}
