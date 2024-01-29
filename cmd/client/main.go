package main

import (
	"context"
	"github.com/MirableOne/word-of-wisdom/pkg/factory"
	"github.com/MirableOne/word-of-wisdom/pkg/protocol"
	"net"
	"time"
)

func main() {
	log := factory.MakeLogger()
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3000")

	if err != nil {
		log.Fatal(err.Error())
	}

	ticker := time.NewTicker(5 * time.Second)

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()
	for {
		<-ticker.C

		err = protocol.Send(ctx, conn, &protocol.Message{
			Type: protocol.QuoteRequest,
			Body: "hello there;",
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

	conn.Close()
}
