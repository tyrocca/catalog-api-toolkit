package main

import (
	"context"
	"log"
	"net"
	"time"

	// importing generated stubs
	"google.golang.org/protobuf/types/known/timestamppb"

	gen "github.com/tyrocca/catalog-api-toolkit/gen/go/catalog/v1"
	"google.golang.org/grpc"
)

// GreeterServerImpl will implement the service defined in protocol buffer definitions
type CatalogServicerImpl struct {
	gen.UnimplementedCatalogServiceServer
}

// SayHello is the implementation of RPC call defined in protocol definitions.
// This will take HelloRequest message and return HelloReply
func (s *CatalogServicerImpl) CreateCompany(ctx context.Context, request *gen.CreateCompanyRequest) (*gen.Company, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &gen.Company{
		Name:        request.Company.Name,
		Uid:         "someid",
		DisplayName: request.Company.DisplayName,
		CreateTime:  timestamppb.New(time.Now().UTC()),
		UpdateTime:  timestamppb.New(time.Now().UTC()),
	}, nil
}

func (s *CatalogServicerImpl) GetCompany(ctx context.Context, request *gen.GetCompanyRequest) (*gen.Company, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	log.Println(request)
	log.Println("oh shit")
	return &gen.Company{
		Name:        request.Name,
		Uid:         "someid",
		DisplayName: "Oh this is a name",
		CreateTime:  timestamppb.Now(),
		UpdateTime:  timestamppb.Now(),
	}, nil
}

func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gen.RegisterCatalogServiceServer(server, &CatalogServicerImpl{})
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatal("error in listening on port :8080", err)
	} else {
		// the start gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
