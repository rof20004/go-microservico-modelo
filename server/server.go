package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/rof20004/mensagem_microservico"
	"golang.org/x/net/context"
)

var (
	address  = "localhost:10000"
	endpoint = flag.String("mensagem_endpoint", address, "Serviço Mensagem")
)

// Server struct
type Server struct{}

// Echo handler
func (s *Server) Echo(ctx context.Context, in *pb.MensagemRequest) (*pb.MensagemResponse, error) {
	fmt.Printf("rpc request Echo(%q)\n", in.Entrada)
	return &pb.MensagemResponse{Saida: "SAÍDA OK"}, nil
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// func serveSwagger(mux *http.ServeMux) {
// 	mime.AddExtensionType(".svg", "image/svg+xml")

// 	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
// 	fileServer := http.FileServer(&assetfs.AssetFS{
// 		Asset:    swagger.Asset,
// 		AssetDir: swagger.AssetDir,
// 		Prefix:   "third_party/swagger-ui",
// 	})
// 	prefix := "/swagger-ui/"
// 	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
// }

// Run -> start rpc server
func Run() error {
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, address))}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMensagemServiceServer(grpcServer, &Server{})
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs:    demoCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()
	// mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	// 	io.Copy(w, strings.NewReader(pb.Swagger))
	// })

	gwmux := runtime.NewServeMux()
	err := pb.RegisterMensagemServiceHandlerFromEndpoint(ctx, gwmux, *endpoint, dopts)
	if err != nil {
		return err
	}

	mux.Handle("/", gwmux)
	// serveSwagger(mux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    address,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("Serviço gRPC/REST MENSAGEM iniciado às %q na porta %d\n", time.Now(), port)
	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))

	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		glog.Fatal(err.Error())
	}
}
