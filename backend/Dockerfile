FROM golang:1.22 AS builder

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /api/cmd/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Etapa de execução
FROM alpine:latest

WORKDIR /api

COPY --from=builder /api/cmd/server/server .

COPY --from=builder /api/cmd/server/.env . 

RUN chmod +x server

CMD ["./server"]