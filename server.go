package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Chandra5468/grpc-practise-2/coffeeshop_protos"
	"google.golang.org/grpc"
)

// This below UnimplementedCoffeeShopServer is provided by our generated code (grpc.pb.go) and embedd it into our server.
// This struct has methods of receiver function associated with it.

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		{
			Id:   "1",
			Name: "Black Coffee",
		},
		{
			Id:   "2",
			Name: "Pulao",
		},
		{
			Id:   "3",
			Name: "Chole Bature",
		},
	}

	for i, _ := range items {
		srv.Send(&pb.Menu{ // Here we are doing streaming as mentioned in proto file
			Items: items[0 : i+1],
		})
	}

	return nil
}
func (s *server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}
func (s *server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "In Progress",
	}, nil
}
func main() {
	list, err := net.Listen("tcp", "localhost:9001") // tcp listerner on 9001. How do we mention IP if this app is running in different servers as container

	if err != nil {
		log.Fatalf("Failed to listen on port %s", err.Error())
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Unable to start grpc server %v", err)
	}
}
