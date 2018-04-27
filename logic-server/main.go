package main

import (
	pb "github.com/LarsMiren/accountancy/proto/logic"

	_ "github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

type logicServer struct {
	pb.LogicServer
}

func main() {
	server := grpc.NewServer()
	pb.RegisterLogicServer(server, new(logicServer))
}
