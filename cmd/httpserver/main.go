package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	gen "github.com/tyrocca/catalog-api-toolkit/gen/go/catalog/v1"

	// "github.com/felixge/httpsnoop"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in allowed header, don't send the header
	return s, false
}

const GRPC_ENDPOINT = "localhost:8080"
const SERVER_PORT = ":8081"

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux(
		// convert header in response(going from gateway) from metadata received.
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}))
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := gen.RegisterCatalogServiceHandlerFromEndpoint(context.Background(), mux, GRPC_ENDPOINT, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/heartbeat"))
	r.HandleFunc("/api/*", func(w http.ResponseWriter, r *http.Request) {
		// Lifted from https://github.com/grpc-ecosystem/grpc-gateway/issues/769#issuecomment-478307237
		// gateway is generated to match for /v1/ and not /api/v1
		// we could update the gateway proto to match for /api/v1 but
		// it shouldn't care where it's mounted to, hence we just rewrite the path here
		r.URL.Path = strings.Replace(r.URL.Path, "/api", "", -1)
		log.Println("Here")
		log.Println(r.URL.Path)
		mux.ServeHTTP(w, r)
	})

	err = http.ListenAndServe(SERVER_PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}
