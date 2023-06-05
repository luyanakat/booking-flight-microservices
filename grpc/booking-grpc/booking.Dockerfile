# Build Stage
FROM golang:alpine AS build
WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o booking grpc/booking-grpc/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/booking .

COPY --from=build /app/grpc/booking-grpc/config.yml .

CMD ["./booking"]