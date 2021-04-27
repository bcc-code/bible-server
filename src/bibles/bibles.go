package bibles

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"regexp"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library

	"go.bcc.media/bibleserver/log"
)

var (
	bibles        map[string]*sql.DB
	bibleIDRegexp *regexp.Regexp
)

func init() {
	bibleIDRegexp = regexp.MustCompile("/([^/]+).bible")
}

// Get open bible
func Get(bibleID string) *sql.DB {
	if db, ok := bibles[bibleID]; ok {
		return db
	}

	return nil
}

// MustLoadBibles that are available in the path
// It may panic if there are erorrs
func MustLoadBibles(basePath string) {
	bibles = map[string]*sql.DB{}

	bibleFiles, err := filepath.Glob(fmt.Sprintf("%s/*.bible", basePath))
	if err != nil {
		log.L.Fatal().Err(err).Msg("Error globbing for bibles")
	}

	for _, biblePath := range bibleFiles {
		db, err := sql.Open("sqlite3", fmt.Sprintf("%s?immutable=true", biblePath))
		if err != nil {
			log.L.Error().Err(err).Msgf("Unable to read bible at path %s", biblePath)
			continue
		}

		bibleID := bibleIDRegexp.FindStringSubmatch(biblePath)
		bibles[bibleID[1]] = db
	}

	log.L.Info().Msgf("Loaded %d bibles", len(bibles))
}

// CloseBibles that are open
func CloseBibles() {
	for id, db := range bibles {
		err := db.Close()
		if err != nil {
			log.L.Error().Err(err).Msgf("Error while closing %s", id)
		}
	}
}
