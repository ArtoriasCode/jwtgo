package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"jwtgo/internal/app/api/config"
	"jwtgo/internal/app/api/controller/http/middleware"
	"jwtgo/internal/app/api/controller/http/v1"
	errorService "jwtgo/internal/pkg/error"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	jwtService "jwtgo/internal/pkg/jwt"
	authPb "jwtgo/internal/pkg/proto/auth"
	"jwtgo/pkg/logging"
)

type ApiGateway struct {
	Config            *config.Config
	Logger            *logging.Logger
	Router            *gin.Engine
	JWTService        pkgServiceIface.JWTServiceIface
	ErrorService      pkgServiceIface.ErrorServiceIface
	ValidatorClient   *validator.Validate
	AuthServiceClient authPb.AuthServiceClient
}

func NewApiGateway() *ApiGateway {
	logger := logging.GetLogger("info")

	return &ApiGateway{
		Logger: &logger,
	}
}

func (app *ApiGateway) InitializeConfig() {
	app.Logger.Info("Reading API gateway config...")
	app.Config = config.GetConfig(app.Logger)
}

func (app *ApiGateway) initializeValidatorClient() {
	app.ValidatorClient = validator.New()
}

func (app *ApiGateway) initializeAuthServiceClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(app.Config.Service.Auth.Container+":"+app.Config.Service.Auth.Port, opts...)
	if err != nil {
		app.Logger.Fatal("Failed to connect to Auth server: ", err)
	}

	app.AuthServiceClient = authPb.NewAuthServiceClient(conn)
}

func (app *ApiGateway) InitializeClients() {
	app.Logger.Info("Clients initialization...")
	app.initializeValidatorClient()
	app.initializeAuthServiceClient()
}

func (app *ApiGateway) InitializeJWTService() {
	app.JWTService = jwtService.NewJWTService(
		app.Config.Security.Secret,
		app.Config.Security.AccessLifetime,
		app.Config.Security.RefreshLifetime,
		app.Logger,
	)
}

func (app *ApiGateway) InitializeErrorService() {
	app.ErrorService = errorService.NewErrorService()
}

func (app *ApiGateway) InitializeServices() {
	app.InitializeErrorService()
	app.InitializeJWTService()
}

func (app *ApiGateway) SetGinMode() {
	if app.Config.Service.Api.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (app *ApiGateway) InitializeRouter() {
	app.Logger.Info("API gateway initialization...")
	app.Router = gin.New()

	app.Router.Use(gin.Logger())
}

func (app *ApiGateway) InitializeControllers() {
	apiGroup := app.Router.Group("/api")

	authController := v1.NewAuthController(app.AuthServiceClient, app.ErrorService, app.ValidatorClient, app.Logger)
	authController.Register(apiGroup)

	app.Router.Use(middleware.Authentication(app.JWTService, app.AuthServiceClient))
}

func (app *ApiGateway) Run() {
	app.Logger.Info("API gateway is running on http://" + app.Config.Service.Api.Host + ":" + app.Config.Service.Api.Port)
	err := app.Router.Run(app.Config.Service.Api.Host + ":" + app.Config.Service.Api.Port)
	if err != nil {
		app.Logger.Fatal("Failed to start the application", err)
	}
}

func (app *ApiGateway) Initialize() {
	app.InitializeConfig()
	app.SetGinMode()
	app.InitializeRouter()
	app.InitializeClients()
	app.InitializeServices()
	app.InitializeControllers()
}
