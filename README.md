# identicon-service
"identicon-service" is a gRPC-based microservice written in Go that generates unique Identicons for use in user profiles and avatars.
It uses the gRPC protocol to provide Identicons based on user-provided strings. 
「identicon-service」 はGoで書かれたgRPCベースのサービスで、ユーザプロファイルやアバターで使用できる一意のIdenticonsを生成します。
gRPCプロトコルを使用し、ユーザーから提供された文字列に基づいてIdenticonを提供します。

## 生成用設定ファイルコピー
コマンドを実行し、`config.ini`を作成する。
ファイル内に任意の値に設定する。
```sh
cp ./config.ini.example ./config.ini

```

##　 サーバの起動
```sh
go run ./cmd/server/main.go 
```