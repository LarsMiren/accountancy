package main

import (
	"context"
	"flag"
	"net"

	pb "github.com/LarsMiren/accountancy/proto/auth"
	"github.com/LarsMiren/accountancy/proto/general"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
)

type authServer struct{}

var (
	db   *gorm.DB
	host string
	port string
)

func init() {
	db, err := gorm.Open("mysql", "$user:$password@$host/$dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

}

func newAuthServer() pb.AuthServer {
	return new(authServer)
}

func (s *authServer) Login(
	c context.Context,
	u *general.User) (*general.User, error) {
	return u, nil
}
func (s *authServer) Logout(
	c context.Context,
	e *empty.Empty) (*general.Confirmation, error) {
	return new(general.Confirmation), nil
}
func (s *authServer) Signup(
	c context.Context,
	u *general.User) (*general.Confirmation, error) {
	return new(general.Confirmation), nil
}

func run() error {
	listen, err := net.Listen("tcp", "5051")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterAuthServer(server, newAuthServer())
	return server.Serve(listen)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
