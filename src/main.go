package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/sirupsen/logrus"

	"go.bcc.media/bibleserver/log"

	"encoding/json"
	"fmt"

	"net/http"
)

var bibles map[string]*sql.DB

func main() {
	log.ConfigureGlobalLogger(logrus.DebugLevel)
	bibles = map[string]*sql.DB{}

	db, _ := sql.Open("sqlite3", "../bibles/nb-1930.sqlite")
	defer db.Close()
	bibles["NB-1930"] = db

	http.HandleFunc("/books", listBooks)
	http.ListenAndServe(":8001", nil)
}

// Book represents a book in the bible
type Book struct {
	Number    uint16
	LongName  string
	ShortName string
	ID        string
}

func listBooks(w http.ResponseWriter, req *http.Request) {
	db := bibles["NB-1930"]

	row, err := db.Query("SELECT book_number, long_name, short_name FROM books")
	if err != nil {
		log.L.Fatal(err)
	}
	defer row.Close()

	books := []Book{}

	for row.Next() { // Iterate and fetch the records from result cursor
		b := Book{}
		row.Scan(&b.Number, &b.LongName, &b.ShortName)
		books = append(books, b)
	}

	json.NewEncoder(w).Encode(books)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")

}
