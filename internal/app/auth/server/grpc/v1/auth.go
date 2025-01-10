package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"jwtgo/internal/app/auth/server/grpc/mapper"
	customErr "jwtgo/internal/pkg/error"
	serviceInterface "jwtgo/internal/pkg/interface/service"
	pb "jwtgo/internal/proto/auth"
	"jwtgo/pkg/logging"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	authService serviceInterface.AuthService
	logger      *logging.Logger
}

func NewAuthServer(authService serviceInterface.AuthService, logger *logging.Logger) *AuthServer {
	return &AuthServer{
		authService: authService,
		logger:      logger,
	}
}

func (s *AuthServer) handeError(err error) codes.Code {
	var alreadyExistsErr *customErr.AlreadyExistsError
	var invalidCredentialsErr *customErr.InvalidCredentialsError
	var invalidTokenError *customErr.InvalidTokenError
	var expiredTokenError *customErr.ExpiredTokenError
	var notFoundError *customErr.NotFoundError

	var statusCode codes.Code

	switch {
	case errors.As(err, &alreadyExistsErr):
		statusCode = codes.AlreadyExists
	case errors.As(err, &invalidCredentialsErr), errors.As(err, &invalidTokenError), errors.As(err, &expiredTokenError):
		statusCode = codes.Unauthenticated
	case errors.As(err, &notFoundError):
		statusCode = codes.NotFound
	default:
		statusCode = codes.Internal
	}

	return statusCode
}

func (s *AuthServer) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	userCredentialsDTO := mapper.MapSignUpRequestToUserCredentialsDTO(request)

	_, err := s.authService.SignUp(ctx, userCredentialsDTO)
	if err != nil {
		return nil, status.Errorf(s.handeError(err), err.Error())
	}

	return &pb.SignUpResponse{Message: "User successfully registered"}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	userCredentialsDTO := mapper.MapSignInRequestToUserCredentialsDTO(request)

	userTokensDTO, err := s.authService.SignIn(ctx, userCredentialsDTO)
	if err != nil {
		return nil, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapUserTokensDTOToSignInResponse(userTokensDTO, "User successfully logged in"), nil
}

func (s *AuthServer) SignOut(ctx context.Context, request *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	userAccessTokenDTO := mapper.MapSignOutRequestToUserTokenDTO(request)

	err := s.authService.SignOut(ctx, userAccessTokenDTO)
	if err != nil {
		return nil, status.Errorf(s.handeError(err), err.Error())
	}

	return &pb.SignOutResponse{Message: "User successfully logged out"}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	userRefreshTokenDTO := mapper.MapRefreshRequestToUserTokenDTO(request)

	userTokensDTO, err := s.authService.Refresh(ctx, userRefreshTokenDTO)
	if err != nil {
		return nil, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapUserTokensDTOToRefreshResponse(userTokensDTO, "Tokens successfully updated"), nil
}
