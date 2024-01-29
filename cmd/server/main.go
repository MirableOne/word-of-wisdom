package main

import (
	"github.com/MirableOne/word-of-wisdom/pkg/factory"
	"github.com/MirableOne/word-of-wisdom/pkg/server"
)

func main() {
	log := factory.MakeLogger()
	handler := factory.MakeHandler()

	serverConfig := &server.Config{
		Host:   "0.0.0.0",
		Port:   "3000",
		Logger: log,
	}

	srv := server.NewServer(serverConfig)

	log.Error(srv.Listen(handler).Error())
}
