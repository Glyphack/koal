FROM golang:1.17-alpine AS buildenv

RUN apk add --no-cache git make build-base

WORKDIR /app/koal

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o ./bin ./cmd/main.go

FROM alpine

WORKDIR /app/koal

COPY ./config.env ./
COPY ./api-docs ./api-docs
COPY --from=buildenv  /app/koal/bin ./bin
EXPOSE 8090
CMD ["./bin"]
