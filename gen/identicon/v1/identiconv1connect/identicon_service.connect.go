// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: identicon/v1/identicon_service.proto

package identiconv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/kushidam/identicon-service/gen/identicon/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// IdenticonServiceName is the fully-qualified name of the IdenticonService service.
	IdenticonServiceName = "identicon.v1.IdenticonService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IdenticonServiceGenerateIdenticonProcedure is the fully-qualified name of the IdenticonService's
	// GenerateIdenticon RPC.
	IdenticonServiceGenerateIdenticonProcedure = "/identicon.v1.IdenticonService/GenerateIdenticon"
)

// IdenticonServiceClient is a client for the identicon.v1.IdenticonService service.
type IdenticonServiceClient interface {
	// テキストデータを受け取り、Identiconを生成し、格納パスを返すメソッド
	GenerateIdenticon(context.Context, *connect.Request[v1.GenerateIdenticonRequest]) (*connect.Response[v1.GenerateIdenticonResponse], error)
}

// NewIdenticonServiceClient constructs a client for the identicon.v1.IdenticonService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIdenticonServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IdenticonServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &identiconServiceClient{
		generateIdenticon: connect.NewClient[v1.GenerateIdenticonRequest, v1.GenerateIdenticonResponse](
			httpClient,
			baseURL+IdenticonServiceGenerateIdenticonProcedure,
			opts...,
		),
	}
}

// identiconServiceClient implements IdenticonServiceClient.
type identiconServiceClient struct {
	generateIdenticon *connect.Client[v1.GenerateIdenticonRequest, v1.GenerateIdenticonResponse]
}

// GenerateIdenticon calls identicon.v1.IdenticonService.GenerateIdenticon.
func (c *identiconServiceClient) GenerateIdenticon(ctx context.Context, req *connect.Request[v1.GenerateIdenticonRequest]) (*connect.Response[v1.GenerateIdenticonResponse], error) {
	return c.generateIdenticon.CallUnary(ctx, req)
}

// IdenticonServiceHandler is an implementation of the identicon.v1.IdenticonService service.
type IdenticonServiceHandler interface {
	// テキストデータを受け取り、Identiconを生成し、格納パスを返すメソッド
	GenerateIdenticon(context.Context, *connect.Request[v1.GenerateIdenticonRequest]) (*connect.Response[v1.GenerateIdenticonResponse], error)
}

// NewIdenticonServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIdenticonServiceHandler(svc IdenticonServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	identiconServiceGenerateIdenticonHandler := connect.NewUnaryHandler(
		IdenticonServiceGenerateIdenticonProcedure,
		svc.GenerateIdenticon,
		opts...,
	)
	return "/identicon.v1.IdenticonService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IdenticonServiceGenerateIdenticonProcedure:
			identiconServiceGenerateIdenticonHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIdenticonServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIdenticonServiceHandler struct{}

func (UnimplementedIdenticonServiceHandler) GenerateIdenticon(context.Context, *connect.Request[v1.GenerateIdenticonRequest]) (*connect.Response[v1.GenerateIdenticonResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("identicon.v1.IdenticonService.GenerateIdenticon is not implemented"))
}
