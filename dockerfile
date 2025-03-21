# BUILD
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./


RUN --mount=type=secret,id=github_pat git config --global url."https://$(cat /run/secrets/github_pat)@github.com/".insteadOf "https://github.com/"
RUN go mod download

COPY . .

RUN go build -o server main.go

# RUN
FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/server .

# ENV
ENV RUNNING_IN_DOCKER=true

EXPOSE 3005

CMD ["./server", "server"]