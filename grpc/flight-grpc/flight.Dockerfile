# Build Stage
FROM golang:alpine AS build
WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o flight grpc/flight-grpc/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/flight .

COPY --from=build /app/grpc/flight-grpc/config.yml .

CMD ["./flight"]