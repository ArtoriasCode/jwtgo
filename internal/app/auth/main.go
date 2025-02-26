package auth

import (
	"google.golang.org/grpc/credentials/insecure"
	service3 "jwtgo/internal/pkg/interface/service"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"jwtgo/internal/app/auth/config"
	service2 "jwtgo/internal/app/auth/interface/service"
	server "jwtgo/internal/app/auth/server/grpc/v1"
	"jwtgo/internal/app/auth/service"
	authPb "jwtgo/internal/pkg/proto/auth"
	userPb "jwtgo/internal/pkg/proto/user"
	servicePkg "jwtgo/internal/pkg/service"
	"jwtgo/pkg/logging"
)

type AuthMicroService struct {
	Config            *config.Config
	Logger            *logging.Logger
	Router            *gin.Engine
	JWTService        service3.JWTService
	PasswordService   service3.PasswordService
	AuthService       service2.AuthService
	UserServiceClient userPb.UserServiceClient
}

func NewAuthMicroService() *AuthMicroService {
	logger := logging.GetLogger("info")

	return &AuthMicroService{
		Logger: &logger,
	}
}

func (app *AuthMicroService) InitializeConfig() {
	app.Logger.Info("Reading application config...")
	app.Config = config.GetConfig(app.Logger)
}

func (app *AuthMicroService) initializeUserServiceClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(app.Config.Service.User.Container+":"+app.Config.Service.User.Port, opts...)
	if err != nil {
		app.Logger.Fatal("Failed to connect to User server: ", err)
	}

	app.UserServiceClient = userPb.NewUserServiceClient(conn)
}

func (app *AuthMicroService) InitializeClients() {
	app.initializeUserServiceClient()
}

func (app *AuthMicroService) InitializeJWTService() {
	app.JWTService = servicePkg.NewJWTService(app.Config.Security.Secret, app.Config.Security.AccessLifetime, app.Config.Security.RefreshLifetime)
}

func (app *AuthMicroService) InitializePasswordService() {
	app.PasswordService = servicePkg.NewPasswordService(app.Config.Security.BcryptCost, app.Config.Security.Salt)
}

func (app *AuthMicroService) InitializeAuthService() {
	app.AuthService = service.NewAuthService(app.UserServiceClient, app.JWTService, app.PasswordService, app.Logger)
}

func (app *AuthMicroService) InitializeServices() {
	app.InitializeJWTService()
	app.InitializePasswordService()
	app.InitializeAuthService()
}

func (app *AuthMicroService) Initialize() {
	app.InitializeConfig()
	app.InitializeClients()
	app.InitializeServices()
}

func (app *AuthMicroService) Run() {
	grpcServer := grpc.NewServer()
	authPb.RegisterAuthServiceServer(grpcServer, server.NewAuthServer(app.AuthService, app.Logger))

	listener, err := net.Listen("tcp", ":"+app.Config.App.Port)
	if err != nil {
		app.Logger.Fatal("Failed to start the Auth microservice: ", err)
	}

	app.Logger.Info("Auth microservice is running on port :" + app.Config.App.Port)

	if err := grpcServer.Serve(listener); err != nil {
		app.Logger.Fatal("Failed to serve gRPC server: ", err)
	}
}
