package handler

import (
	"context"
	"fmt"
	"github.com/MirableOne/word-of-wisdom/pkg/hashcash"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
	"github.com/MirableOne/word-of-wisdom/pkg/protocol"
	"github.com/MirableOne/word-of-wisdom/pkg/quotes"
	"net"
)

type Config struct {
	Log     *logger.Logger
	Storage hashcash.Storage
}

type Handler struct {
	log     *logger.Logger
	storage hashcash.Storage
}

func NewHandler(cfg *Config) *Handler {
	return &Handler{
		log:     cfg.Log,
		storage: cfg.Storage,
	}
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	for {
		message, err := protocol.Read(ctx, conn)

		if err != nil {
			h.disconnect(conn)
			return
		}
		h.log.Info(fmt.Sprintf("incoming message[%s]: %s", message.Type, message.Body))

		if !hashcash.VerifyWithStorage(message.Body, h.storage) {
			err = protocol.Send(ctx, conn, &protocol.Message{
				Type: protocol.Error,
				Body: "invalid header; disconnecting",
			})
			h.disconnect(conn)
			return
		}

		err = protocol.Send(ctx, conn, &protocol.Message{
			Type: protocol.QuoteResponse,
			Body: quotes.RandomQuote(),
		})

		if err != nil {
			h.log.Error(err.Error())
		}
	}
}

func (h *Handler) disconnect(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		h.log.Error(err.Error())
	}
}
