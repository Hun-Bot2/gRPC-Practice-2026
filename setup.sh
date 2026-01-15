#!/bin/bash

echo "=== Go gRPC 도구 설치 시작 ==="

# Protoc 플러그인 설치
echo "Installing protoc-gen-go..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

echo "Installing protoc-gen-go-grpc..."
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo "=== 설치 완료! ==="