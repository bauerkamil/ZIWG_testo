package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
	"src/model"
)

// AddQuestion            godoc
// @Summary      Add question
// @Description  Add question from json body
// @Tags         question
// @Produce      json
// @Param        question body model.QuestionRequest true "Payload"
// @Success      200  {object} model.BaseResponse
// @Failure     400  {object} model.ErrorResponse
// @Failure     500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question [post]
func AddQuestionHandle(ctx *gin.Context) {
	var request model.QuestionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.NewV4()
	Question := &model.Question{
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

// GetQuestions            godoc
// @Summary      Get questions
// @Description  Get all questions
// @Tags         question
// @Produce      json
// @Success      200  {array}  model.Question
// @Failure     500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question [get]
func GetQuestionsHandle(ctx *gin.Context) {
	questions, err := db.GetQuestionsFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, questions)
}

// GetQuestion            godoc
// @Summary      Get question
// @Description  Get question by id
// @Tags         question
// @Produce      json
// @Param        id path string true "Question ID"
// @Success      200  {object} model.Question
// @Failure     404  {object} model.ErrorResponse
// @Failure     500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id} [get]
func GetQuestionHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	question, err := db.GetQuestionFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, question)
}

// UpdateQuestion            godoc
// @Summary      Update question
// @Description  Update question by id
// @Tags         question
// @Produce      json
// @Param        id  path  string  true  "Question ID"
// @Param        updatedQuestion body model.QuestionRequest true "Payload"
// @Success      200  {object} model.IdResponse
// @Failure    404  {object} model.ErrorResponse
// @Failure    500  {object} model.ErrorResponse
// @Failure    400  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id} [put]
func UpdateQuestionHandle(ctx *gin.Context) {
	var request model.QuestionRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	question := &model.Question{
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
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

// DeleteQuestion            godoc
// @Summary      Delete question
// @Description  Delete question by id
// @Tags         question
// @Produce      json
// @Param        id  path  string  true  "Question ID"
// @Success      200  {object} model.BaseResponse
// @Failure  404  {object} model.ErrorResponse
// @Failure  500  {object} model.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id} [delete]
func DeleteQuestionHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteQuestionFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddQuestionHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/question", RequireAuth)
	subGroup.POST("", AddQuestionHandle)
	subGroup.GET("", GetQuestionsHandle)
	subGroup.GET(":id", GetQuestionHandle)
	subGroup.PUT(":id", UpdateQuestionHandle)
	subGroup.DELETE(":id", DeleteQuestionHandle)
}
