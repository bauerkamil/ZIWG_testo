package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
	"src/model"
)

// AddAnswer            godoc
// @Summary      Add answer
// @Description  Add answer from json body
// @Tags         answer
// @Produce      json
// @Param        answer body model.AnswerRequest true "Payload"
// @Success      200  {object} model.IdResponse
// @Failure     400  {object} model.ErrorResponse
// @Failure     500  {object} model.ErrorResponse
// @Router       /api/v1/answer [post]
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

	ctx.JSON(200, gin.H{"id": id})
}

// GetAnswers            godoc
// @Summary      Get answers
// @Description  Get all answers
// @Tags         answer
// @Produce      json
// @Success      200  {array}  model.Answer
// @Failure     500  {object} model.ErrorResponse
// @Router       /api/v1/answer [get]
func GetAnswersHandle(ctx *gin.Context) {
	answers, err := db.GetAnswersFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, answers)
}

// GetAnswer            godoc
// @Summary      Get answer
// @Description  Get answer by id
// @Tags         answer
// @Produce      json
// @Param        id path string true "Answer ID"
// @Success      200  {object} model.Answer
// @Failure    404  {object} model.ErrorResponse
// @Failure    500  {object} model.ErrorResponse
// @Router       /api/v1/answer/{id} [get]
func GetAnswerHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	answer, err := db.GetAnswerFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, answer)
}

// UpdateAnswer            godoc
// @Summary      Update answer
// @Description  Update answer by id
// @Tags         answer
// @Produce      json
// @Param        id  path  string  true  "Answer ID"
// @Param        updatedAnswer body model.AnswerRequest true "Payload"
// @Success      200  {object} model.BaseResponse
// @Failure    404  {object} model.ErrorResponse
// @Failure    500  {object} model.ErrorResponse
// @Failure    400  {object} model.ErrorResponse
// @Router       /api/v1/answer/{id} [put]
func UpdateAnswerHandle(ctx *gin.Context) {
	var request model.AnswerRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
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
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

// DeleteAnswer            godoc
// @Summary      Delete answer
// @Description  Delete answer by id
// @Tags         answer
// @Produce      json
// @Param        id  path  string  true  "Answer ID"
// @Success      200  {object} model.BaseResponse
// @Failure   404  {object} model.ErrorResponse
// @Failure   500  {object} model.ErrorResponse
// @Router       /api/v1/answer/{id} [delete]
func DeleteAnswerHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteAnswerFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
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
