package user

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"jwtgo/internal/app/user/adapter/mongodb/repository"
	"jwtgo/internal/app/user/config"
	serviceInterface "jwtgo/internal/app/user/interface/service"
	server "jwtgo/internal/app/user/server/grpc/v1"
	"jwtgo/internal/app/user/service"
	userPb "jwtgo/internal/pkg/proto/user"
	"jwtgo/pkg/client"
	"jwtgo/pkg/logging"
)

type UserMicroservice struct {
	Config      *config.Config
	Logger      *logging.Logger
	Router      *gin.Engine
	Validator   *validator.Validate
	MongoClient *mongo.Client
	UserService serviceInterface.UserService
}

func NewUserMicroservice() *UserMicroservice {
	logger := logging.GetLogger("info")

	return &UserMicroservice{
		Logger: &logger,
	}
}

func (app *UserMicroservice) InitializeConfig() {
	app.Logger.Info("Reading application config...")
	app.Config = config.GetConfig(app.Logger)
}

func (app *UserMicroservice) InitializeClients() {
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

func (app *UserMicroservice) InitializeServices() {
	userRepository := repository.NewUserRepository(app.MongoClient, app.Config.MongoDB.Database, "users", app.Logger)
	app.UserService = service.NewUserService(userRepository, app.Logger)
}

func (app *UserMicroservice) Initialize() {
	app.InitializeConfig()
	app.InitializeClients()
	app.InitializeServices()
}

func (app *UserMicroservice) Run() {
	grpcServer := grpc.NewServer()
	userPb.RegisterUserServiceServer(grpcServer, server.NewUserServer(app.UserService, app.Logger))

	listener, err := net.Listen("tcp", ":"+app.Config.App.Port)
	if err != nil {
		app.Logger.Fatal("Failed to start the User microservice: ", err)
	}

	app.Logger.Info("User microservice is running on port :" + app.Config.App.Port)

	if err := grpcServer.Serve(listener); err != nil {
		app.Logger.Fatal("Failed to serve gRPC server: ", err)
	}
}
