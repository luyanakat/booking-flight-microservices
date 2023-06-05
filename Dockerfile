# Build Stage
FROM golang:1.20-alpine AS build
WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/config.yaml .

CMD ["./server"]