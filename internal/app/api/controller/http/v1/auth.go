package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"jwtgo/internal/app/api/controller/http/dto"
	"jwtgo/internal/app/api/controller/http/mapper"
	"jwtgo/internal/app/api/controller/http/middleware"
	authPb "jwtgo/internal/pkg/proto/auth"
	"jwtgo/internal/pkg/request"
	"jwtgo/internal/pkg/request/schema"
	"jwtgo/pkg/logging"
)

type AuthController struct {
	authMicroService authPb.AuthServiceClient
	requestValidator *validator.Validate
	logger           *logging.Logger
}

func NewAuthController(
	authMicroService authPb.AuthServiceClient,
	requestValidator *validator.Validate,
	logger *logging.Logger,
) *AuthController {
	return &AuthController{
		authMicroService: authMicroService,
		requestValidator: requestValidator,
		logger:           logger,
	}
}

func (ac *AuthController) Register(apiGroup *gin.RouterGroup) {
	authV1Group := apiGroup.Group("/v1/auth")

	authV1Group.POST("/signup", middleware.Validator[dto.UserCredentialsDTO](ac.requestValidator), ac.SignUp())
	authV1Group.POST("/signin", middleware.Validator[dto.UserCredentialsDTO](ac.requestValidator), ac.SignIn())
	authV1Group.POST("/signout", ac.SignOut())
	authV1Group.POST("/refresh", ac.Refresh())
}

func (ac *AuthController) handleError(c *gin.Context, err error, defaultMessage string) {
	statusData, ok := status.FromError(err)
	if !ok {
		ac.logger.Error(defaultMessage+": ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	message := statusData.Message()

	switch statusData.Code() {
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"message": message})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"message": message})
	case codes.NotFound:
		c.JSON(http.StatusUnauthorized, gin.H{"message": message})
	default:
		ac.logger.Error(defaultMessage+": ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": message})
	}
}

func (ac *AuthController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userCredentialsDTO := c.MustGet("validatedBody").(dto.UserCredentialsDTO)
		signUpRequest := mapper.MapUserCredentialsDTOToAuthSignUpRequest(&userCredentialsDTO)

		signUpResponse, err := ac.authMicroService.SignUp(ctx, signUpRequest)
		if err != nil {
			ac.handleError(c, err, "Error while sign up")
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": signUpResponse.Message})
	}
}

func (ac *AuthController) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userCredentialsDTO := c.MustGet("validatedBody").(dto.UserCredentialsDTO)
		signInRequest := mapper.MapUserCredentialsDTOToAuthSignInRequest(&userCredentialsDTO)

		signInResponse, err := ac.authMicroService.SignIn(ctx, signInRequest)
		if err != nil {
			ac.handleError(c, err, "Error while sign in")
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
			ac.handleError(c, err, "Error while sign out")
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
			ac.handleError(c, err, "Error while refresh")
			return
		}

		request.SetCookies(c, []schema.Cookie{
			{Name: "access_token", Value: refreshResponse.AccessToken, Duration: 7 * 24 * time.Hour},
			{Name: "refresh_token", Value: refreshResponse.RefreshToken, Duration: 7 * 24 * time.Hour},
		})

		c.JSON(http.StatusOK, gin.H{"message": refreshResponse.Message})
	}
}
