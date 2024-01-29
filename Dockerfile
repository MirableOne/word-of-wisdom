FROM golang:1.21-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app ./cmd/server

FROM alpine:3.12
COPY --from=builder /app ./app

CMD ["./app"]