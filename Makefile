include .makeenv

SHELL = /bin/sh
UID := $(shell id -u)
GID := $(shell id -g)

export APP_NAME
export API_PORT
export MG_PORT
export DB_PORT
export DB_PASSWORD
export DB_USER
export UID
export GID

build:
	docker-compose build
	@mkdir -p ./docker/var/postgres/dbdata
	@chown -R ${UID}:${GID} ./docker/var/postgres

up:
	docker-compose up -d --no-build

down: 
	docker-compose down

builder-proto-compile:
	protoc -I proto/ \
	proto/mission_generator.proto \
	--go_out=./mission-generator \
	--go-grpc_out=./mission-generator  


	protoc -I proto/ \
	proto/mission_generator.proto \
	--php_out=./api-gateway/grpc \
	--grpc_out=./api-gateway/grpc \
	--plugin=protoc-gen-grpc=grpc_php_plugin

builder-proto-gateway-compile:
	protoc -I proto/ --grpc-gateway_out pb/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	proto/mission_generator.proto

builder-buf-compile:
	@cd proto
	buf build 