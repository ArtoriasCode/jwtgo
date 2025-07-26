package service

import (
	"google.golang.org/grpc/codes"

	customErr "jwtgo/internal/pkg/error/type"
)

type ErrorServiceIface interface {
	ErrToGrpcCode(err customErr.BaseErrorIface) codes.Code
	GrpcCodeToHttpErr(err error) (int, string)
}
