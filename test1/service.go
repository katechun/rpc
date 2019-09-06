package test1

import (
	"context"
	"github.com/katechun/rpc/test1/tools"
	"github.com/smallnest/rpcx/server"
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *tools.Args, reply *tools.Reply) error {
	reply.C = args.A + args.B
	return nil
}

func Register() {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", ":8972")
}

func main() {
	Register()
}
