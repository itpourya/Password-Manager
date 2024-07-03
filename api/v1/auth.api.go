package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nothyphen/Password-Manager/models"
	jwt "github.com/nothyphen/Password-Manager/pkg"
	"github.com/nothyphen/Password-Manager/serilizers"
	"github.com/nothyphen/Password-Manager/services"
)

type AuthAPI interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authAPI struct {
	authService services.AuthService
}

func NewAuthAPI(service services.AuthService) AuthAPI {
	return &authAPI{
		authService: service,
	}
}

func (a authAPI) Register(context *gin.Context) {
	var registerRequest serilizers.RegisterRequest

	err := context.ShouldBind(&registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	user, err := a.authService.AddUserService(registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"data": user, "status": true})
	return
	
}

func (a authAPI) Login(context *gin.Context) {
	var loginRequest serilizers.LoginRequest
	var users 		 models.User

	err := context.ShouldBind(&loginRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}

	user, err := a.authService.LoginVerify(loginRequest.Email, loginRequest.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{user: "password/email incorect", "status": false})
		return
	}
	
	users.Email = loginRequest.Email
	jwt := jwt.Jwt{}
	token, err := jwt.CreateToken(users)
	context.JSON(http.StatusOK, gin.H{"data": token, "status": true})
	return
}