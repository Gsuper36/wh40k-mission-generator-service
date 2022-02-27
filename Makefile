proto-compile:
	protoc -I pb/ \
	pb/mission_generator.proto \
	--go_out=. \
	--go-grpc_out=.