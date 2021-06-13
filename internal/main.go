package internal

import (
	"context"
	gw "github.com/code-newbee/protocol/geeker"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const  (
	geekerEndPoint = "localhost:50001"
)

func Run()  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()
	err := gw.RegisterGeekerHandlerFromEndpoint(ctx, mux, geekerEndPoint, []grpc.DialOption{
		grpc.WithInsecure(),
	})

	if err != nil{
		log.Fatal("fail to register endpoint")
	}

	if err := http.ListenAndServe(":8081", mux); err != nil{
		log.Fatalf("fail to setup http svr")
	}
}
