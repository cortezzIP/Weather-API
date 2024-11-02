FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o weather-api cmd/main.go

FROM alpine:latest

COPY --from=builder /app/weather-api .
COPY .env .

EXPOSE 8080

CMD ["./weather-api"]