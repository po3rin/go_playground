# Try Go + AWS + Docker

required file ```/.aws/credentials```

```
[default]
aws_access_key_id = <<ACCESS_KEY>>
aws_secret_access_key = <<SERCRET_KEY>>
```

and set environment values.

```bash
$ source app.env
```

# Go言語 + Docker で AWS と接続できる開発環境を作る

AWS の S3 や DynamoDB に Docker で作った開発環境から接続したかったのでサクッと準備してみたので共有。
今回紹介するサンプルではS3のバケット一覧を所得するところまでやっている。

ぶっちゃけDocker知ってて、ここの公式ドキュメント読めればそんな苦労しない。
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html

## AWS のサービスに繋げるためにそもそも何がいるのか

ACCESS_KEY
SERCRET_KEY
リージョン情報

がいる。ポイントは Key をどのように読み込みか。
Key の読み込みには 4パターンのやり方がある。

1: 環境変数使うパターン
2: credentials file を作ってそこに記載しておくパターン
3: EC2インスタンスで　IAM Roles から接続するパターン
4: コードの中にハードコーディングしちゃうやり方

今回はローカルでの開発環境を想定かつ、必要な証明書ファイルを自動で見つけてくれるやり方は、Dockerと相性がいいように思われる為、今回は2番を採用する。他の方法についての詳細は、公式ドキュメントへ
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html


今回、最終的にはこんな感じになる。

```
.
├── .aws
│   └── credentials
├── .gitignore
├── Dockerfile
├── Makefile
└── main.go
```

さて、まずは早速 ACCESS_KEY と SERCRET=KEY を記載した ./.aws/credentials というファイルを作る、これが証明書になる。

```
[default]
aws_access_key_id = <<ACCESS_KEY>>
aws_secret_access_key = <<SERCRET_KEY>>
```

ちなみに環境別に Key を変えているときは下記のように数パターン記載しておける。こちらで明示的に指定しない限り[default]を読みに行く。別パターンの読み込みはのちに紹介する。

```
[default]
aws_access_key_id = <YOUR_DEFAULT_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_DEFAULT_SECRET_ACCESS_KEY>

[test-account]
aws_access_key_id = <YOUR_TEST_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_TEST_SECRET_ACCESS_KEY>
```

gitに間違えてあげると怖いので ```.gitignore``` に追加しておく

```bash
echo "credentials" >> .gitignore
```

これで credentials file が完成した。しかし、注意したいのが AWS の SDK は、```~/.aws``` にあるファイルしか見に行ってくれない。つまり今回のようにローカルのプロジェクトディレクトリの中に作ったファイルは検出されない。ではなぜ作ったのかというと、Dockerで環境を作る際に、Dockerコンテナ内の```~/.aws```にファイルを置いて、自動で読み込ませる為です。早速Dockerfileを作ります。

```
FROM golang:1.11.0

WORKDIR /go/src/try-go-aws

COPY . .

RUN mkdir -p ~/.aws
RUN mv ./.aws/credentials ~/.aws/

RUN go get -d -v ./...
```

先ほど作った credentials ファイルを Dockerコンテナ内```~/.aws```にセットしているのがわかると思います。これで Docker 内の SDK が AWS に接続できるようになります。

では動かす用の Goファイルを作りましょう。

```go
ackage main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	sess, err := session.NewSession(&aws.Config{
        // ここでリージョン情報を指定　もちろん環境変数にしても良い。
        Region: aws.String("ap-northeast-1"),
        // 必要ならここで credentials のどの Key を使うかを指定できる。
        // Credentials: credentials.NewSharedCredentials("", "test-account"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
```

今回は S3 を使っているが、もし、DynamoDBを使う際は s3.NEW() の部分を dynamo.New() とすれば良い。


あとはDockerコマンドを叩くだけだが、毎回叩くのもめんどいので、Makefileを作っておく。

```bash
IMAGENAME=try-go-aws
PROJECTNAME=try-go-aws

all: build run

.PHONY: build
build: ## build images of go1.11.0 + aws config
	docker build -t $(IMAGENAME) .

.PHONY: run
run: ## run go1.11.0 + aws config container
	docker run --name $(IMAGENAME) --rm -it -p 8081:8081 -v $(shell pwd):/go/src/${PROJECTNAME} $(IMAGENAME) bash

.PHONY: exec
exec: ## exec go1.11.0 + aws config container
	docker exec -it $(IMAGENAME) /bin/bash

.PHONY: help
help: ## show help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
```

Makefile の最後の help target は make のヘルプを表示する為のもの
[超小技!! Makefile に help をつけて「こいつ...できる!」と言われたい](https://qiita.com/po3rin/items/7875ef9db5ca994ff762)

これで準備完了。下記のコマンドでビルドから立ち上げまでやってくれる。

```bash
$ make
```

コンテナが起動するとそのままコンテナ中に入るので main.go を実行してみよう。

```bash
# in docker containe
$ go run main.go
```

これで Key に紐づけられた S3 からバケットの一覧返ってくる。あとは煮るなり焼くなり

