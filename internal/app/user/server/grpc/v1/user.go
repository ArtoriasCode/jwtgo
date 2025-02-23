package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	serviceInterface "jwtgo/internal/app/user/interface/service"
	"jwtgo/internal/app/user/server/grpc/mapper"
	customErr "jwtgo/internal/pkg/error"
	userPb "jwtgo/internal/pkg/proto/user"
	"jwtgo/pkg/logging"
)

type UserServer struct {
	userPb.UnimplementedUserServiceServer
	userService serviceInterface.UserService
	logger      *logging.Logger
}

func NewUserServer(userService serviceInterface.UserService, logger *logging.Logger) *UserServer {
	return &UserServer{
		userService: userService,
		logger:      logger,
	}
}

func (s *UserServer) handeError(err error) codes.Code {
	var alreadyExistsErr *customErr.AlreadyExistsError
	var invalidCredentialsErr *customErr.InvalidCredentialsError
	var notFoundError *customErr.NotFoundError

	var statusCode codes.Code

	switch {
	case errors.As(err, &alreadyExistsErr):
		statusCode = codes.AlreadyExists
	case errors.As(err, &invalidCredentialsErr):
		statusCode = codes.Unauthenticated
	case errors.As(err, &notFoundError):
		statusCode = codes.NotFound
	default:
		statusCode = codes.Internal
	}

	return statusCode
}

func (s *UserServer) GetById(ctx context.Context, request *userPb.GetByIdRequest) (*userPb.GetByIdResponse, error) {
	userIdDTO := mapper.MapUserGetByIdRequestToUserIdDTO(request)

	userDTO, err := s.userService.GetById(ctx, userIdDTO)
	if err != nil {
		return &userPb.GetByIdResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.GetByIdResponse{}, nil
	}

	return mapper.MapUserDTOToUserGetByIdResponse(userDTO), nil
}

func (s *UserServer) GetByEmail(ctx context.Context, request *userPb.GetByEmailRequest) (*userPb.GetByEmailResponse, error) {
	userEmailDTO := mapper.MapUserGetByEmailRequestToUserEmailDTO(request)

	userDTO, err := s.userService.GetByEmail(ctx, userEmailDTO)
	if err != nil {
		return &userPb.GetByEmailResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.GetByEmailResponse{}, nil
	}

	return mapper.MapUserDTOToUserGetByEmailResponse(userDTO), nil
}

func (s *UserServer) Create(ctx context.Context, request *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	userCreateDTO := mapper.MapUserCreateRequestToUserCreateDTO(request)

	userDTO, err := s.userService.Create(ctx, userCreateDTO)
	if err != nil {
		return &userPb.CreateResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapUserDTOToUserCreateResponse(userDTO), nil
}

func (s *UserServer) Update(ctx context.Context, request *userPb.UpdateRequest) (*userPb.UpdateResponse, error) {
	userUpdateDTO := mapper.MapUserUpdateRequestToUserUpdateDTO(request)

	userDTO, err := s.userService.Update(ctx, userUpdateDTO)
	if err != nil {
		return &userPb.UpdateResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapUserDTOToUserUpdateResponse(userDTO), nil
}

func (s *UserServer) Delete(ctx context.Context, request *userPb.DeleteRequest) (*userPb.DeleteResponse, error) {
	userIdDTO := mapper.MapUserDeleteRequestToUserIdDTO(request)

	isDeleted, err := s.userService.Delete(ctx, userIdDTO)
	if err != nil {
		return &userPb.DeleteResponse{}, status.Errorf(s.handeError(err), err.Error())
	}

	return mapper.MapStatusToDeleteResponse(isDeleted), nil
}
