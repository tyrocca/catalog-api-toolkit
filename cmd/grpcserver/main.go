package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	// importing generated stubs
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/oklog/ulid/v2"
	gen "github.com/tyrocca/catalog-api-toolkit/gen/go/catalog/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Ulid generator
func genUlid() ulid.ULID {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

// GreeterServerImpl will implement the service defined in protocol buffer definitions
type CatalogServicerImpl struct {
	gen.UnimplementedCatalogServiceServer
}

// buildName creates the name of a company
func (s *CatalogServicerImpl) buildCompanyName(uid string) string {
	return fmt.Sprintf("companies/%s", uid)
}

// parseName parses the name into its components
func (s *CatalogServicerImpl) parseName(name string) []string {
	return strings.Split(name, "/")
}

// uidFromName will return the uid from a Name
func (s *CatalogServicerImpl) uidFromName(name string) string {
	splitName := s.parseName(name)
	return splitName[len(splitName)-1]
}

// SayHello is the implementation of RPC call defined in protocol definitions.
// This will take HelloRequest message and return HelloReply
func (s *CatalogServicerImpl) CreateCompany(ctx context.Context, request *gen.CreateCompanyRequest) (*gen.Company, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println("Metadata: ", md)
	}

	uid := genUlid()

	return &gen.Company{
		Name:        s.buildCompanyName(uid.String()),
		Uid:         uid.String(),
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
		Uid:         s.uidFromName(request.Name),
		DisplayName: "Oh this is a name",
		CreateTime:  timestamppb.Now(),
		UpdateTime:  timestamppb.Now(),
	}, nil
}

func (s *CatalogServicerImpl) UpdateCompany(ctx context.Context, request *gen.UpdateCompanyRequest) (*gen.Company, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	log.Println(request)
	log.Println("oh shit")
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println("Metadata: ", md)
	}
	return &gen.Company{
		Name:        request.Company.Name,
		Uid:         s.uidFromName(request.Company.Name),
		DisplayName: request.Company.DisplayName,
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
