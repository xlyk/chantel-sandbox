FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /go/bin/chantel-api main.go

FROM alpine:latest
RUN apk --update --no-cache add libc6-compat curl
COPY --from=builder /go/bin/chantel-api /go/bin/chantel-api
ENTRYPOINT ["/go/bin/chantel-api"]
