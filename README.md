# grpc-server

##Command to generate code from proto file
`protoc ./proto/user.proto --go_out=./ --go-grpc_out=./`
`grpcui -plaintext localhost:8080`