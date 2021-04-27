package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go.bcc.media/bibleserver/bibles"
	"go.bcc.media/bibleserver/log"
	"go.bcc.media/bibleserver/proto"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.ConfigureGlobalLogger(zerolog.DebugLevel)
	biblePath := getEnv("BIBLE_PATH", "../bibles/")
	bibles.MustLoadBibles(biblePath)
	defer bibles.CloseBibles()

	// Create the main listener.

	l, err := net.Listen("tcp", (fmt.Sprintf(":%s", getEnv("PORT", "8000"))))
	if err != nil {
		log.L.Fatal().Err(err).Msg("")
	}

	grpcS := grpc.NewServer()
	proto.RegisterBibleServerServer(grpcS, &GRPCBibleServer{})
	reflection.Register(grpcS)

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.Any())

	router := gin.Default()
	router.Use(logger.SetLogger(logger.Config{
		Logger: log.L,
	}))
	router.GET("v1/:bible/books", listBooks)
	router.GET("v1/:bible/:book/:chapter/:verse_from", getVersesHandler)
	router.GET("v1/:bible/:book/:chapter/:verse_from/:verse_to", getVersesHandler)

	go grpcS.Serve(grpcL)
	go router.RunListener(httpL)
	m.Serve()
}
