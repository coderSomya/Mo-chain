package main

import (
	"mochain/network"
	"time"
)

func main(){

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	go func(){
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello world"))
			time.Sleep(5* time.Second)
		}
	}()

	server := network.NewServer(opts)

	print(server)

	server.Start()
}