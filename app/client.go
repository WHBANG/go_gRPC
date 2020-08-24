package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	server "icey.com/server"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:7890", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Dial error:", err)
		return
	}
	defer conn.Close()
	c := server.NewMathClient(conn)
	defer func() {
		res, err := c.Add(context.Background(), &server.Req{A: 10, B: 20})
		if err != nil {
			log.Fatal("grpc error: ", err)
		}
		fmt.Println(res.Sum)
	}()

	defer func() {
		res, err := c.Muil(context.Background(), &server.Req{A: 10, B: 20})
		if err != nil {
			log.Fatal("grpc error: ", err)
		}
		fmt.Println(res.Muil)
	}()

}
