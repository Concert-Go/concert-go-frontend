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

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)


	r.Handle("/", routeHandler())
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("static/styles/"))))

	log.Fatal(http.ListenAndServe(port, r))

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "./static/404.html")

}

func routeHandler() http.Handler {
	return http.FileServer(http.Dir("static/"))
}
