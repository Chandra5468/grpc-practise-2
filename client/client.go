package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Chandra5468/grpc-practise-2/coffeeshop_protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials())) // always try to implement in secure way. Eventhough you use service to service comminication

	if err != nil {
		log.Fatalf("Failed to connect to grpc server %v", err)
	}

	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1) // it would not take much time to get resources

	defer cancel()

	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})

	if err != nil {
		cancel()
		log.Fatalf("error calling function getMenu %s", err.Error())
	}

	done := make(chan bool)

	var items []*pb.Item

	go func() {
		// for {

		// resp, err := menuStream.Recv()
		// if err == io.EOF {
		// 	done <- true
		// 	return
		// }
		// if err != nil {
		// 	cancel()
		// 	log.Fatalf("cannot receive %v", err) // handle this error robustly without Fatalf as this is like api request, response
		// }

		// items = resp.Items
		// log.Println("Resp received ", items)
		// }
		items = append(items, menuStream.Items...)
		// items = append(items, menuStream.RecvMsg())
		log.Println("Resp items received are ", items)
		done <- true
	}()

	<-done

	recepipt, _ := c.PlaceOrder(ctx, &pb.Order{Items: items}) // we are ordering all items so, sending items

	log.Println("receipt is ", recepipt)

	status, _ := c.GetOrderStatus(ctx, recepipt)

	log.Printf("Order status is %v", status)
}
