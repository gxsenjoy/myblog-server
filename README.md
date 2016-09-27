# myblog-server

[![CircleCI](https://circleci.com/gh/nomkhonwaan/myblog-server.svg?style=shield)](https://circleci.com/gh/nomkhonwaan/myblog-server)
[![Coverage Status](https://coveralls.io/repos/github/nomkhonwaan/myblog-server/badge.svg?branch=develop)](https://coveralls.io/github/nomkhonwaan/myblog-server?branch=develop)
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

### gRPC RESTful Gateway
```
$ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:. \
 *.proto
```

### Swagger JSON
```
 $ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --swagger_out=logtostderr=true:. \
 *.proto
 ```

## Compile the Source Code

### gRPC Server
```
$ cd $GOPATH/src/github.com/nomkhonwaan/myblog-server/grpc-server
$ go build -o ../bin/grpc_server .
```

### gRPC RESTful Gateway
```
$ cd $GOPATH/src/github.com/nomkhonwaan/myblog-server/grpc-gateway
$ go build -o ../bin/grpc-gateway
```
