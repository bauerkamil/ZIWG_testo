package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
)

type CourseRequest struct {
	Name       string    `json:"name"`
	Teacher    uuid.UUID `json:"teacher"`
	SchoolYear int       `json:"school_year"`
}

func AddCourseHandle(ctx *gin.Context) {
	var request CourseRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.NewV4()
	course := &db.Course{
		Id:         id,
		Name:       request.Name,
		Teacher:    request.Teacher,
		SchoolYear: request.SchoolYear,
	}

	err = db.AddCourseToDB(course)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func GetCoursesHandle(ctx *gin.Context) {
	courses, err := db.GetCoursesFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, courses)
}

func GetCourseHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	course, err := db.GetCourseFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, course)
}

func UpdateCourseHandle(ctx *gin.Context) {
	var request CourseRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"Json decode error: ": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	Course := &db.Course{
		Id:         id,
		Name:       request.Name,
		Teacher:    request.Teacher,
		SchoolYear: request.SchoolYear,
	}

	err = db.UpdateCourseInDB(Course)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

func DeleteCourseHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteCourseFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})

}

func AddCourseHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/course")
	subGroup.POST("", AddCourseHandle)
	subGroup.GET("", GetCoursesHandle)
	subGroup.GET(":id", GetCourseHandle)
	subGroup.PUT(":id", UpdateCourseHandle)
	subGroup.DELETE(":id", DeleteCourseHandle)
}
