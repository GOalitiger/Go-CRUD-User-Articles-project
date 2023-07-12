package routes

import (
	"crudOperations/controllers"
	"crudOperations/shared"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SettingRoutes(e *echo.Echo) {

	//adding  jwtMiddleware to all calls except login
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(shared.CnfigDB.JWTSecretSigningKey),
		AuthScheme: "Bearer",
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "login")
		},
		TokenLookup: "header:x-auth-token",
	}

	e.Use(middleware.JWTWithConfig(JWTConfig))

	//login route
	e.GET("/login", controllers.Login)

	//users route
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/user/:id", controllers.GetUserById)
	e.POST("/user", controllers.CreateUser)
	e.PUT("/user", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUserById)

	//not working
	// e.GET("/user-articles/query", controllers.GetUsersWithThereArticlesUsingJOIN)

	e.GET("/user-articles", controllers.GetUsersWithThereArticlesUsingJOINWithLogic)
	e.GET("/user-article", controllers.GettingUserArticlesTogetherFromJoin)

	//articles route
	e.GET("/articles", controllers.GetAllArticles)
	e.GET("/article/:id", controllers.GetArticleById)
	e.POST("/article", controllers.CreateArticle)
	e.PUT("/article", controllers.UpdateArticle)
	e.DELETE("/article/:id", controllers.DeleteArticleById)

}
