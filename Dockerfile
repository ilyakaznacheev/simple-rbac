FROM golang:1.20 as builder

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -o /usr/local/bin/app cmd/main.go

FROM alpine:3.16

COPY --from=builder /usr/local/bin/app /app

CMD ["/app"]
