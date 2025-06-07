FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod init rss2mail-ai && go mod tidy && go build -o rss2mail

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/rss2mail /usr/bin/rss2mail
COPY config.yaml .
CMD ["rss2mail"]
