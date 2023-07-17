
#protoc --go-grpc_out=.. --proto_path=./pb hello.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pb/*.proto