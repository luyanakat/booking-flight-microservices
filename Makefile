gen-gql:
	go run github.com/99designs/gqlgen generate
gen-proto:
	protoc --go_out=./pb --go_opt=paths=source_relative \
        --go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
        proto/*.proto