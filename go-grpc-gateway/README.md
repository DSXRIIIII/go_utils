学习链接 https://www.cnblogs.com/hacker-linner/p/14618862.html

``` go

go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

// 下载buf
go install github.com/bufbuild/buf/cmd/buf@v1.52.1

// grpc-gen-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2

// 生成gateway文件
protoc -I=proto  --go_out=proto --go_opt=paths=source_relative  --go-grpc_out=proto --go-grpc_opt=paths=source_relative --grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative  proto/hello.proto 

```