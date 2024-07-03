package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nothyphen/Password-Manager/serilizers"
	"github.com/nothyphen/Password-Manager/services"
)

type UserAPI interface {
	Save(ctx *gin.Context)
	Show(ctx *gin.Context)
}

type userAPI struct {
	userService services.UserService
}

func NewUserAPI(userService services.UserService) UserAPI {
	return &userAPI{
		userService: userService,
	}
}

type Email struct {
	Email string `json:"email" form:"email" binding:"email"`
}


func (a userAPI) Save(ctx *gin.Context) {
	var saveRequest serilizers.SaveRequest
	var email Email

	err := ctx.ShouldBind(&saveRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}

	useremail := ctx.Value("userEmail")
	email = Email{useremail.(string)}
	
	result, err := a.userService.SavePassword(saveRequest, email.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result, "status": true})
	return

}

func (a userAPI) Show(ctx *gin.Context) {
	var email Email

	useremail := ctx.Value("userEmail")
	email = Email{useremail.(string)}

	result, err := a.userService.ListPassword(email.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	fmt.Println(result)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result, "status": true})
	return
}