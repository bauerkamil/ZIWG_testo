package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
	"src/model"
)

// AddTeacher            godoc
// @Summary      Add teacher
// @Description  Add teacher from json body
// @Tags         teacher
// @Produce      json
// @Success      200  {array}  model.TeacherRequest
// @Router       /api/v1/teacher [get]
func AddTeacherHandle(ctx *gin.Context) {
	var request model.TeacherRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.NewV4()
	teacher := &model.Teacher{
		Id:         id,
		Name:       request.Name,
		SecondName: request.SecondName,
		Surname:    request.Surname,
	}

	err = db.AddTeacherToDB(teacher)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func GetTeachersHandle(ctx *gin.Context) {
	teachers, err := db.GetTeachersFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, teachers)
}

func GetTeacherHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	teacher, err := db.GetTeacherFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, teacher)
}

func UpdateTeacherHandle(ctx *gin.Context) {
	var request model.TeacherRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"Json decode error: ": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	teacher := &model.Teacher{
		Id:         id,
		Name:       request.Name,
		SecondName: request.SecondName,
		Surname:    request.Surname,
	}

	err = db.UpdateTeacherInDB(teacher)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func DeleteTeacherHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteTeacherFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddTeacherHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/teacher")
	subGroup.POST("", AddTeacherHandle)
	subGroup.GET("", GetTeachersHandle)
	subGroup.GET(":id", GetTeacherHandle)
	subGroup.PUT(":id", UpdateTeacherHandle)
	subGroup.DELETE(":id", DeleteTeacherHandle)
}
