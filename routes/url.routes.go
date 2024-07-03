package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nothyphen/Password-Manager/api/v1"
	"github.com/nothyphen/Password-Manager/db"
	"github.com/nothyphen/Password-Manager/middleware"
	jwt "github.com/nothyphen/Password-Manager/pkg"
	"github.com/nothyphen/Password-Manager/repository"
	"github.com/nothyphen/Password-Manager/services"
	"gorm.io/gorm"
)

var (
	sqlitedb			*gorm.DB						= db.ConnectDB()
	authrepository		repository.AuthRepository		= repository.NewAuthRepository(sqlitedb)
	authservice			services.AuthService			= services.NewAuthService(authrepository)
	authAPI 			v1.AuthAPI						= v1.NewAuthAPI(authservice)
	userrepository		repository.UserRepository		= repository.NewUserRepository(sqlitedb)
	userservice			services.UserService			= services.NewUserService(userrepository)
	userAPI				v1.UserAPI						= v1.NewUserAPI(userservice)
	jwtauth				jwt.Jwt
)

func Urls() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.CORSMiddleware())
	route.Use(middleware.NoRouteHandler())
	route.HandleMethodNotAllowed = true
	route.Use(middleware.NoMethodHandler())

	apiv1 := route.Group("/api/v1")
	{
		auth := apiv1.Group("auth")
		{
			auth.POST("/register", authAPI.Register)
			auth.POST("/login", authAPI.Login)
			//auth.POST("/forget", authAPI.Forget)
			//auth.POST("/verify", authAPI.Verify)
		}

		user := route.Group("user", middleware.AthorizationJWT(jwtauth))
		{
			user.POST("/save", userAPI.Save)
			user.GET("/list", userAPI.Show)
			//user.DELETE("/delete")
			//user.POST("/update")
		}
	}


	return route
}