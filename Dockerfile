FROM golang:1.21 AS build
ARG GO111MODULE=on
ARG GOPROXY=https://goproxy.io,direct
WORKDIR /go/src
COPY . .
RUN go mod download
RUN make build-api-server-linux

FROM alpine:latest AS api-server
RUN apk --no-cache add ca-certificates
WORKDIR /go/src
COPY --from=build /go/src/bin ./
CMD ["./api-server"]