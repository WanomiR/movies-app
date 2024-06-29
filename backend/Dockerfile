FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod init backend && go mod tidy

RUN go build -o server ./cmd/api

# second stage
FROM alpine:latest

COPY --from=builder /app/server /server
COPY --from=builder /app/.env /.env

EXPOSE 8888

CMD ["/server"]
