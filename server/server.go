package icey_com

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct {
}

/*
	Add(context.Context, *Req) (*Res, error)
	Muil(context.Context, *Req) (*Res, error)
*/

func (this *server) Add(ctx context.Context, in *Req) (*Res, error) {
	sum := in.A + in.B
	return &Res{Sum: sum}, nil
}

func (this *server) Muil(ctx context.Context, in *Req) (*Res, error) {
	muil := in.A * in.B
	return &Res{Muil: muil}, nil
}

func Server() {
	//绑定端口
	lis, err := net.Listen("tcp", ":7890")
	if err != nil {
		log.Fatal("listen error: ", err)
		return
	}
	//创建grpc服务
	s := grpc.NewServer()
	//绑定服务接口
	RegisterMathServer(s, &server{})
	//将端口服务由grpc服务代理
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Serve error: ", err)
	}

}
