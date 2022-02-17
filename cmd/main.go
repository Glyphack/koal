package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/migrate"
	authv1 "github.com/glyphack/koal/gen/proto/go/auth/v1"
	"github.com/glyphack/koal/internal/config"
	authapi "github.com/glyphack/koal/internal/module/auth/api"
	"github.com/glyphack/koal/pkg/corsutils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	config.InitConfig()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	client := newClient()
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	defer client.Close()
	authv1.RegisterAuthServiceServer(s, authapi.NewServer(client))

	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	r := http.NewServeMux()

	fs := http.FileServer(http.Dir("./api-docs/"))
	r.Handle("/api-docs/", http.StripPrefix("/api-docs/", fs))

	gwmux := runtime.NewServeMux()
	r.Handle("/", corsutils.Cors(gwmux, corsutils.AllowOrigin))
	err = authv1.RegisterAuthServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: corsutils.Cors(r, corsutils.AllowOrigin),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

func newClient() *ent.Client {
	if viper.GetBool("debug") == true {
		client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatal(err)
		}
		return client
	} else {
		client, err := ent.Open("postgres", viper.GetString("postgres_uri"))
		if err != nil {
			log.Fatal(err)
		}
		return client
	}

}
