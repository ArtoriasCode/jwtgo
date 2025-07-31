package user

import (
	"context"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"jwtgo/internal/app/user/adapter/mongodb/repository"
	"jwtgo/internal/app/user/config"
	server "jwtgo/internal/app/user/controller/grpc/v1"
	userRepositoryIface "jwtgo/internal/app/user/interface/repository"
	userServiceIface "jwtgo/internal/app/user/interface/service"
	userService "jwtgo/internal/app/user/service"
	errorService "jwtgo/internal/pkg/error"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	userPb "jwtgo/internal/pkg/proto/user"
	"jwtgo/pkg/client"
	"jwtgo/pkg/logging"
)

type UserMicroService struct {
	Config         *config.Config
	Logger         *logging.Logger
	Router         *gin.Engine
	Validator      *validator.Validate
	MongoClient    *mongo.Client
	UserRepository userRepositoryIface.UserRepositoryIface
	UserService    userServiceIface.UserServiceIface
	ErrorService   pkgServiceIface.ErrorServiceIface
}

func NewUserMicroService() *UserMicroService {
	logger := logging.GetLogger("info")

	return &UserMicroService{
		Logger: &logger,
	}
}

func (app *UserMicroService) InitializeConfig() {
	app.Logger.Info("Reading application config...")
	app.Config = config.GetConfig(app.Logger)
}

func (app *UserMicroService) InitializeDatabaseClient() {
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

func (app *UserMicroService) InitializeClients() {
	app.InitializeDatabaseClient()
}

func (app *UserMicroService) InitializeUserRepository() {
	app.UserRepository = repository.NewUserRepository(
		app.MongoClient,
		app.Config.MongoDB.Database,
		"users",
		app.Logger,
	)

	ctx := context.Background()

	if err := app.UserRepository.PrepareDatabase(ctx); err != nil {
		app.Logger.Fatal("Failed to prepare user collection: ", err)
	}
}

func (app *UserMicroService) InitializeUserRepositories() {
	app.InitializeUserRepository()
}

func (app *UserMicroService) InitializeUserService() {
	app.UserService = userService.NewUserService(app.UserRepository, app.Logger)
}

func (app *UserMicroService) InitializeErrorService() {
	app.ErrorService = errorService.NewErrorService()
}

func (app *UserMicroService) InitializeServices() {
	app.InitializeErrorService()
	app.InitializeUserService()
}

func (app *UserMicroService) Initialize() {
	app.InitializeConfig()
	app.InitializeClients()
	app.InitializeUserRepositories()
	app.InitializeServices()
}

func (app *UserMicroService) Run() {
	grpcServer := grpc.NewServer()
	userPb.RegisterUserServiceServer(grpcServer, server.NewUserServer(app.UserService, app.ErrorService, app.Logger))

	listener, err := net.Listen("tcp", ":"+app.Config.App.Port)
	if err != nil {
		app.Logger.Fatal("Failed to start the User microservice: ", err)
	}

	app.Logger.Info("User microservice is running on port :" + app.Config.App.Port)

	if err := grpcServer.Serve(listener); err != nil {
		app.Logger.Fatal("Failed to serve gRPC server: ", err)
	}
}
