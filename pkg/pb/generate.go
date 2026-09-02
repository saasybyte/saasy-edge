package pb

//go:generate protoc --go_out=../.. --go_opt=module=github.com/saasybyte/saasy-edge --go-grpc_out=../.. --go-grpc_opt=module=github.com/saasybyte/saasy-edge --proto_path=../../saasy-proto/protos ../../saasy-proto/protos/edge/v1/edge.proto
