package server

import (
	"context"
	"fmt"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
	"net"
)

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
}

type Config struct {
	Host string `env:"HOST" validate:"required"`
	Port string `env:"PORT" validate:"required"`

	Logger *logger.Logger
}

type Server struct {
	host string
	port string

	log *logger.Logger

	listener net.Listener
}

func NewServer(cfg *Config) *Server {
	return &Server{
		host: cfg.Host,
		port: cfg.Port,

		log: cfg.Logger,
	}
}

func (s *Server) Listen(handler Handler) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))

	if err != nil {
		s.log.Fatal(err.Error())
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			s.log.Fatal(err.Error())
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go handler.Handle(context.Background(), conn)
	}
}
