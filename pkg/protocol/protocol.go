package protocol

import (
	"bufio"
	"context"
	"encoding/json"
	"net"
)

const delim = '\n'

func Send(ctx context.Context, conn net.Conn, message *Message) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = conn.Write(msg)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte{delim})
	if err != nil {
		return err
	}

	return nil
}

func Read(ctx context.Context, conn net.Conn) (*Message, error) {
	reader := bufio.NewReader(conn)
	msg, err := reader.ReadString(delim)
	if err != nil {
		return nil, err
	}

	var message Message

	err = json.Unmarshal([]byte(msg), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
