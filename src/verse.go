package main

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"go.bcc.media/bibleserver/proto"
)

var (
	tagRegxp       *regexp.Regexp
	footnoteRegexp *regexp.Regexp
)

func init() {
	tagRegxp = regexp.MustCompile("<[a-zA-Z/]+>")
	footnoteRegexp = regexp.MustCompile("{(\\*+) (.+?)}")
}

type VerseList []Verse

func (vl VerseList) ToProto() []*proto.Verse {
	pl := []*proto.Verse{}
	for _, v := range vl {
		pl = append(pl, v.ToProto())
	}

	return pl
}

// Verse represents a single verse from the bible
type Verse struct {
	Number    uint32
	Text      string
	Footnotes map[string]string
}

func (v Verse) removeTags() Verse {
	v.Text = tagRegxp.ReplaceAllString(v.Text, "")
	return v
}

func (v Verse) parseFootnotes() Verse {
	matches := footnoteRegexp.FindAllStringSubmatch(v.Text, -1)
	for _, note := range matches {
		v.Footnotes[note[1]] = note[2]
	}

	if len(matches) > 0 {
		v.Text = footnoteRegexp.ReplaceAllString(v.Text, "")
	}

	return v
}

func (v Verse) ToProto() *proto.Verse {
	return &proto.Verse{
		Number:    v.Number,
		Text:      v.Text,
		Footnotes: v.Footnotes,
	}
}

// NewVerse constructs and parses the text for a verse
func NewVerse(number uint32, text string) Verse {
	v := Verse{
		Number:    number,
		Text:      text,
		Footnotes: map[string]string{},
	}

	v = v.removeTags()
	v = v.parseFootnotes()
	v.Text = strings.TrimSpace(v.Text)

	return v
}

func getVerses(ctx context.Context, bibles map[string]*sql.DB, bibleID string, book, chapter, verseFrom, verseTo uint32) (VerseList, error) {
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

	for row.Next() {
		var n uint32
		var t string
		row.Scan(&n, &t)

		v := NewVerse(n, t)
		verses = append(verses, v)
	}

	return verses, nil
}
