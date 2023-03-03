FROM  golang:1.18-alpine as builder

WORKDIR /root
COPY . .

RUN go build

FROM alpine:3.9 as runner

RUN apk --no-cache add ca-certificates curl tzdata

WORKDIR /root

COPY --from=builder /root/fm.auth /root

# for grpc
EXPOSE 6000
# for http
EXPOSE 6001

CMD ["./fm.auth", "serve"]
