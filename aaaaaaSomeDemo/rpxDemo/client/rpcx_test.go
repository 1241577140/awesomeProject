package main

import (
	"awesomeProject/aaaaaaSomeDemo/rpxDemo/server"
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"testing"
)

func BenchmarkRpcx(b *testing.B) {
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+"localhost:8972", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	defer xclient.Close()
	for i := 0; i < b.N; i++ {
		args := &server.Args{
			A: 10,
			B: 20,
		}

		reply := &Reply{}

		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

	}
}
