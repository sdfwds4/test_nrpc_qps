package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/nats-io/nats.go"

	// This is the package containing the generated *.pb.go and *.nrpc.go
	// files.
	pb "github.com/sdfwds4/test_nrpc_qps/proto"
)

var counter int64

// server implements the pb.GreeterServer interface.
type server struct{}

// SayHello is an implementation of the SayHello method from the definition of
// the Greeter service.
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloReply, err error) {

	atomic.AddInt64(&counter, 1)

	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
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

	// Our server implementation.
	s := &server{}

	// The NATS handler from the pb.nrpc.proto file.
	h := pb.NewGreeterHandler(context.TODO(), nc, s)

	// Start a NATS subscription using the handler. You can also use the
	// QueueSubscribe() method for a load-balanced set of servers.
	sub, err := nc.Subscribe(h.Subject(), h.Handler)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	go func() {
		var t = time.Now().UnixNano() / 1e6
		for {
			select {
			case <-time.After(time.Second * 5):
				now := time.Now().UnixNano() / 1e6
				v := atomic.SwapInt64(&counter, 0)
				log.Print("count: ", float64(v)/float64((now-t)/1000), "/s")
				t = now
			}
		}
	}()

	// Keep running until ^C.
	fmt.Println("server is running, ^C quits.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(c)
}
