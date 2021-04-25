package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/rs/zerolog"
	"github.com/soheilhy/cmux"

	"go.bcc.media/bibleserver/log"
)

var bibles map[string]*sql.DB

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.ConfigureGlobalLogger(zerolog.DebugLevel)
	bibles = map[string]*sql.DB{}

	// TODO: Better path logic, potentially only a location and autoload all *.sqlite
	// Also we could load bibles on demand later
	db, _ := sql.Open("sqlite3", "../bibles/nb-1930.sqlite")
	defer db.Close()
	bibles["NB-1930"] = db

	log.L.Info().Msgf("Loaded %d bibles", len(bibles))

	// Create the main listener.

	l, err := net.Listen("tcp", (fmt.Sprintf(":%s", getEnv("PORT", "8000"))))
	if err != nil {
		log.L.Fatal().Err(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	//	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.Any())

	//go grpcS.Serve(grpcL)

	router := gin.Default()
	router.Use(logger.SetLogger(logger.Config{
		Logger: log.L,
	}))
	router.GET("v1/:bible/books", listBooks)
	router.GET("v1/:bible/:book/:chapter/:verse_from", getVersesHandler)
	router.GET("v1/:bible/:book/:chapter/:verse_from/:verse_to", getVersesHandler)

	go router.RunListener(httpL)
	m.Serve()
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

type verseResponse struct {
	Verses []Verse
}

func getVersesHandler(c *gin.Context) {
	bibleID := c.Param("bible")
	book, _ := strconv.ParseInt(c.Param("book"), 10, 32)
	chapter, _ := strconv.ParseInt(c.Param("chapter"), 10, 32)
	verseFrom, _ := strconv.ParseInt(c.Param("verse_from"), 10, 32)
	verseTo, _ := strconv.ParseInt(c.Param("verse_to"), 10, 32)

	if verseTo < verseFrom {
		verseTo = verseFrom
	}

	verses, err := getVerses(c.Request.Context(), bibles, bibleID, uint32(book), uint32(chapter), uint32(verseFrom), uint32(verseTo))
	if err != nil {
		log.L.Err(err)
		c.AbortWithStatus(500)
		return
	}

	if len(verses) == 0 {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, verseResponse{Verses: verses})
}
