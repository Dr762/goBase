package server

import (
	"errors"
	"net/rpc"
	"net/http"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func HttpRpcServer() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe("localhost:2100", nil)
	if err != nil {
		log.Println(err)
	}
}


func TcpRpcServer(){
	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", "localhost:2200")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		rpc.ServeConn(conn)
	}
}


func JsonRpcServer(){
	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", "localhost:2300")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		jsonrpc.ServeConn(conn)
	}
}