FROM golang:1.21 as builder

WORKDIR ./bin/stock-services
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/main.go

FROM alpine:latest
WORKDIR .
COPY --from=builder /go/bin/stock-services/main .
COPY --from=builder /go/bin/stock-services/envs/ envs/
CMD ["./main"]