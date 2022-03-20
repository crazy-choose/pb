#!/usr/bin/env bash

SRC_DIR="./"
RPC_DIR="../../"
META="stream.proto"

protoc --go-grpc_out=$RPC_DIR $SRC_DIR/$META
