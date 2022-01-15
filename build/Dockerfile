## 2021/12/27最新versionを取得
FROM golang:1.17.5-alpine
## アップデートとgitのインストール
RUN apk add --update && apk add git
## appディレクトリの作成
RUN mkdir /go/src/app
## ワーキングディレクトリの設定
WORKDIR /go/src/app

## 依存ライブラリをダウンロードする
ENV GO111MODULE=on
COPY cmd/go.mod /go/src/app/go.mod
COPY cmd/go.sum /go/src/app/go.sum
RUN go mod download

## ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

## go run を実行
CMD go run apitest/main.go
