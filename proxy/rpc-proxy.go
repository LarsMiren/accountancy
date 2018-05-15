package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/LarsMiren/accountancy/env"
	gwAuth "github.com/LarsMiren/accountancy/proto/auth"
	gwLogic "github.com/LarsMiren/accountancy/proto/logic"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type session struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	bindAddr      = flag.String("bind_address", "0.0.0.0"+env.GetPort("proxy"), "Bind endpoint")
	endpointAuth  = flag.String("auth_endpoint", env.GetFullAddr("auth"), "Auth endpoint")
	endpointLogic = flag.String("logic_endpoint", env.GetFullAddr("logic"), "Logic enpoint")
)

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Millisecond * 3000),
		grpc.FailOnNonTempDialError(true),
		grpc.WithBlock(),
	}
	fmt.Println("Connecting to auth service")
	err := gwAuth.RegisterAuthHandlerFromEndpoint(ctx, mux, *endpointAuth, opts)
	if err != nil {
		glog.Fatal(err)
		return err
	}
	fmt.Println("Connecting to logic server")
	err = gwLogic.RegisterLogicHandlerFromEndpoint(ctx, mux, *endpointLogic, opts)
	if err != nil {
		glog.Fatal(err)
		return err
	}
	fmt.Println("Allowing CORS")
	return http.ListenAndServe(*bindAddr, allowCORS(mux))
}

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	if err := run(ctx); err != nil {
		glog.Fatal(err)
	}
}
