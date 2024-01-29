package handler

import (
	"bufio"
	"context"
	"fmt"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
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
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			err := conn.Close()
			h.log.Error(err.Error())
			return
		}
		h.log.Info(fmt.Sprintf("get message: %s", message))

		_, err = conn.Write([]byte("hello there.\n"))
		if err != nil {
			h.log.Error(err.Error())
		}
	}
}
