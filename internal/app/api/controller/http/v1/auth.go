package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"jwtgo/internal/app/api/controller/http/dto"
	"jwtgo/internal/app/api/controller/http/mapper"
	"jwtgo/internal/app/api/controller/http/middleware"
	pkgServiceIface "jwtgo/internal/pkg/interface/service"
	authPb "jwtgo/internal/pkg/proto/auth"
	"jwtgo/internal/pkg/request"
	"jwtgo/internal/pkg/request/schema"
	"jwtgo/pkg/logging"
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

type AuthController struct {
	authMicroService authPb.AuthServiceClient
	errorService     pkgServiceIface.ErrorServiceIface
	requestValidator *validator.Validate
	logger           *logging.Logger
}

func NewAuthController(
	authMicroService authPb.AuthServiceClient,
	errorService pkgServiceIface.ErrorServiceIface,
	requestValidator *validator.Validate,
	logger *logging.Logger,
) *AuthController {
	return &AuthController{
		authMicroService: authMicroService,
		errorService:     errorService,
		requestValidator: requestValidator,
		logger:           logger,
	}
}

func (ac *AuthController) RegisterNoAuth(apiGroup *gin.RouterGroup) {
	authV1Group := apiGroup.Group("/v1/auth")

	authV1Group.POST("/signup", middleware.Validator[dto.SignUpRequestDTO](ac.requestValidator), ac.SignUp())
	authV1Group.POST("/signin", middleware.Validator[dto.SignInRequestDTO](ac.requestValidator), ac.SignIn())
	authV1Group.POST("/refresh", ac.Refresh())
}

func (ac *AuthController) RegisterWithAuth(apiGroup *gin.RouterGroup) {
	authV1Group := apiGroup.Group("/v1/auth")

	authV1Group.POST("/signout", ac.SignOut())
}

func (ac *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		signUpRequestDTO := c.MustGet("validatedBody").(dto.SignUpRequestDTO)
		authSignUpRequest := mapper.MapSignUpRequestDTOToAuthSignUpRequest(&signUpRequestDTO)

		authSignUpResponse, err := ac.authMicroService.SignUp(ctx, authSignUpRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": authSignUpResponse.Message})
	}
}

func (ac *AuthController) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		signInRequestDTO := c.MustGet("validatedBody").(dto.SignInRequestDTO)
		authSignInRequest := mapper.MapSignInRequestDTOToAuthSignInRequest(&signInRequestDTO)

		authSignInResponse, err := ac.authMicroService.SignIn(ctx, authSignInRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: authSignInResponse.AccessToken, Duration: Week},
			{Name: "refresh_token", Value: authSignInResponse.RefreshToken, Duration: Week},
		})

		c.JSON(http.StatusOK, gin.H{"message": authSignInResponse.Message})
	}
}

func (ac *AuthController) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		id := c.MustGet("id").(string)
		authSignOutRequest := mapper.MapAccessTokenToAuthSignOutRequest(id)

		authSignOutResponse, err := ac.authMicroService.SignOut(ctx, authSignOutRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: "", Duration: Week},
			{Name: "refresh_token", Value: "", Duration: Week},
		})

		c.JSON(http.StatusOK, gin.H{"message": authSignOutResponse.Message})
	}
}

func (ac *AuthController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		refreshToken, err := c.Cookie("refresh_token")
		if err != nil || refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}

		authRefreshRequest := mapper.MapRefreshTokenToAuthRefreshRequest(refreshToken)

		authRefreshResponse, err := ac.authMicroService.Refresh(ctx, authRefreshRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: authRefreshResponse.AccessToken, Duration: Week},
			{Name: "refresh_token", Value: authRefreshResponse.RefreshToken, Duration: Week},
		})

		c.JSON(http.StatusOK, gin.H{"message": authRefreshResponse.Message})
	}
}
