package auth

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"jwtgo/internal/app/auth/adapter/mongodb/repository"
	"jwtgo/internal/app/auth/config"
	server "jwtgo/internal/app/auth/server/grpc/v1"
	"jwtgo/internal/app/auth/service"
	serviceInterface "jwtgo/internal/pkg/interface/service"
	servicePkg "jwtgo/internal/pkg/service"
	pb "jwtgo/internal/proto/auth"
	"jwtgo/pkg/client"
	"jwtgo/pkg/logging"
)

type AuthMicroservice struct {
	Config          *config.Config
	Logger          *logging.Logger
	Router          *gin.Engine
	Validator       *validator.Validate
	MongoClient     *mongo.Client
	JWTService      serviceInterface.JWTService
	PasswordService serviceInterface.PasswordService
	AuthService     serviceInterface.AuthService
}

func NewAuthMicroservice() *AuthMicroservice {
	logger := logging.GetLogger("info")

	return &AuthMicroservice{
		Logger: &logger,
	}
}

func (app *AuthMicroservice) InitializeConfig() {
	app.Logger.Info("Reading application config...")
	app.Config = config.GetConfig(app.Logger)
}

func (app *AuthMicroservice) InitializeClients() {
	databaseUrl := fmt.Sprintf(
		"%s://%s:%s@%s:%d/",
		app.Config.MongoDB.Uri,
		app.Config.MongoDB.User,
		app.Config.MongoDB.Password,
		app.Config.MongoDB.Host,
		app.Config.MongoDB.Port,
	)
	app.MongoClient = client.NewMongodbClient(databaseUrl, app.Logger).Connect()
}

func (app *AuthMicroservice) InitializeServices() {
	app.JWTService = servicePkg.NewJWTService(app.Config.Security.Secret, app.Config.Security.AccessLifetime, app.Config.Security.RefreshLifetime)
	app.PasswordService = servicePkg.NewPasswordService(app.Config.Security.BcryptCost, app.Config.Security.Salt)

	userRepository := repository.NewUserRepository(app.MongoClient, app.Config.MongoDB.Database, "users", app.Logger)
	app.AuthService = service.NewAuthService(userRepository, app.JWTService, app.PasswordService, app.Logger)
}

func (app *AuthMicroservice) Initialize() {
	app.InitializeConfig()
	app.InitializeClients()
	app.InitializeServices()
}

func (app *AuthMicroservice) Run() {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server.NewAuthServer(app.AuthService, app.Logger))

	listener, err := net.Listen("tcp", ":"+app.Config.App.Port)
	if err != nil {
		app.Logger.Fatal("Failed to start the Auth microservice: ", err)
	}

	app.Logger.Info("Auth microservice is running on port :" + app.Config.App.Port)

	if err := grpcServer.Serve(listener); err != nil {
		app.Logger.Fatal("Failed to serve gRPC server: ", err)
	}
}
