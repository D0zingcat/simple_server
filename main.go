package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	user     string
	password string
	dir      string
)

type authenticationMiddleware struct{}

func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xTokenSession := r.Header.Get("X-Session-Token")
		if user+":"+password == xTokenSession {
			log.Printf("Authenticated user %s %s\n", user, password)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {
	flag.StringVar(&user, "u", "d0zingcat", "username")
	flag.StringVar(&password, "p", "Hello@World11235", "token")
	flag.StringVar(&dir, "d", "./static", "static folder location")
	flag.Parse()
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	amw := authenticationMiddleware{}
	r.Use(amw.Middleware)
	log.Fatal(srv.ListenAndServe())

}
