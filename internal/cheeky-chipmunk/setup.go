package cheeky_chipmunk

import (
	"BeastMaster/internal/debug"
	"BeastMaster/pkg/logging"
	"github.com/gorilla/websocket"
	"net/http"
	"net/rpc"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartWebSocketListener(address string, connections *[]*websocket.Conn) error {
	var returnErr error
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			returnErr = err
			return
		}
		index := len(*connections)
		*connections = append(*connections, conn)

		conn.SetCloseHandler(func(code int, text string) error {
			*connections = append((*connections)[:index], (*connections)[index+1:]...)
			debug.Log("Websocket connection lost")
			return nil
		})

		debug.Log("Websocket connection established")

		select {}
	})

	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return returnErr
}

func RegisterRpcServices() error {
	loggerPluginService := new(logging.PluginService)

	err := rpc.Register(loggerPluginService)
	if err != nil {
		return err
	}

	return nil
}
