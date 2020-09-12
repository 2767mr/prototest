package prototest

//go:generate protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=. --go-grpc_out=. example.proto
