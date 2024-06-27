package main

import (
	"mochain/network"
)

func main(){

	trlocal := network.NewLocalTransport("LOCAL")
	opts := network.ServerOpts{
		Transports: []network.Transport{trlocal},
	}

	server := network.NewServer(opts)

	print(server)

	server.Start()
}