package services

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"golang.org/x/text/encoding/charmap"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"path/filepath"
	"src/dal"
	"src/model"
	"src/model/dto"
	"strings"
	"time"
)

// AddTestHandle            godoc
// @Summary      Add test
// @Description  Add test from json body
// @Tags         test
// @Produce      json
// @Param        test body dto.TestRequest true "Payload"
// @Success      200  {object} dto.IdResponse
// @Failure     400  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test [post]
func AddTestHandle(ctx *gin.Context) {
	var request dto.TestRequest
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
		Id:         id,
		Name:       request.Name,
		CreatedBy:  createdBy,
		CourseId:   request.CourseId,
		CreatedAt:  date,
		ChangedBy:  nil,
		SchoolYear: request.SchoolYear,
	}
	err = dal.AddTestToDB(Test)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

// GetTestsHandle            godoc
// @Summary      Get tests
// @Description  Get all tests
// @Tags         test
// @Produce      json
// @Success      200  {array}  dto.ListTest
// @Failure     500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test [get]
func GetTestsHandle(ctx *gin.Context) {
	tests, err := dal.GetTestsFromDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//loop over tests and convert to listTest
	output := make([]dto.ListTest, len(tests))
	for i, test := range tests {
		output[i] = dto.ToListTest(test)
	}
	ctx.JSON(200, output)
}

// GetActiveTestsHandle            godoc
// @Summary      Get active tests
// @Description  Get all user active tests
// @Tags         test
// @Produce      json
// @Success      200  {array}  dto.ListTest
// @Failure     500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/active [get]
func GetActiveTestsHandle(ctx *gin.Context) {
	_, activeCourses := GetActiveUserCourses(ctx)
	tests, err := dal.GetTestsById(activeCourses)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//loop over tests and convert to listTest
	output := make([]dto.ListTest, len(tests))
	for i, test := range tests {
		output[i] = dto.ToListTest(test)
	}
	ctx.JSON(200, output)
}

// GetTestHandle            godoc
// @Summary      Get test
// @Description  Get test by id
// @Tags         test
// @Produce      json
// @Param        id  path  string  true  "Test ID"
// @Success      200  {object} dto.FullTest
// @Failure    404  {object} dto.ErrorResponse
// @Failure    500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/{id} [get]
func GetTestHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	test, err := dal.GetTestFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, dto.ToFullTest(*test))
}

// UpdateTestHandle            godoc
// @Summary      Update test
// @Description  Update test by id
// @Tags         test
// @Produce      json
// @Param        id  path  string  true  "Test ID"
// @Param        updatedTest body dto.EditTestRequest true "Payload"
// @Success      200  {object} dto.BaseResponse
// @Failure   404  {object} dto.ErrorResponse
// @Failure   500  {object} dto.ErrorResponse
// @Failure   400  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/{id} [put]
func UpdateTestHandle(ctx *gin.Context) {
	var request dto.EditTestRequest
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
		Id:         id,
		Name:       request.Name,
		ChangedBy:  &changedBy,
		CourseId:   request.CourseId,
		ChangedAt:  &changedAt,
		SchoolYear: request.SchoolYear,
	}

	err = dal.UpdateTestInDB(Test)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "OK"})
}

// DeleteTestHandle            godoc
// @Summary      Delete test
// @Description  Delete test by id
// @Tags         test
// @Produce      json
// @Param        id  path  string  true  "Test ID"
// @Success      200  {object} dto.BaseResponse
// @Failure  404  {object} dto.ErrorResponse
// @Failure  500  {object} dto.ErrorResponse
// @Security     BearerAuth
// @Router       /api/v1/test/{id} [delete]
func DeleteTestHandle(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	err = dal.DeleteTestFromDB(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{"Record not found with id": id})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OK"})
}

// ImportTestHandle            godoc
// @Summary      Import test
// @Description  Import test from zip file
// @Tags         test
// @Produce      json
// @Success      200  {object} dto.IdResponse
// @Failure     400  {object} dto.ErrorResponse
// @Failure     500  {object} dto.ErrorResponse
// @Param			file formData file true "file"
// @Security     BearerAuth
// @Router       /api/v1/test/import [post]
func ImportTestHandle(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		}
	}(file)

	if header.Size > 10<<20 {
		ctx.JSON(400, gin.H{"error": "File size too big"})
		return
	}

	fileNameParts := strings.Split(header.Filename, ".")
	if fileNameParts[len(fileNameParts)-1] != "zip" {
		ctx.JSON(400, gin.H{"error": "File must be a zip archive"})
		return
	}

	archiveDest, err := UnzipArchive(file, header)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	*archiveDest += "//" + fileNameParts[0]

	// Read all .txt and .png files from the unzipped directory
	err = filepath.Walk(*archiveDest, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if ext == ".png" {
				//data, err := os.ReadFile(path)
				if err != nil {
					return err
				}
			} else if ext == ".txt" {
				file, err = os.Open(path)
				reader := charmap.Windows1250.NewDecoder().Reader(file)
				scanner := bufio.NewScanner(reader)

				for scanner.Scan() {
					fmt.Println(scanner.Text())
				}
			}
		}

		return nil
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func AddTestHandlers(router *gin.RouterGroup) {
	var subGroup = router.Group("/test", RequireAuth)
	subGroup.POST("", AddTestHandle)
	subGroup.GET("", GetTestsHandle)
	subGroup.GET(":id", GetTestHandle)
	subGroup.PUT(":id", UpdateTestHandle)
	subGroup.DELETE(":id", DeleteTestHandle)
	subGroup.GET("active", GetActiveTestsHandle)
	subGroup.POST("import", ImportTestHandle)
}
