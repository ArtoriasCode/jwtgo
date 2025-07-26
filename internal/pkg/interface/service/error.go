package service

import (
	"google.golang.org/grpc/codes"

	customErr "jwtgo/internal/pkg/error/type"
)

type ErrorService interface {
	ErrToGrpcCode(err customErr.BaseErrorInterface) codes.Code
	GrpcCodeToHttpErr(err error) (int, string)
}
