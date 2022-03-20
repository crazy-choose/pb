#!/usr/bin/env bash

SRC_DIR="./"
GO_DST_DIR="../../"
CPP_DST_DIR="../../out/cpp"
META="*.proto"

protoc -I=$SRC_DIR --go_out=$GO_DST_DIR $SRC_DIR/$META
protoc -I=$SRC_DIR --cpp_out=$CPP_DST_DIR $SRC_DIR/$META
