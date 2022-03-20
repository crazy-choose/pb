#!/usr/bin/env bash

SRC_DIR="./"
RPC_DIR="../../out/cpp"
META="stream.proto"

protoc -I $SRC_DIR --grpc_out=$RPC_DIR --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $SRC_DIR$META
