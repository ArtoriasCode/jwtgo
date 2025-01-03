package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"

	"jwtgo/internal/adapter/mongodb/repository"
	"jwtgo/internal/config"
	"jwtgo/internal/controller/http/middleware"
	v1 "jwtgo/internal/controller/http/v1"
	"jwtgo/internal/domain/service"
	"jwtgo/pkg/client"
	"jwtgo/pkg/logging"
	"jwtgo/pkg/security"
)

type Application struct {
	Config          *config.Config
	Logger          *logging.Logger
	Router          *gin.Engine
	Validator       *validator.Validate
	MongoClient     *mongo.Client
	TokenManager    *security.TokenManager
	PasswordManager *security.PasswordManager
	AuthService     *service.AuthService
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
	app.MongoClient = client.NewMongodbClient(app.Config.MongoDB.Url, app.Logger).Connect()
	app.TokenManager = security.NewTokenManager(app.Config.Security.Secret, app.Config.Security.AccessLifetime, app.Config.Security.RefreshLifetime)
	app.PasswordManager = security.NewPasswordManager(app.Config.Security.BcryptCost, app.Config.Security.Salt)
}

func (app *Application) InitializeServices() {
	userRepository := repository.NewUserRepository(app.MongoClient, app.Config.MongoDB.Database, "users", app.Logger)
	app.AuthService = service.NewAuthService(userRepository, app.TokenManager, app.PasswordManager, app.Logger)
}

func (app *Application) InitializeControllers() {
	authController := v1.NewAuthController(app.AuthService, app.Validator, app.PasswordManager, app.Logger)
	authController.Register(app.Router)

	app.Router.Use(middleware.Authentication(app.TokenManager))
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
