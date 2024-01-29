package main

import (
	"context"
	"github.com/MirableOne/word-of-wisdom/pkg/factory"
	"github.com/MirableOne/word-of-wisdom/pkg/hashcash"
	"github.com/MirableOne/word-of-wisdom/pkg/protocol"
	"net"
)

func main() {
	log := factory.MakeLogger()
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3000")

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
