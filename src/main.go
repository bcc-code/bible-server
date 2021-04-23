package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/rs/zerolog"

	"go.bcc.media/bibleserver/log"
)

var bibles map[string]*sql.DB

func main() {
	log.ConfigureGlobalLogger(zerolog.DebugLevel)
	bibles = map[string]*sql.DB{}

	db, _ := sql.Open("sqlite3", "../bibles/nb-1930.sqlite")
	defer db.Close()
	bibles["NB-1930"] = db

	router := gin.Default()
	router.Use(logger.SetLogger(logger.Config{
		Logger: log.L,
	}))
	router.GET("v1/:bible/books", listBooks)
	router.GET("v1/:bible/:book/:chapter/:verse_from", getSingleVerse)
	router.GET("v1/:bible/:book/:chapter/:verse_from/:verse_to", getMultipleVerses)
	router.Run(":8080")
}

// Book represents a book in the bible
type Book struct {
	// This shoudl be the canonical english version as defined in TODO.
	// The reason for this is so that we can have huma readable canonical
	// representations, f.ex. `1Pet 2/7-8`
	ID string

	Number    uint16 // Mostly for sorting
	LongName  string // Localized long name
	ShortName string // Localized short name
}

func listBooks(c *gin.Context) {
	bibleID := c.Param("bible")

	var bible *sql.DB
	if b, ok := bibles[bibleID]; ok {
		bible = b
	} else {
		c.AbortWithStatus(404)
	}

	row, err := bible.QueryContext(c.Request.Context(), "SELECT book_number, long_name, short_name FROM books")
	if err != nil {
		log.L.Fatal().Err(err)
	}

	defer row.Close()

	books := []Book{}

	for row.Next() { // Iterate and fetch the records from result cursor
		b := Book{}
		row.Scan(&b.Number, &b.LongName, &b.ShortName)
		books = append(books, b)
	}

	c.JSON(http.StatusOK, books)
}

type Verse struct {
	Number    uint32
	Text      string
	Footnotes []string
}

func getVerses(ctx context.Context, bibleID string, book, chapter, verseFrom, verseTo uint32) ([]Verse, error) {
	var bible *sql.DB
	if b, ok := bibles[bibleID]; ok {
		bible = b
	} else {
		return nil, fmt.Errorf("Bible %s not found", bibleID)
	}

	row, err := bible.QueryContext(ctx, "SELECT verse, text FROM verses WHERE book_number = ? AND chapter = ? AND verse >= ? AND verse <= ?", book, chapter, verseFrom, verseTo)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	verses := []Verse{}

	for row.Next() { // Iterate and fetch the records from result cursor
		v := Verse{}
		row.Scan(&v.Number, &v.Text)
		verses = append(verses, v)
	}

	return verses, nil
}

func getSingleVerse(c *gin.Context) {
	bibleID := c.Param("bible")
	book, _ := strconv.ParseInt(c.Param("book"), 10, 32)
	chapter, _ := strconv.ParseInt(c.Param("chapter"), 10, 32)
	verseFrom, _ := strconv.ParseInt(c.Param("verse_from"), 10, 32)

	verses, err := getVerses(c.Request.Context(), bibleID, uint32(book), uint32(chapter), uint32(verseFrom), uint32(verseFrom))
	if err != nil {
		log.L.Err(err)
		c.AbortWithStatus(500)
		return
	}

	if len(verses) == 0 {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, verses)
}

func getMultipleVerses(c *gin.Context) {
	bibleID := c.Param("bible")
	book, _ := strconv.ParseInt(c.Param("book"), 10, 32)
	chapter, _ := strconv.ParseInt(c.Param("chapter"), 10, 32)
	verseFrom, _ := strconv.ParseInt(c.Param("verse_from"), 10, 32)
	verseTo, _ := strconv.ParseInt(c.Param("verse_to"), 10, 32)

	verses, err := getVerses(c.Request.Context(), bibleID, uint32(book), uint32(chapter), uint32(verseFrom), uint32(verseTo))
	if err != nil {
		log.L.Err(err)
		c.AbortWithStatus(500)
		return
	}

	if len(verses) == 0 {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, verses)
}
