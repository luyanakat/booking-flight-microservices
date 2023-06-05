# Build Stage
FROM golang:alpine AS build
WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o customer grpc/customer-grpc/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/customer .

COPY --from=build /app/grpc/customer-grpc/config.yml .

CMD ["./customer"]