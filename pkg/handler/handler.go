package handler

import (
	"context"
	"fmt"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
	"github.com/MirableOne/word-of-wisdom/pkg/protocol"
	"net"
)

type Config struct {
	Log *logger.Logger
}

type Handler struct {
	log *logger.Logger
}

func NewHandler(cfg *Config) *Handler {
	return &Handler{
		log: cfg.Log,
	}
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	for {
		message, err := protocol.Read(ctx, conn)

		if err != nil {
			err := conn.Close()
			h.log.Error(err.Error())
			return
		}
		h.log.Info(fmt.Sprintf("get message: %s", message.Body))

		err = protocol.Send(ctx, conn, &protocol.Message{
			Type: protocol.QuoteResponse,
			Body: "hello there;",
		})

		if err != nil {
			h.log.Error(err.Error())
		}
	}
}
