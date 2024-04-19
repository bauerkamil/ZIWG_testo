package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"src/api"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()
	var defaultGroup = router.Group("api/v1")
	api.AddTeacherHandlers(defaultGroup)
	api.AddCourseHandlers(defaultGroup)
	api.AddQuestionHandlers(defaultGroup)
	api.AddAnswerHandlers(defaultGroup)
	api.AddTestHandlers(defaultGroup)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run("localhost:9090")
	if err != nil {
		fmt.Println("Gin error occurred: ", err)
	}
}
