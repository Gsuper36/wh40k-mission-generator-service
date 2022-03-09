proto-compile:
	protoc -I pb/ \
	pb/mission_generator.proto \
	--go_out=. \
	--go-grpc_out=. 

proto-gateway-compile:
	protoc -I pb/ --grpc-gateway_out pb/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	pb/mission_generator.proto

buf-compile:
	