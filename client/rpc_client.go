package client

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func HttpRpcClient(server string){
	clt,err:= rpc.DialHTTP("tcp",server)
	if err != nil {
		log.Fatal(err)
	}

	args := Args{17,8}
	var reply int

	err=clt.Call("Arith.Multiply",args,&reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient

	err = clt.Call("Arith.Divide",args,&quot)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}


func TcpRpcClient(server string){
	clt,err := rpc.Dial("tcp",server)
	if err != nil {
		log.Fatal(err)
	}

	args := Args{23,9}
	var reply int

	err=clt.Call("Arith.Multiply",args,&reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient

	err = clt.Call("Arith.Divide",args,&quot)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}


func JsonRpcClient(server string){
	clt,err := jsonrpc.Dial("tcp",server)
	if err != nil {
		log.Fatal(err)
	}

	args := Args{45,6}
	var reply int

	err=clt.Call("Arith.Multiply",args,&reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient

	err = clt.Call("Arith.Divide",args,&quot)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}