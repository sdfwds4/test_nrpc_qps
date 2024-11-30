package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"

	// This is the package containing the generated *.pb.go and *.nrpc.go
	// files.
	pb "github.com/sdfwds4/test_nrpc_qps/proto"
)

func main() {

	fmt.Println("client start ...")

	start := time.Now()

	var natsURL = nats.DefaultURL
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	// Connect to the NATS server.
	nc, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// This is our generated client.
	cli := pb.NewGreeterClient(nc)

	var resp *pb.HelloReply
	for {
		// Contact the server and print out its response.
		resp, err = cli.SayHello(&pb.HelloRequest{Name: "world"})
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("duration:", time.Since(start))

	// print
	fmt.Printf("Greeting: %s\n", resp.GetMessage())
}
