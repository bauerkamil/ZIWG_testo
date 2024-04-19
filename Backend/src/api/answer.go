package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
	"src/model"
)

func AddAnswerHandle(ctx *gin.Context) {
	var request model.AnswerRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.NewV4()
	answer := &model.Answer{
		Id:         id,
		Body:       request.Body,
		QuestionId: request.QuestionId,
		Valid:      request.Valid,
	}
	err = db.AddAnswerToDB(answer)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func GetAnswersHandle(ctx *gin.Context) {
	answers, err := db.GetAnswersFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, answers)
}

func GetAnswerHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	answer, err := db.GetAnswerFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, answer)
}

func UpdateAnswerHandle(ctx *gin.Context) {
	var request model.AnswerRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"Json decode error: ": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	answer := &model.Answer{
		Id:         id,
		Body:       request.Body,
		Valid:      request.Valid,
		QuestionId: request.QuestionId,
	}
	err = db.UpdateAnswerInDB(answer)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func DeleteAnswerHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteAnswerFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddAnswerHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/answer")
	subGroup.POST("", AddAnswerHandle)
	subGroup.GET("", GetAnswersHandle)
	subGroup.GET(":id", GetAnswerHandle)
	subGroup.PUT(":id", UpdateAnswerHandle)
	subGroup.DELETE(":id", DeleteAnswerHandle)
}
