package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
	"github.com/MirableOne/word-of-wisdom/pkg/server"
	"net"
)

type Handler struct {
	log *logger.Logger
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

func main() {
	log := logger.NewLogger()

	serverConfig := &server.Config{
		Host:   "0.0.0.0",
		Port:   "3000",
		Logger: log,
	}

	h := &Handler{
		log: log,
	}

	srv := server.NewServer(serverConfig)

	log.Error(srv.Listen(h).Error())
}
