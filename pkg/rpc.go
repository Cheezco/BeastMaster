package pkg

import (
	"BeastMaster/internal"
	"BeastMaster/internal/debug"
	"errors"
	"net"
	"net/rpc"
	"time"
)

func RemoteServerExists(network, address string, retryCount int) bool {
	retries := 0
	timeout := 2 * time.Second
	dialer := &net.Dialer{
		Timeout: timeout,
	}
	success := false
	for retries < retryCount && !success {
		client, err := dialer.Dial(network, address)
		if err != nil {
			retries++
			continue
		}
		err = client.Close()
		if err != nil {
			continue
		}
		success = true
	}

	return success
}

func StartRpcServer(network string, address string) {
	listener, err := net.Listen(network, address)
	internal.CheckError(err)
	for {
		conn, err := listener.Accept()
		internal.CheckError(err)
		go rpc.ServeConn(conn)
	}
}

func dialAndCall(network, address, method string, args any, reply any) error {
	if !RemoteServerExists(network, address, 5) {
		return errors.New("remote server doesn't exist")
	}

	client, err := rpc.Dial(network, address)
	if err != nil {
		return err
	}
	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {
			debug.Log(err)
			return
		}
	}(client)

	err = client.Call(method, args, reply)
	return err
}

func DialNoArgs(network string, address string, method string, reply any) error {
	return dialAndCall(network, address, method, "", reply)
}

func DialNoReply(network string, address string, method string, args any) error {
	var reply interface{}
	return dialAndCall(network, address, method, args, &reply)
}
