FROM golang:1.17-alpine AS buildenv

RUN apk add --no-cache git make build-base

WORKDIR /app/koal

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /koal/bin ./cmd/main.go

FROM alpine

WORKDIR /

COPY --from=buildenv  /koal/bin /koal/bin

CMD ["/koal-bin"]
