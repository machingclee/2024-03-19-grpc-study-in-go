#!/bin/bash

MODULE_NAME=github.com/machingclee/gogrpc
protoc -I greet/proto \
    --go_out=. --go_opt=module=$MODULE_NAME \
    --go-grpc_out=. --go-grpc_opt=module=$MODULE_NAME \
    greet/proto/greet.proto