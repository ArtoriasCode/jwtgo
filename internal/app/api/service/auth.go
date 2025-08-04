package service

import (
	"context"

	"jwtgo/internal/app/api/controller/http/dto"
	"jwtgo/internal/app/api/controller/http/mapper"
	authPb "jwtgo/internal/generated/proto/auth"
	customErr "jwtgo/internal/pkg/error/type"
	"jwtgo/pkg/logging"
)

type AuthService struct {
	authMicroService authPb.AuthServiceClient
	logger           *logging.Logger
}

func NewAuthService(
	authMicroService authPb.AuthServiceClient,
	logger *logging.Logger,
) *AuthService {
	return &AuthService{
		authMicroService: authMicroService,
		logger:           logger,
	}
}

func (s *AuthService) SignUp(ctx context.Context, signUpRequestDTO *dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, customErr.BaseErrorIface) {
	authSignUpRequest := mapper.MapSignUpRequestDTOToAuthSignUpRequest(signUpRequestDTO)

	authSignUpResponse, err := s.authMicroService.SignUp(ctx, authSignUpRequest)
	if err != nil {
		return nil, err
	}

	return mapper.MapAuthSignUpResponseToSignUpResponseDTO(authSignUpResponse), nil
}

func (s *AuthService) SignIn(ctx context.Context, signInRequestDTO *dto.SignInRequestDTO) (*dto.SignInResponseDTO, customErr.BaseErrorIface) {
	authSignInRequest := mapper.MapSignInRequestDTOToAuthSignInRequest(signInRequestDTO)

	authSignInResponse, err := s.authMicroService.SignIn(ctx, authSignInRequest)
	if err != nil {
		return nil, err
	}

	return mapper.MapAuthSignInResponseToSignInResponseDTO(authSignInResponse), nil
}

func (s *AuthService) SignOut(ctx context.Context, signOutRequestDTO *dto.SignOutRequestDTO) (*dto.SignOutResponseDTO, customErr.BaseErrorIface) {
	authSignOutRequest := mapper.MapSignOutRequestDTOToAuthSignOutRequest(signOutRequestDTO)

	authSignOutResponse, err := s.authMicroService.SignOut(ctx, authSignOutRequest)
	if err != nil {
		return nil, err
	}

	return mapper.MapAuthSignOutResponseToSignOutResponseDTO(authSignOutResponse), nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshRequestDTO *dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, customErr.BaseErrorIface) {
	authRefreshRequest := mapper.MapRefreshRequestDTOToAuthRefreshRequest(refreshRequestDTO)

	authRefreshResponse, err := s.authMicroService.Refresh(ctx, authRefreshRequest)
	if err != nil {
		return nil, err
	}

	return mapper.MapAuthRefreshResponseToRefreshResponseDTO(authRefreshResponse), nil
}
