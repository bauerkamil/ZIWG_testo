package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"src/db"
	"src/model"
)

// AddCourse            godoc
// @Summary      Add course
// @Description  Add course from json body
// @Tags         course
// @Produce      json
// @Param        course body model.CourseRequest true "Payload"
// @Success      200  {object} model.IdResponse
// @Failure     400  {object} model.ErrorResponse
// @Failure     500  {object} model.ErrorResponse
// @Router       /api/v1/course [post]
func AddCourseHandle(ctx *gin.Context) {
	var request model.CourseRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.NewV4()
	course := &model.Course{
		Id:         id,
		Name:       request.Name,
		TeacherId:  request.TeacherId,
		SchoolYear: request.SchoolYear,
	}

	err = db.AddCourseToDB(course)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

// GetCourses            godoc
// @Summary      Get courses
// @Description  Get all courses
// @Tags         course
// @Produce      json
// @Success      200  {array}  model.Course
// @Failure     500  {object} model.ErrorResponse
// @Router       /api/v1/course [get]
func GetCoursesHandle(ctx *gin.Context) {
	courses, err := db.GetCoursesFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, courses)
}

// GetCourse            godoc
// @Summary      Get course
// @Description  Get course by id
// @Tags         course
// @Produce      json
// @Param        id path string true "Course ID"
// @Success      200  {object} model.Course
// @Failure    404  {object} model.ErrorResponse
// @Failure    500  {object} model.ErrorResponse
// @Router       /api/v1/course/{id} [get]
func GetCourseHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	course, err := db.GetCourseFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, course)
}

// UpdateCourse            godoc
// @Summary      Update course
// @Description  Update course by id
// @Tags         course
// @Produce      json
// @Param        id  path  string  true  "Course ID"
// @Param        updatedCourse body model.CourseRequest true "Payload"
// @Success      200  {object} model.BaseResponse
// @Failure    404  {object} model.ErrorResponse
// @Failure    500  {object} model.ErrorResponse
// @Failure    400  {object} model.ErrorResponse
// @Router       /api/v1/course/{id} [put]
func UpdateCourseHandle(ctx *gin.Context) {
	var request model.CourseRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.FromString(ctx.Param("id"))
	Course := &model.Course{
		Id:         id,
		Name:       request.Name,
		TeacherId:  request.TeacherId,
		SchoolYear: request.SchoolYear,
	}

	err = db.UpdateCourseInDB(Course)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

// DeleteCourse            godoc
// @Summary      Delete course
// @Description  Delete course by id
// @Tags         course
// @Produce      json
// @Param        id  path  string  true  "Course ID"
// @Success      200  {object} model.BaseResponse
// @Failure   404  {object} model.ErrorResponse
// @Failure   500  {object} model.ErrorResponse
// @Router       /api/v1/course/{id} [delete]
func DeleteCourseHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = db.DeleteCourseFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
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
