package v1

import (
	"context"
	"google.golang.org/grpc/status"

	"jwtgo/internal/app/user/controller/grpc/mapper"
	userServiceIface "jwtgo/internal/app/user/interface/service"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	userPb "jwtgo/internal/pkg/proto/user"
	"jwtgo/pkg/logging"
)

type UserServer struct {
	userPb.UnimplementedUserServiceServer
	userService  userServiceIface.UserServiceIface
	errorService pkgServiceIface.ErrorServiceIface
	logger       *logging.Logger
}

func NewUserServer(
	userService userServiceIface.UserServiceIface,
	errorService pkgServiceIface.ErrorServiceIface,
	logger *logging.Logger,
) *UserServer {
	return &UserServer{
		userService:  userService,
		errorService: errorService,
		logger:       logger,
	}
}

func (s *UserServer) GetById(ctx context.Context, request *userPb.GetByIdRequest) (*userPb.GetByIdResponse, error) {
	userIdDTO := mapper.MapUserGetByIdRequestToUserIdDTO(request)

	userDTO, err := s.userService.GetById(ctx, userIdDTO)
	if err != nil {
		return &userPb.GetByIdResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
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
		return &userPb.GetByEmailResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
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
		return &userPb.CreateResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return mapper.MapUserDTOToUserCreateResponse(userDTO), nil
}

func (s *UserServer) Update(ctx context.Context, request *userPb.UpdateRequest) (*userPb.UpdateResponse, error) {
	userUpdateDTO := mapper.MapUserUpdateRequestToUserUpdateDTO(request)

	userDTO, err := s.userService.Update(ctx, userUpdateDTO)
	if err != nil {
		return &userPb.UpdateResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return mapper.MapUserDTOToUserUpdateResponse(userDTO), nil
}

func (s *UserServer) Delete(ctx context.Context, request *userPb.DeleteRequest) (*userPb.DeleteResponse, error) {
	userIdDTO := mapper.MapUserDeleteRequestToUserIdDTO(request)

	isDeleted, err := s.userService.Delete(ctx, userIdDTO)
	if err != nil {
		return &userPb.DeleteResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	return mapper.MapStatusToDeleteResponse(isDeleted), nil
}
