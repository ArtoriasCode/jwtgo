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

func (ac *AuthController) Register(apiGroup *gin.RouterGroup) {
	authV1Group := apiGroup.Group("/v1/auth")

	authV1Group.POST("/signup", middleware.Validator[dto.UserSignUpDTO](ac.requestValidator), ac.SignUp())
	authV1Group.POST("/signin", middleware.Validator[dto.UserSignInDTO](ac.requestValidator), ac.SignIn())
	authV1Group.POST("/signout", ac.SignOut())
	authV1Group.POST("/refresh", ac.Refresh())
}

func (ac *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userSignUpDTO := c.MustGet("validatedBody").(dto.UserSignUpDTO)
		signUpRequest := mapper.MapUserSignUpDTOToAuthSignUpRequest(&userSignUpDTO)

		signUpResponse, err := ac.authMicroService.SignUp(ctx, signUpRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": signUpResponse.Message})
	}
}

func (ac *AuthController) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userSignInDTO := c.MustGet("validatedBody").(dto.UserSignInDTO)
		signInRequest := mapper.MapUserSignInDTOToAuthSignInRequest(&userSignInDTO)

		signInResponse, err := ac.authMicroService.SignIn(ctx, signInRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: signInResponse.AccessToken, Duration: 7 * 24 * time.Hour},
			{Name: "refresh_token", Value: signInResponse.RefreshToken, Duration: 7 * 24 * time.Hour},
		})

		c.JSON(http.StatusOK, gin.H{"message": signInResponse.Message})
	}
}

func (ac *AuthController) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}

		signOutRequest := mapper.MapAccessTokenToAuthRefreshRequest(accessToken)

		signOutResponse, err := ac.authMicroService.SignOut(ctx, signOutRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: "", Duration: 7 * 24 * time.Hour},
			{Name: "refresh_token", Value: "", Duration: 7 * 24 * time.Hour},
		})

		c.JSON(http.StatusOK, gin.H{"message": signOutResponse.Message})
	}
}

func (ac *AuthController) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}

		refreshRequest := mapper.MapRefreshTokenToAuthRefreshRequest(refreshToken)

		refreshResponse, err := ac.authMicroService.Refresh(ctx, refreshRequest)
		if err != nil {
			code, message := ac.errorService.GrpcCodeToHttpErr(err)
			c.JSON(code, gin.H{"error": message})
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: refreshResponse.AccessToken, Duration: 7 * 24 * time.Hour},
			{Name: "refresh_token", Value: refreshResponse.RefreshToken, Duration: 7 * 24 * time.Hour},
		})

		c.JSON(http.StatusOK, gin.H{"message": refreshResponse.Message})
	}
}
