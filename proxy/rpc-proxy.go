package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	gwAuth "github.com/LarsMiren/accountancy/proto/auth"
	gwLogic "github.com/LarsMiren/accountancy/proto/logic"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	bindAddr      = flag.String("bind_address", "0.0.0.0:11000", "Bind endpoint")
	endpointAuth  = flag.String("auth_endpoint", "127.0.0.1:5051", "Auth endpoint")
	endpointLogic = flag.String("logic_endpoint", "127.0.0.1:5052", "Logic enpoint")
)

var memCacheClient *memcache.Client

func init() {
	memCacheClient = memcache.New("127.0.0.1:11211")
}

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

	err := gwAuth.RegisterAuthHandlerFromEndpoint(ctx, mux, *endpointAuth, opts)
	if err != nil {
		return err
	}
	err = gwLogic.RegisterLogicHandlerFromEndpoint(ctx, mux, *endpointLogic, opts)
	if err != nil {
		return err
	}
	fmt.Printf("Starting http server")
	err = http.ListenAndServe(*bindAddr, allowCORS(mux))
	return err
}

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	if err := run(ctx); err != nil {
		glog.Fatal(err)
	}
}
