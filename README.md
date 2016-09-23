# myblog-server

[![Stories in Ready](https://badge.waffle.io/nomkhonwaan/myblog-server.svg?label=ready&title=Ready)](http://waffle.io/nomkhonwaan/myblog-server)

## Installation

### Golang Packages
```
$ go get -u golang.org/x/tools/cmd/...
$ go get -u github.com/golang/lint/golint
$ go get -u github.com/nsf/gocode
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### Atom Editor Packages
```
$ apm install go-plus language-docker language-protobuf
```

## Compile the Protobufs

### gRPC Stub
```
$ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
 *.proto
```

### Reverse Proxy Gateway
```
$ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:. \
 *.proto
 ```

 ### Swagger
 ```
 $ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --swagger_out=logtostderr=true:. \
 *.proto
 ```
