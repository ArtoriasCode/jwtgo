package error

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	customErr "jwtgo/internal/pkg/error/type"
)

type ErrorService struct{}

func NewErrorService() *ErrorService {
	return &ErrorService{}
}

func (s *ErrorService) ErrToGrpcCode(err customErr.BaseErrorIface) codes.Code {
	var (
		alreadyExistsErr      *customErr.AlreadyExistsError
		invalidCredentialsErr *customErr.InvalidCredentialsError
		invalidTokenErr       *customErr.InvalidTokenError
		expiredTokenErr       *customErr.ExpiredTokenError
		notFoundError         *customErr.NotFoundError
	)

	switch {
	case errors.As(err, &alreadyExistsErr):
		return codes.AlreadyExists
	case errors.As(err, &invalidCredentialsErr), errors.As(err, &invalidTokenErr), errors.As(err, &expiredTokenErr):
		return codes.Unauthenticated
	case errors.As(err, &notFoundError):
		return codes.NotFound
	default:
		return codes.Internal
	}
}

func (s *ErrorService) GrpcCodeToHttpErr(err error) (int, string) {
	statusData, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, err.Error()
	}

	message := statusData.Message()

	switch statusData.Code() {
	case codes.AlreadyExists:
		return http.StatusConflict, message
	case codes.Unauthenticated:
		return http.StatusUnauthorized, message
	case codes.NotFound:
		return http.StatusNotFound, message
	default:
		return http.StatusInternalServerError, message
	}
}
