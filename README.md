# myblog-server

[![CircleCI](https://circleci.com/gh/nomkhonwaan/myblog-server.svg?style=shield)](https://circleci.com/gh/nomkhonwaan/myblog-server)
[![Coverage Status](https://coveralls.io/repos/github/nomkhonwaan/myblog-server/badge.svg?branch=develop)](https://coveralls.io/github/nomkhonwaan/myblog-server?branch=develop)
[![Stories in Ready](https://badge.waffle.io/nomkhonwaan/myblog-server.svg?label=ready&title=Ready)](http://waffle.io/nomkhonwaan/myblog-server)

## Installation

### Golang Packages
```
$ go get -u github.com/tools/godep
$ godep restore
```

### Development Packages
```
$ go get -u golang.org/x/tools/cmd/...
$ go get -u github.com/golang/lint/golint
$ go get -u github.com/nsf/gocode
$ apm install go-plus language-docker language-protobuf
```

## Run the Tests
```
$ godep go test ./blog-service/...
$ godep go test ./grpc-gateway/...
$ godep go test ./grpc-server/...
```

## Compile the Protobufs

### gRPC Stub
```
$ protoc \
    -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
    *.proto
```

### gRPC RESTful Gateway
```
$ protoc \
    -I/usr/local/include -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    *.proto
```

### Swagger JSON
```
 $ protoc \
    -I/usr/local/include -I. \
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

## SSL

Generate self-signed certificate using CloudFlare SSL tools
```
$ cd $GOPATH/src/github.com/nomkhonwaan/myblog-server/ssl
$ cfssl gencert -initca myblog-csr-ca.json | cfssljson -bare myblog
```

Then copy the certificate and key by encode it with base64 algorithm

```
$ cat myblog.pem | base64 | pbcopy
```

put the certificate base64 to `myblog-sercret.yaml tls.crt` and do it same on the certificate key file

```
$ cat myblog-key.pem | base64 | pbcopy
```

And then edit the `myblog-ingress.yaml` in `spec.tls` with `secretName` field

```yaml
...
spec:
  tls:
  - secretName: myblog
...
```

## Deployment
### Kubernetes
#### Create the Deployment
```
$ kubectl create -f myblog-server-deployment.yaml
$ kubectl create -f myblog-gateway-deployment.yaml
```

#### Create the Service
```
$ kubectl create -f myblog-server-service.yaml
$ kubectl create -f myblog-gateway-service.yaml
```

#### Create the Secert
```
$ kubectl create -f myblog-secret.yaml
```

#### Create the Ingress
```
$ kubectl create -f myblog-ingress.yaml
```
