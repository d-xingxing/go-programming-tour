1. 安装protoc 编译器

```shell
$ wget https://github.com/google/protobuf/releases/download/v3.11.2/protobuf-all-3.11.2.zip
$ unzip protobuf-all-3.11.2.zip && cd protobuf-3.11.2/
$ ./configure
$ make
$ make install
```

2. 对应 Go 语言就是 protoc-gen-go 插件安装

```shell
 go get -u google.golang.org/protobuf/protoc-gen-go@v1.3.2
 
 https://grpc.io/docs/languages/go/quickstart/
 
 export PATH="$PATH:$(go env GOPATH)/bin"
或者 sudo mv $GOPATH/bin/protoc-gen-go /usr/local/go/bin/

3. 生成 proto 文件
```shell
go mod init github.com/go-programming-tour/grpc-demo

protoc --go_out=plugins=grpc:. ./proto/*.proto
```




