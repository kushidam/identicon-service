package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	identiconv1 "github.com/kushidam/identicon-service/gen/identicon/v1"
	"github.com/kushidam/identicon-service/gen/identicon/v1/identiconv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GenerateIdenticonServer struct {}

func (s *GenerateIdenticonServer) GenerateIdenticon(
    ctx context.Context,
    req *connect.Request[identiconv1.GenerateIdenticonRequest],
) (*connect.Response[identiconv1.GenerateIdenticonResponse], error) {


    res := connect.NewResponse(&identiconv1.GenerateIdenticonResponse{
        ImageData: []byte{0x0a, 0x24, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f},
    })

    return res, nil
}

func main() {
	generateServer := &GenerateIdenticonServer{}
	mux := http.NewServeMux()
	path, handler := identiconv1connect.NewIdenticonServiceHandler(generateServer)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:50051",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}