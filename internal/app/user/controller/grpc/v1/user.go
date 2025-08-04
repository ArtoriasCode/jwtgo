package v1

import (
	"context"

	"google.golang.org/grpc/status"

	"jwtgo/internal/app/user/controller/grpc/mapper"
	userServiceIface "jwtgo/internal/app/user/interface/service"
	userPb "jwtgo/internal/generated/proto/user"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
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
	getByIdRequestDTO := mapper.MapUserGetByIdRequestToGetByIdRequestDTO(request)

	userDTO, err := s.userService.GetById(ctx, getByIdRequestDTO)
	if err != nil {
		return &userPb.GetByIdResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.GetByIdResponse{User: nil}, nil
	}

	return mapper.MapUserDTOToUserGetByIdResponse(userDTO), nil
}

func (s *UserServer) GetByEmail(ctx context.Context, request *userPb.GetByEmailRequest) (*userPb.GetByEmailResponse, error) {
	getByEmailRequestDTO := mapper.MapUserGetByEmailRequestToGetByEmailRequestDTO(request)

	userDTO, err := s.userService.GetByEmail(ctx, getByEmailRequestDTO)
	if err != nil {
		return &userPb.GetByEmailResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.GetByEmailResponse{User: nil}, nil
	}

	return mapper.MapUserDTOToUserGetByEmailResponse(userDTO), nil
}

func (s *UserServer) Create(ctx context.Context, request *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	createRequestDTO := mapper.MapUserCreateRequestToCreateRequestDTO(request)

	userDTO, err := s.userService.Create(ctx, createRequestDTO)
	if err != nil {
		return &userPb.CreateResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.CreateResponse{User: nil}, nil
	}

	return mapper.MapUserDTOToUserCreateResponse(userDTO), nil
}

func (s *UserServer) Update(ctx context.Context, request *userPb.UpdateRequest) (*userPb.UpdateResponse, error) {
	updateRequestDTO := mapper.MapUserUpdateRequestToUpdateRequestDTO(request)

	userDTO, err := s.userService.Update(ctx, updateRequestDTO)
	if err != nil {
		return &userPb.UpdateResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.UpdateResponse{User: nil}, nil
	}

	return mapper.MapUserDTOToUserUpdateResponse(userDTO), nil
}

func (s *UserServer) Delete(ctx context.Context, request *userPb.DeleteRequest) (*userPb.DeleteResponse, error) {
	deleteRequestDTO := mapper.MapUserDeleteRequestToDeleteRequestDTO(request)

	userDTO, err := s.userService.Delete(ctx, deleteRequestDTO)
	if err != nil {
		return &userPb.DeleteResponse{}, status.Errorf(s.errorService.ErrToGrpcCode(err), err.Error())
	}

	if userDTO == nil {
		return &userPb.DeleteResponse{User: nil}, nil
	}

	return mapper.MapUserDTOToUserDeleteResponse(userDTO), nil
}
