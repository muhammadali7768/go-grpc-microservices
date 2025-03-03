FROM golang:1.23.1-alpine3.20 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/muhammadali7768/go-grpc-microservices
COPY go.mod go.sum ./
COPY vendor vendor
COPY account account
RUN go build -mod vendor -o /go/bin/app ./account/cmd/account

FROM alpine:3.20
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
