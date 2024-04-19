package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"src/db"
	"src/model"
)

func AddTestHandle(ctx *gin.Context) {
	var request model.TestRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	date := *timestamppb.Now()
	id, _ := uuid.NewV4()
	Test := &model.Test{
		Id:        id,
		Name:      request.Name,
		CreatedBy: request.User,
		CourseId:  request.CourseId,
		CreatedAt: date,
		ChangedBy: "",
	}

	err = db.AddTestToDB(Test)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func GetTestsHandle(ctx *gin.Context) {
	Tests, err := db.GetTestsFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, Tests)
}

func GetTestHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	Test, err := db.GetTestFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, Test)
}

func UpdateTestHandle(ctx *gin.Context) {
	var request model.TestRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"Json decode error: ": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	Test := &model.Test{
		Id:        id,
		Name:      request.Name,
		ChangedBy: request.User,
		CourseId:  request.CourseId,
	}

	err = db.UpdateTestInDB(Test)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func DeleteTestHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteTestFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddTestHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/test")
	subGroup.POST("", AddTestHandle)
	subGroup.GET("", GetTestsHandle)
	subGroup.GET(":id", GetTestHandle)
	subGroup.PUT(":id", UpdateTestHandle)
	subGroup.DELETE(":id", DeleteTestHandle)
}
