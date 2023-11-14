# identicon-service
"identicon-service" is a gRPC-based microservice written in Go that generates unique Identicons for use in user profiles and avatars.
It uses the gRPC protocol to provide Identicons based on user-provided strings. 

`identicon-service` はGoで書かれたgRPCベースのサービスで、ユーザプロファイルやアバターで使用できる一意のIdenticonsを生成します。

gRPCプロトコルを使用し、ユーザーから提供された文字列に基づいてIdenticonを提供します。

サーバは、Identiconのバイナリデータを返します。
クライアント側でバイナリデータを画像に変換（デコード）を行ってください。

# Identicon
与えられた文字列に応じて、アイコンを生成する。

![icon](assets/identicon_default.png)
![icon](assets/identicon_argtext.png)

## 生成用設定ファイルコピー
コマンドを実行し、`config.ini`を作成する。
ファイル内に任意の値に設定する。
```sh
cp ./config.ini.example ./config.ini
```
```sh
# config.ini
# サービスの名前（任意）
namespace = "サービス名など任意の名前"
# 画像のサイズ（n*n）
blocks_size = n
# 画像の密度
density = 5
```

##　 サーバの起動
```sh
go run ./cmd/server/main.go 
```

## クライアントCLIから動作確認
同じ文字列からは一意のIdenticonsが生成されます。
`-text`オプションで任意の文字列を引数に与えれます。
この場合、生成されるIdenticonsが違うことが確認できます。

* デフォルト
```sh
go run ./cmd/client/main.go 
```
* 任意の文字列
```sh
go run ./cmd/client/main.go -text Sample
```
