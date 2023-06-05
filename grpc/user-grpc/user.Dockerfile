# Build Stage
FROM golang:alpine AS build
WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o user grpc/user-grpc/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/user .

COPY --from=build /app/grpc/user-grpc/config.yml .

CMD ["./user"]