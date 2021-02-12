FROM golang:latest as builder
MAINTAINER Tony<asong4love@gmail.com>
ENV APP_NAME="simple_server"
ARG BUILD_WORK_DIR="/srv/$APP_NAME/"
COPY . $BUILD_WORK_DIR
WORKDIR $BUILD_WORK_DIR
RUN CGO_ENABLED=0 go build -o bin/$APP_NAME

FROM alpine:latest
ENV APP_NAME="simple_server"
ENV user d0zingcat
ENV password Hello@World11235
ARG BUILD_WORK_DIR="/srv/$APP_NAME/"
ARG BIZ_WORK_DIR="/srv/"
WORKDIR $BIZ_WORK_DIR
RUN apk --no-cache add ca-certificates
COPY . $BIZ_WORK_DIR
COPY --from=builder $BUILD_WORK_DIR/bin/"$APP_NAME" $BIZ_WORK_DIR/
ENTRYPOINT ["/bin/sh", "-c", "./$APP_NAME -u $user -p $password"]
