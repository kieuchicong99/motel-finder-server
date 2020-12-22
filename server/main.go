package main

import (
	"server/controller"
	"server/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	//port := os.Getenv("PORT") // deploy
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + "8081"
	//docs.SwaggerInfo.Host = "go-react-heroku.herokuapp.com" // deploy
	docs.SwaggerInfo.BasePath = "/api/v1"
	//docs.SwaggerInfo.Schemes = []string{"https", "http"} // deploy
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	c := controller.NewController()
	//r.Use(static.Serve("/", static.LocalFile("./web", true))) // old version
	r.Use(CORSMiddleware())
	v1 := r.Group("/api/v1")
	{
		//
		//accounts := v1.Group("/accounts")
		//{
		//	accounts.GET(":id", c.ShowAccount)
		//	accounts.GET("", c.ListAccounts)
		//	accounts.POST("", c.AddAccount)
		//	accounts.DELETE(":id", c.DeleteAccount)
		//	accounts.PATCH(":id", c.UpdateAccount)
		//	accounts.POST(":id/images", c.UploadAccountImage)
		//}

		// motel api router
		motel := v1.Group("/motel")
		{
			motel.GET("", c.GetMotelsByFilter)
			motel.POST("", c.CreateMotel)
			motel.PATCH(":code", c.UpdateMotel)
			motel.GET(":code", c.GetMotelByCode)
		}

		// user api router
		user := v1.Group("/user")
		{
			user.POST("", c.CreateUser)
			user.POST("/login", c.LoginUser)
			user.GET("", c.GetUsersByFilter)
			user.PATCH("/info", c.UpdateUser)
			user.GET(":code", c.GetUserByCode)
			user.PATCH("/change-pass", c.ChangePass)
			user.POST("/add-favourite-motel", c.AddMotelFavourites)
		}
	}
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":" + "8081")
	//r.Run(":" + port) // deploy
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", "0")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
