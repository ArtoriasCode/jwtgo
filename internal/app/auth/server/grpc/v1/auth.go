package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	serviceInterface "jwtgo/internal/app/auth/interface/service"
	"jwtgo/internal/app/auth/server/grpc/mapper"
	customErr "jwtgo/internal/pkg/error"
	authPb "jwtgo/internal/pkg/proto/auth"
	"jwtgo/pkg/logging"
)

type AuthServer struct {
	authPb.UnimplementedAuthServiceServer
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

func (s *AuthServer) SignUp(ctx context.Context, request *authPb.SignUpRequest) (*authPb.SignUpResponse, error) {
	userCredentialsDTO := mapper.MapAuthSignUpRequestToUserCredentialsDTO(request)

	_, err := s.authService.SignUp(ctx, userCredentialsDTO)
	if err != nil {
		return &authPb.SignUpResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return &authPb.SignUpResponse{Message: "User successfully registered"}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, request *authPb.SignInRequest) (*authPb.SignInResponse, error) {
	userCredentialsDTO := mapper.MapAuthSignInRequestToUserCredentialsDTO(request)

	userTokensDTO, err := s.authService.SignIn(ctx, userCredentialsDTO)
	if err != nil {
		return &authPb.SignInResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapUserTokensDTOToAuthSignInResponse(userTokensDTO, "User successfully logged in"), nil
}

func (s *AuthServer) SignOut(ctx context.Context, request *authPb.SignOutRequest) (*authPb.SignOutResponse, error) {
	userAccessTokenDTO := mapper.MapAuthSignOutRequestToUserTokenDTO(request)

	_, err := s.authService.SignOut(ctx, userAccessTokenDTO)
	if err != nil {
		return &authPb.SignOutResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return &authPb.SignOutResponse{Message: "User successfully logged out"}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, request *authPb.RefreshRequest) (*authPb.RefreshResponse, error) {
	userRefreshTokenDTO := mapper.MapAuthRefreshRequestToUserTokenDTO(request)

	userTokensDTO, err := s.authService.Refresh(ctx, userRefreshTokenDTO)
	if err != nil {
		return &authPb.RefreshResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapUserTokensDTOToAuthRefreshResponse(userTokensDTO, "Tokens successfully updated"), nil
}
