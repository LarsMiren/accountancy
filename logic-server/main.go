package main

import (
	"context"
	"flag"
	"net"

	"github.com/LarsMiren/accountancy/env"
	"github.com/LarsMiren/accountancy/proto/general"
	pb "github.com/LarsMiren/accountancy/proto/logic"
	"github.com/golang/glog"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
)

type logicServer struct{}

func newLogicServer() *logicServer {
	return new(logicServer)
}

func (l *logicServer) GetUser(ctx context.Context, id *pb.Id) (*general.User, error) {
	user := new(general.User)
	user.Id = id.Id
	user.Name = "John Doe"
	user.Email = "johndoe@gmail.com"
	user.Username = "user"
	return user, nil
}
func (l *logicServer) GetProductById(ctx context.Context, id *pb.Id) (*pb.Product, error) {
	product := new(pb.Product)
	product.Id = "1"
	product.Name = "Product1"
	product.Image = []byte{1}
	product.SupplierId = "1"
	product.Type = "type1"
	product.Description = "first product"
	return product, nil
}
func (l *logicServer) GetAllUsers(ctx context.Context, e *google_protobuf1.Empty) (*pb.Users, error) {
	return nil, nil
}
func (l *logicServer) GetProductsByType(ctx context.Context, t *pb.ProductType) (*pb.Products, error) {
	return nil, nil
}
func (l *logicServer) GetProductsByUser(ctx context.Context, id *pb.Id) (*pb.Products, error) {
	return nil, nil
}
func (l *logicServer) UpdateUser(ctx context.Context, u *general.User) (*general.Confirmation, error) {
	return nil, nil
}
func (l *logicServer) UpdateProduct(ctx context.Context, p *pb.Product) (*general.Confirmation, error) {
	return nil, nil
}
func (l *logicServer) CreateProduct(ctx context.Context, p *pb.Product) (*general.Confirmation, error) {
	return nil, nil
}
func (l *logicServer) DeleteUser(ctx context.Context, id *pb.Id) (*general.Confirmation, error) {
	return nil, nil
}
func (l *logicServer) DeleteProduct(ctx context.Context, id *pb.Id) (*general.Confirmation, error) {
	return nil, nil
}
func (l *logicServer) Subscribe(ctx context.Context, id *pb.Id) (*general.Confirmation, error) {
	return nil, nil
}
func run(port string) error {

	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	check(err)
	server := grpc.NewServer()
	pb.RegisterLogicServer(server, newLogicServer())
	return server.Serve(listen)
}

func check(err error) {
	if err != nil {
		glog.Fatal(err)
		panic(err)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()
	db = connectToDB()
	defer db.DB().Close()
	port := env.GetPort("logic")
	err := run(port)
	check(err)
}
