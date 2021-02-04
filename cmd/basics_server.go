package main

import (
	"net/http"
	"strings"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2021-workspace/denniskr/golang-http-tasks/books"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Say hello by writing a http response
func sayHello(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")             // Extract the name parameter from the http request
	name = strings.Title(strings.ToLower(name)) // Make the names pretty
	w.Write([]byte("Hello " + name + "!"))      // Say hello to the client by writing a string to the response
}

func main() {
	db := books.NewBooksDB()

	// Chi handles most of point 2 for us,
	// But we have to specify some regex for the last part.
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello/{name:[a-zA-Z-]+}", sayHello)
	r.Get("/books/{id:[0-9]+}", books.NewBookGetter(&db))

	http.ListenAndServe(":3000", r)
}
