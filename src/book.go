package main

import (
	"context"
	"fmt"

	"go.bcc.media/bibleserver/bibles"
	"go.bcc.media/bibleserver/proto"
)

func getBooks(ctx context.Context, bibleID string) ([]*proto.Book, error) {
	bible := bibles.Get(bibleID)
	if bible == nil {
		return nil, fmt.Errorf("Unable to find bible %s", bibleID)
	}

	row, err := bible.QueryContext(ctx, "SELECT book_number, long_name, short_name FROM books")
	if err != nil {
		return nil, err
	}

	defer row.Close()

	books := []*proto.Book{}

	for row.Next() {
		b := &proto.Book{}
		row.Scan(&b.Number, &b.LongName, &b.ShortName)
		b.Id = bookNumberToCanonical[b.Number]
		books = append(books, b)
	}

	return books, err
}
