FROM golang:1.17-alpine

RUN apk add --no-cache git

WORKDIR /app/koal

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/koal ./cmd/main.go

EXPOSE 8090

CMD ["./out/koal"]
