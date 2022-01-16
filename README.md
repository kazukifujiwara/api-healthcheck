# simple-golang-docker

apitest/apiendpoints.csv 内に書かれたAPIエンドポイント(URI)に対して、GETリクエストを送信し、レスポンスのヘッダからステータスコードのみを抽出する。

## Sample

### Input

apitest/apiendpoints.csv
```
id,name,description,uri
1,google,www.google.com,http://www.google.com/
2,yahoo,www.yahoo.co.jp,https://www.yahoo.co.jp/
3,bing,www.bing.com,https://www.bing.com/
```

### Output

```
# go run apitest/main.go
Status: 200 http://www.google.com/
Status: 200 https://www.yahoo.co.jp/
Status: 200 https://www.bing.com/
```

## Directories

```
.
├── README.md
├── build
│   └── Dockerfile
├── cmd
│   ├── apitest
│   │   ├── apiendpoints.csv
│   │   ├── example.csv
│   │   └── main.go
│   ├── example.csv
│   ├── go.mod
│   └── go.sum
└── docker-compose.yml
```

## How to use (初回実行手順)


### ①コンテナ外部で実行

```
# コンテナイメージのビルド
$ docker-compose build

# コンテナイメージの確認
$ docker images | grep simple_golang_docker

# コンテナ起動
$ docker-compose up -d

# コンテナステータス確認
$ docker-compose ps

# 作成したコンテナに直接アクセス
# (aplineベースのimageのcontainerの場合 /bin/bash ではなく /bin/ash を利用する)
$ docker exec -it simple_golang_docker /bin/ash 
```

### ②コンテナ内で実行

go.mod の初期化(不要かも)
```
go mod init
```

go.mod の更新を行い、 go.sum も生成する
```
go mod tidy
```

### ③コンテナの外部から実行
```
# 作成したファイルの実行
$ docker-compose exec simple_golang_docker go run hello/main.go
```

## API実行時jのパラメータ

* ヘッダ情報などを記載するならばcsvではなくてYAMLなどの方が良いかも

## csv読み込み時のファイルのパス

go run を実行するディレクトリが、ファイル実行時のcurrent directoryになる。
すなわち、
```
csvtest
┝ file.csv
┗ main.go
```
のような構成の場合、
go csvtest/main.go 実行時にcsvファイルを読み込む場合のパスは```csvtest/file.csv``` となり、
csvtestディレクトリ内で go run main.go 実行時にcsvファイルを読み込む場合```file.csv```となる。

## TODO

* リファクタリング
* Readme整理

## References

* [Go言語(golang)でcsvの読み書き Reader/Writer](https://golang.hateblo.jp/entry/golang-encoding-csv)
* [Goでhttpリクエストを送信する方法](https://qiita.com/taizo/items/c397dbfed7215969b0a5)
* [goでjson apiを叩く](https://qiita.com/nyamage/items/46f17318473aa141ffd5)
* [csvをGoの構造体にマッピングする](https://nakawatch.hatenablog.com/entry/2018/05/29/184737)
* [【golang】CSV の読み書きができるライブラリ「Go CSV」紹介](https://baba-s.hatenablog.com/entry/2018/11/23/220600)
* [Goのforとgoroutineでやりがちなミスとたった一つの冴えたgo vetと](https://qiita.com/sudix/items/67d4cad08fe88dcb9a6d)
* [サンプルで学ぶ Go 言語：JSON](https://www.spinute.org/go-by-example/json.html)
* [並行で発行したgoroutineをまとめ上げる処理](https://gist.github.com/paperlefthand/55c8147d5ebe7a243e9f65191be1d282)
* [Go+GinでCors設定を行い、クロスオリジンのアクセスを制御する](https://ti-tomo-knowledge.hatenablog.com/entry/2020/06/15/213401)
 * flaskからgoのAPIを呼び出そうとしたときに403:forbiddenとなったので参照した
* [Go gin framework CORS](https://stackoverflow.com/questions/29418478/go-gin-framework-cors)
 * AllowOriginsでアスタリスクも使える
* [HTTP response status codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)