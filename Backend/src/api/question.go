package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
)

type QuestionRequest struct {
	Body    string    `json:"body"`
	ImgFile string    `json:"img_file"`
	TestId  uuid.UUID `json:"test_id"`
}

func AddQuestionHandle(ctx *gin.Context) {
	var request QuestionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.NewV4()
	Question := &db.Question{
		Id:      id,
		Body:    request.Body,
		ImgFile: request.ImgFile,
		TestId:  request.TestId,
	}

	err = db.AddQuestionToDB(Question)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func GetQuestionsHandle(ctx *gin.Context) {
	questions, err := db.GetQuestionsFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, questions)
}

func GetQuestionHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	question, err := db.GetQuestionFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, question)
}

func UpdateQuestionHandle(ctx *gin.Context) {
	var request QuestionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"Json decode error: ": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	question := &db.Question{
		Id:      id,
		Body:    request.Body,
		ImgFile: request.ImgFile,
		TestId:  request.TestId,
	}

	err = db.UpdateQuestionInDB(question)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func DeleteQuestionHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteQuestionFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddQuestionHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/question")
	subGroup.POST("", AddQuestionHandle)
	subGroup.GET("", GetQuestionsHandle)
	subGroup.GET(":id", GetQuestionHandle)
	subGroup.PUT(":id", UpdateQuestionHandle)
	subGroup.DELETE(":id", DeleteQuestionHandle)
}
