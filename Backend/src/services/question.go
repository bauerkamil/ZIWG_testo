package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"net/http"
	"src/dal"
	"src/model"
	"src/model/dto"
)

// AddQuestion            godoc
// @Summary      Add question
// @Description  Add question from json body
// @Tags         question
// @Produce      json
// @Param        question body dto.QuestionRequest true "Payload"
// @Success      200  {object} dto.BaseResponse
// @Failure     400  {object} dto.ErrorResponse
// @Failure     500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question [post]
func AddQuestionHandle(ctx *gin.Context) {
	var request dto.QuestionRequest
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

	err = dal.AddQuestionToDB(Question)
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
// @Failure     500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question [get]
func GetQuestionsHandle(ctx *gin.Context) {
	questions, err := dal.GetQuestionsFromDB()
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
// @Failure     404  {object} dto.ErrorResponse
// @Failure     500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id} [get]
func GetQuestionHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	question, err := dal.GetQuestionFromDB(id)
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
// @Param        updatedQuestion body dto.QuestionRequest true "Payload"
// @Success      200  {object} dto.IdResponse
// @Failure    404  {object} dto.ErrorResponse
// @Failure    500  {object} dto.ErrorResponse
// @Failure    400  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id} [put]
func UpdateQuestionHandle(ctx *gin.Context) {
	var request dto.QuestionRequest
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

	err = dal.UpdateQuestionInDB(question)
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
// @Success      200  {object} dto.BaseResponse
// @Failure  404  {object} dto.ErrorResponse
// @Failure  500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id} [delete]
func DeleteQuestionHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = dal.DeleteQuestionFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

// AddImage            godoc
// @Summary      Add image to question
// @Description  Add image to question by id
// @Tags         image
// @Produce      json
// @Param        id  path  string  true  "Question ID"
// @Param			file formData file true "file"
// @Success      200  {object} dto.BaseResponse
// @Failure  404  {object} dto.ErrorResponse
// @Failure  500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/question/{id}/image [post]
func AddImageHandle(ctx *gin.Context) {

	azureProvider, err := GetAzureProviderInstance()
	if err != nil {
		fmt.Printf("Failed to create AzureProvider: %v\n", err)
		return
	}

	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = azureProvider.UploadFile(ctx, id)
	if err != nil {
		var ginErr *gin.Error
		if errors.As(err, &ginErr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func AddQuestionHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/question", RequireAuth)
	subGroup.POST("", AddQuestionHandle)
	subGroup.GET("", GetQuestionsHandle)
	subGroup.GET(":id", GetQuestionHandle)
	subGroup.PUT(":id", UpdateQuestionHandle)
	subGroup.POST(":id/image", AddImageHandle)
	subGroup.DELETE(":id", DeleteQuestionHandle)
}
