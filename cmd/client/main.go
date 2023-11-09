package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	identiconv1 "github.com/kushidam/identicon-service/gen/identicon/v1"
	identiconv1connect "github.com/kushidam/identicon-service/gen/identicon/v1/identiconv1connect"
)


var (
	userText    = flag.String("text", "nullrocks", "use decides the resulting figure")
)

func main(){
	flag.Parse()
	fmt.Println(userText)

	// gRPCクライアントを作成
	client := identiconv1connect.NewIdenticonServiceClient(
		http.DefaultClient,
		"http://localhost:50051",
	)


	// GenerateIdenticonメソッドを呼び出す
	res, err := client.GenerateIdenticon(
        context.Background(),
        connect.NewRequest(&identiconv1.GenerateIdenticonRequest{Text: *userText}),
    )

	if err != nil {
		log.Fatalf("GenerateIdenticonの呼び出しエラー: %v", err)
	}
	log.Printf("生成されたアイデンティコンのバイトデータ: %v", res.Msg.ImageData)
	
	// 生成されたアイデンティコンのバイトデータをファイルに書き込む
	file, err := os.Create("identicon.png")
	if err != nil {
		log.Fatalf("ファイルの作成エラー: %v", err)
	}
	defer file.Close()

	_, err = file.Write(res.Msg.ImageData)
	if err != nil {
		log.Fatalf("ファイルの書き込みエラー: %v", err)
	}

}