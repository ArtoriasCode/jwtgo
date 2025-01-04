package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"

	"jwtgo/internal/app/adapter/mongodb/repository"
	"jwtgo/internal/app/config"
	"jwtgo/internal/app/controller/http/middleware"
	"jwtgo/internal/app/controller/http/v1"
	serviceInterface "jwtgo/internal/app/interface/service"
	"jwtgo/internal/app/service"
	"jwtgo/pkg/client"
	"jwtgo/pkg/logging"
)

type Application struct {
	Config          *config.Config
	Logger          *logging.Logger
	Router          *gin.Engine
	Validator       *validator.Validate
	MongoClient     *mongo.Client
	JWTService      serviceInterface.JWTService
	PasswordService serviceInterface.PasswordService
	AuthService     serviceInterface.AuthService
}

func NewApplication() *Application {
	logger := logging.GetLogger("info")

	return &Application{
		Logger: &logger,
	}
}

func (app *Application) InitializeConfig() {
	app.Logger.Info("Reading application config...")
	app.Config = config.GetConfig(app.Logger)
}

func (app *Application) SetGinMode() {
	if app.Config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (app *Application) InitializeRouter() {
	app.Logger.Info("Application initialization...")
	app.Router = gin.New()

	app.Router.Use(gin.Logger())
}

func (app *Application) InitializeClients() {
	app.Validator = validator.New()

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

func (app *Application) InitializeServices() {
	app.JWTService = service.NewJWTService(app.Config.Security.Secret, app.Config.Security.AccessLifetime, app.Config.Security.RefreshLifetime)
	app.PasswordService = service.NewPasswordService(app.Config.Security.BcryptCost, app.Config.Security.Salt)

	userRepository := repository.NewUserRepository(app.MongoClient, app.Config.MongoDB.Database, "users", app.Logger)
	app.AuthService = service.NewAuthService(userRepository, app.JWTService, app.PasswordService, app.Logger)
}

func (app *Application) InitializeControllers() {
	authController := v1.NewAuthController(app.AuthService, app.Validator, app.Logger)
	authController.Register(app.Router)

	app.Router.Use(middleware.Authentication(app.JWTService))
}

func (app *Application) Run() {
	app.Logger.Info("Application is running on http://" + app.Config.App.Host + ":" + app.Config.App.Port)
	err := app.Router.Run(app.Config.App.Host + ":" + app.Config.App.Port)
	if err != nil {
		app.Logger.Fatal("Failed to start the application", err)
	}
}

func (app *Application) Initialize() {
	app.InitializeConfig()
	app.SetGinMode()

	app.InitializeRouter()
	app.InitializeClients()
	app.InitializeServices()
	app.InitializeControllers()
}
