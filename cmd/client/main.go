package main

import (
	"bufio"
	"github.com/MirableOne/word-of-wisdom/pkg/factory"
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

	reader := bufio.NewReader(conn)
	for {
		<-ticker.C
		_, err = conn.Write([]byte("msg\n"))
		if err != nil {
			log.Fatal(err.Error())
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Info(message)
	}

	conn.Close()
}
