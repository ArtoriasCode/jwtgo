package v1

import (
	"context"

	"google.golang.org/grpc/status"

	"jwtgo/internal/app/auth/controller/grpc/mapper"
	authServiceIface "jwtgo/internal/app/auth/interface/service"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	authPb "jwtgo/internal/pkg/proto/auth"
	"jwtgo/pkg/logging"
)

type AuthServer struct {
	authPb.UnimplementedAuthServiceServer
	authService  authServiceIface.AuthServiceIface
	errorService pkgServiceIface.ErrorServiceIface
	logger       *logging.Logger
}

func NewAuthServer(
	authService authServiceIface.AuthServiceIface,
	errorService pkgServiceIface.ErrorServiceIface,
	logger *logging.Logger,
) *AuthServer {
	return &AuthServer{
		authService:  authService,
		errorService: errorService,
		logger:       logger,
	}
}

func (s *AuthServer) SignUp(ctx context.Context, request *authPb.SignUpRequest) (*authPb.SignUpResponse, error) {
	signUpRequestDTO := mapper.MapAuthSignUpRequestToSignUpRequestDTO(request)

	_, err := s.authService.SignUp(ctx, signUpRequestDTO)
	if err != nil {
		return &authPb.SignUpResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return &authPb.SignUpResponse{Message: "User successfully registered"}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, request *authPb.SignInRequest) (*authPb.SignInResponse, error) {
	signInRequestDTO := mapper.MapAuthSignInRequestToSignInRequestDTO(request)

	userTokensDTO, err := s.authService.SignIn(ctx, signInRequestDTO)
	if err != nil {
		return &authPb.SignInResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return mapper.MapUserTokensDTOToAuthSignInResponse(userTokensDTO, "User successfully logged in"), nil
}

func (s *AuthServer) SignOut(ctx context.Context, request *authPb.SignOutRequest) (*authPb.SignOutResponse, error) {
	userAccessTokenDTO := mapper.MapAuthSignOutRequestToUserTokenDTO(request)

	_, err := s.authService.SignOut(ctx, userAccessTokenDTO)
	if err != nil {
		return &authPb.SignOutResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return &authPb.SignOutResponse{Message: "User successfully logged out"}, nil
}

func (s *AuthServer) Refresh(ctx context.Context, request *authPb.RefreshRequest) (*authPb.RefreshResponse, error) {
	userRefreshTokenDTO := mapper.MapAuthRefreshRequestToUserTokenDTO(request)

	userTokensDTO, err := s.authService.Refresh(ctx, userRefreshTokenDTO)
	if err != nil {
		return &authPb.RefreshResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return mapper.MapUserTokensDTOToAuthRefreshResponse(userTokensDTO, "Tokens successfully updated"), nil
}
