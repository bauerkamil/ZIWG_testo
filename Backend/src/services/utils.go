package services

import (
	"archive/zip"
	"errors"
	"github.com/gofrs/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"src/model"
	"src/model/dto"
	"strings"
)

func createNewQuestion(question dto.SubQuestion, testId uuid.UUID) model.Question {
	questionId, _ := uuid.NewV4()
	return model.Question{
		Id:      questionId,
		Body:    question.Body,
		ImgFile: question.ImgFile,
		TestId:  testId,
	}
}

func createNewAnswer(answer dto.SubAnswer, questionId uuid.UUID) model.Answer {
	answerId, _ := uuid.NewV4()
	return model.Answer{
		Id:         answerId,
		Body:       answer.Body,
		Valid:      answer.Valid,
		QuestionId: questionId,
	}
}

func containsAnswer(answers []model.Answer, answer model.Answer) bool {
	for _, a := range answers {
		if a.Id == answer.Id {
			return true
		}
	}
	return false
}

func UnzipArchive(file multipart.File, header *multipart.FileHeader) (*string, error) {
	// Create a temporary directory to unzip the archive
	tempDir := "C://users//kacpe//desktop"

	// Create a new zip reader from the file
	zipReader, err := zip.NewReader(file, header.Size)
	if err != nil {
		return nil, err
	}

	// Iterate through each file in the zip archive
	for _, zipFile := range zipReader.File {
		destPath := filepath.Join(tempDir, zipFile.Name)

		// Check for ZipSlip (Directory traversal attack)
		if !strings.HasPrefix(destPath, filepath.Clean(tempDir)+string(os.PathSeparator)) {
			return nil, errors.New("illegal file path")
		}

		if zipFile.FileInfo().IsDir() {
			os.MkdirAll(destPath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
				return nil, err
			}

			outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
			if err != nil {
				return nil, err
			}

			rc, err := zipFile.Open()
			if err != nil {
				return nil, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()
			rc.Close()

			if err != nil {
				return nil, err
			}
		}
	}
	return &tempDir, nil
}
