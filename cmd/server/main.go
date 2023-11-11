package main

import (
	"bytes"
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/kushidam/identicon-service/config"
	identiconv1 "github.com/kushidam/identicon-service/gen/identicon/v1"
	"github.com/kushidam/identicon-service/gen/identicon/v1/identiconv1connect"
	"github.com/nullrocks/identicon"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GenerateIdenticonServer struct{}

func (s *GenerateIdenticonServer) GenerateIdenticon(
	ctx context.Context,
	req *connect.Request[identiconv1.GenerateIdenticonRequest],
) (*connect.Response[identiconv1.GenerateIdenticonResponse], error) {
	// Create the identicon generator
	ig, err := identicon.New(config.Config.NAMESPACE, config.Config.BLOCKS_SIZE, config.Config.DENSITY)

	if err != nil {
		return nil, err
	}

	ii, err := ig.Draw(req.Msg.Text)
	if err != nil {
		return nil, err
	}

	var imgBuffer bytes.Buffer
	if err := ii.Png(300, &imgBuffer); err != nil {
		return nil, err
	}

	imgData := imgBuffer.Bytes()

	return connect.NewResponse(&identiconv1.GenerateIdenticonResponse{
		ImageData: imgData,
	}), nil
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
