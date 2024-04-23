package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
	"src/model"
	"time"
)

// AddTest            godoc
// @Summary      Add test
// @Description  Add test from json body
// @Tags         test
// @Produce      json
// @Param        test body model.TestRequest true "Payload"
// @Success      200  {object} model.IdResponse
// @Failure     400  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test [post]
func AddTestHandle(ctx *gin.Context) {
	var request model.TestRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId, _ := ctx.Get("user")
	createdBy := userId.(string)
	date := time.Now()
	id, _ := uuid.NewV4()
	Test := &model.Test{
		Id:        id,
		Name:      request.Name,
		CreatedBy: createdBy,
		CourseId:  request.CourseId,
		CreatedAt: date,
		ChangedBy: nil,
	}

	err = db.AddTestToDB(Test)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

// GetTests            godoc
// @Summary      Get tests
// @Description  Get all tests
// @Tags         test
// @Produce      json
// @Success      200  {array}  model.Test
// @Failure     500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test [get]
func GetTestsHandle(ctx *gin.Context) {
	Tests, err := db.GetTestsFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, Tests)
}

// GetTest            godoc
// @Summary      Get test
// @Description  Get test by id
// @Tags         test
// @Produce      json
// @Param        id  path  string  true  "Test ID"
// @Success      200  {object} model.Test
// @Failure    404  {object} model.ErrorResponse
// @Failure    500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/{id} [get]
func GetTestHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	Test, err := db.GetTestFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, Test)
}

// UpdateTest            godoc
// @Summary      Update test
// @Description  Update test by id
// @Tags         test
// @Produce      json
// @Param        id  path  string  true  "Test ID"
// @Param        updatedTest body model.TestRequest true "Payload"
// @Success      200  {object} model.BaseResponse
// @Failure   404  {object} model.ErrorResponse
// @Failure   500  {object} model.ErrorResponse
// @Failure   400  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/{id} [put]
func UpdateTestHandle(ctx *gin.Context) {
	var request model.TestRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	userId, _ := ctx.Get("user")
	changedBy := userId.(string)
	changedAt := time.Now()
	Test := &model.Test{
		Id:        id,
		Name:      request.Name,
		ChangedBy: &changedBy,
		CourseId:  request.CourseId,
		ChangedAt: &changedAt,
	}

	err = db.UpdateTestInDB(Test)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "OK"})
}

// DeleteTest            godoc
// @Summary      Delete test
// @Description  Delete test by id
// @Tags         test
// @Produce      json
// @Param        id  path  string  true  "Test ID"
// @Success      200  {object} model.BaseResponse
// @Failure  404  {object} model.ErrorResponse
// @Failure  500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/{id} [delete]
func DeleteTestHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteTestFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddTestHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/test", RequireAuth)
	subGroup.POST("", AddTestHandle)
	subGroup.GET("", GetTestsHandle)
	subGroup.GET(":id", GetTestHandle)
	subGroup.PUT(":id", UpdateTestHandle)
	subGroup.DELETE(":id", DeleteTestHandle)
}
