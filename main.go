package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Serve index page on all unhandled routes
	r.Handle("/", Handler())
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("static/styles/"))))

	log.Fatal(http.ListenAndServe(port, r))

}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "./static/404.html")

}

func Handler() http.Handler {
	return http.FileServer(http.Dir("static/"))
}
