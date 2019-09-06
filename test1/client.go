package test1

import (
	"context"
	"github.com/katechun/rpc/test1/tools"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	Addr := tools.RpcAddr{"localhost"}
	d := client.NewPeer2PeerDiscovery("tpc@"+Addr.Host, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &tools.Args{
		A: 10,
		B: 20,
	}

	reply := &tools.Reply{}

	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
