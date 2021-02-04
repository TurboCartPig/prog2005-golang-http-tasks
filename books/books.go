package books

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Books interface {
	GetBook(id int) (Book, bool)
	InsertBook(id int, book Book)
}

type Book struct {
	Year   int    `json:"year"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type booksDB struct {
	books map[int]Book
}

func (db *booksDB) GetBook(id int) (Book, bool) {
	book, ok := db.books[id]
	return book, ok
}

// FIXME: What if the id belongs to another book?
func (db *booksDB) InsertBook(id int, book Book) {
	db.books[id] = book
}

func NewBooksDB() Books {
	books := &booksDB{
		map[int]Book{
			1: {1220, "What?", "Who?"},
			2: {2020, "What a year", "The survivors"},
		},
	}

	return books
}

func NewBookGetter(db *Books) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validation is done by chi and regex already
		id_string := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(id_string)

		book, ok := (*db).GetBook(id)
		if ok {
			w.Header().Add("Content-Type", "text/json")
			json.NewEncoder(w).Encode(book)
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
}
