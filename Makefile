proto-compile:
	protoc -I proto/ \
	proto/mission_generator.proto \
	--go_out=./mission-generator \
	--go-grpc_out=./mission-generator  


	protoc -I proto/ \
	proto/mission_generator.proto \
	--php_out=./api-gateway/grpc \
	--grpc_out=./api-gateway/grpc \
	--plugin=protoc-gen-grpc=grpc_php_plugin

proto-gateway-compile:
	protoc -I proto/ --grpc-gateway_out pb/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	proto/mission_generator.proto

buf-compile:
	@cd proto
	buf build 