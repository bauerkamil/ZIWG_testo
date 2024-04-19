package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"src/api"
	_ "src/main/docs"
)

// @title           Swagger ZIWG_testo API
// @version         1.0
// @description     REST API for ZIWG_testo project

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
