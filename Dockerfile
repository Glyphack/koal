FROM golang:1.17-alpine

RUN apk add --no-cache git make build-base

WORKDIR /app/koal

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/koal ./cmd/main.go

CMD ["./out/koal"]
