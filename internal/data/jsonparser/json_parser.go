package jsonparser

import (
	"github.com/photowey/tsc/internal/pkg/jsonz"
)

func ParseMap[T any](jsonString string) map[string]T {
	return ParseMapB[T]([]byte(jsonString))
}

func ParseMapB[T any](jsonByte []byte) map[string]T {
	return jsonz.UnmarshalMap[T](jsonByte)
}
