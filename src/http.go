package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library

	"go.bcc.media/bibleserver/log"
)

func listBooks(c *gin.Context) {
	bibleID := c.Param("bible")

	books, err := getBooks(c.Request.Context(), bibleID)

	if err != nil {
		log.L.Error().Err(err).Msg("")
		c.AbortWithStatus(400)
	}

	c.JSON(http.StatusOK, books)
}

type verseResponse struct {
	Verses []Verse
}

func getVersesHandler(c *gin.Context) {
	bibleID := c.Param("bible")
	book := c.Param("book")
	chapter, _ := strconv.ParseInt(c.Param("chapter"), 10, 32)
	verseFrom, _ := strconv.ParseInt(c.Param("verse_from"), 10, 32)
	verseTo, _ := strconv.ParseInt(c.Param("verse_to"), 10, 32)

	if verseTo < verseFrom {
		verseTo = verseFrom
	}

	verses, err := getVerses(c.Request.Context(), bibleID, canonicalBookToNumber[book], uint32(chapter), uint32(verseFrom), uint32(verseTo))
	if err != nil {
		log.L.Err(err).Msg("")
		c.AbortWithStatus(500)
		return
	}

	if len(verses) == 0 {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, verseResponse{Verses: verses})
}
