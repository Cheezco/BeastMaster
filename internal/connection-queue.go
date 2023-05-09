package internal

import "github.com/gorilla/websocket"

type ConnectionQueue struct {
	Connection *websocket.Conn
	Queue      Queue
}

func (c *ConnectionQueue) Send() error {
	for !c.Queue.IsEmpty() {
		value := c.Queue.Pop()
		err := c.Connection.WriteJSON(value)
		if err != nil {
			return err
		}
	}

	return nil
}
