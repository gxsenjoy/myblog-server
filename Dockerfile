# To build a Docker Image for the "develop" branch (this is the default) execute:
#
# $ docker build \
#     --build-arg BUILD_APPLICATION=grpc_server \
#     --build-arg BUILD_BRANCH=develop \
#     --build-arg BUILD_COMMIT=xxxxxxx
#     --build-arg BUILD_ENVIRONMENT=staging \
#     --build-arg BUILD_VERSION=0.0.1-alpha.1 \
#     .
# (or)
# $ docker build .
#
# To build a Docker Image for the "master" branch execute:
#
# $ docker build \
#     --build-arg BUILD_APPLICATION=grpc_gateway \
#     --build-arg BUILD_BRANCH=master \
#     --build-arg BUILD_COMMIT=xxxxxxx
#     --build-arg BUILD_ENVIRONMENT=production \
#     --build-arg BUILD_VERSION=0.0.0 \
#     .

FROM alpine:latest

ARG BUILD_APPLICATION=grpc_server
ARG BUILD_COMMIT=xxxxxxx
ARG BUILD_BRANCH=develop
ARG BUILD_ENVIRONMENT=staging
ARG BUILD_VERSION=0.0.0

LABEL com.nomkhonwaan.application=$BUILD_APPLICATION
LABEL com.nomkhonwaan.commit=$BUILD_COMMIT
LABEL com.nomkhonwaan.branch=$BUILD_BRANCH
LABEL com.nomkhonwaan.environment=$BUILD_ENVIRONMENT
LABEL com.nomkhonwaan.version=$BUILD_VERSION

ENV PORT=9090
ENV APPLICATION=$BUILD_APPLICATION

WORKDIR /usr/bin
ADD ./bin/$BUILD_APPLICATION /usr/bin/

EXPOSE $PORT

CMD ["sh", "-c", "/usr/bin/${APPLICATION}"]
