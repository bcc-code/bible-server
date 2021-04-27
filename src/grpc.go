package main

import (
	"context"

	"go.bcc.media/bibleserver/log"
	"go.bcc.media/bibleserver/proto"
)

// GRPCBibleServer represents the main GRPC Server
type GRPCBibleServer struct {
	proto.UnimplementedBibleServerServer
}

// GetVerses returns a list of verses
func (s GRPCBibleServer) GetVerses(ctx context.Context, req *proto.GetVersesRequest) (*proto.GetVersesResponse, error) {
	bookID := canonicalBookToNumber[req.BookId]
	verses, err := getVerses(ctx, req.BibleId, bookID, req.Chapter, req.VerseFrom, req.VerseTo)
	if err != nil {
		log.L.Error().Err(err).Msg("")
		return nil, err
	}

	return &proto.GetVersesResponse{
		BibleId: req.BibleId,
		BookId:  req.BookId,
		Chapter: req.Chapter,
		Verses:  verses.ToProto(),
	}, nil
}

// ListBooks in the selecteb bible
func (s GRPCBibleServer) ListBooks(ctx context.Context, req *proto.ListBooksRequest) (*proto.ListBooksResponse, error) {
	books, err := getBooks(ctx, req.BibleId)
	if err != nil {
		log.L.Error().Err(err).Msg("")
		return nil, err
	}

	return &proto.ListBooksResponse{BibleId: req.BibleId, Books: books}, nil
}
