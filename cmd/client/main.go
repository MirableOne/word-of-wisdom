package main

import (
	"context"
	"fmt"
	"github.com/MirableOne/word-of-wisdom/pkg/config"
	"github.com/MirableOne/word-of-wisdom/pkg/factory"
	"github.com/MirableOne/word-of-wisdom/pkg/hashcash"
	"github.com/MirableOne/word-of-wisdom/pkg/protocol"
	"net"
)

type Config struct {
	Host string `env:"HOST" validate:"required"`
	Port string `env:"PORT" validate:"required"`
}

func main() {
	log := factory.MakeLogger()

	cfg := &Config{}

	config.Must(config.Configure(cfg))

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))

	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	for {
		hash, err := hashcash.Make().Mint("test@test.com")

		err = protocol.Send(ctx, conn, &protocol.Message{
			Type: protocol.QuoteRequest,
			Body: hash,
		})

		if err != nil {
			log.Fatal(err.Error())
		}

		message, err := protocol.Read(ctx, conn)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Info(message.Body)
	}
}
