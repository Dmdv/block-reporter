package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"time"
	// To explicitly declare gRPC transport
	//"github.com/asim/go-micro/plugins/server/grpc/v4"
	// micro.Server(grpc.NewServer()),
)

var (
	service = "block-reporter"
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		fmt.Println("Ping")
		i++
	}
}

func read() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/22cb2849f5f74b8599f3dc2a23085bd4")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
			fmt.Println(block.GasUsed())           // 7
		}
	}
}

func main() {
	// Create service
	// micro.Server(grpc.NewServer()), // must come before any other options
	// micro.Version(version) - set version
	srv := micro.NewService(
		micro.Name(service),
	)
	srv.Init()

	// Register handler

	go pub()
	go read()

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
