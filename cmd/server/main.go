package main

import (
	"github.com/MirableOne/word-of-wisdom/pkg/config"
	"github.com/MirableOne/word-of-wisdom/pkg/factory"
	"github.com/MirableOne/word-of-wisdom/pkg/server"
)

func main() {
	log := factory.MakeLogger()
	handler := factory.MakeHandler()

	serverConfig := &server.Config{
		Logger: log,
	}

	config.Must(config.Configure(serverConfig))

	srv := server.NewServer(serverConfig)

	log.Error(srv.Listen(handler).Error())
}
