package param

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/condur/matrix/validate"
)

type Param int

const (
	Json Param = iota
	Query
	URI
	Form
)

func Bind[T any](c *gin.Context, param Param, obj T) (T, error) {
	// Define error variable
	var err error

	// Bind
	switch param {
	case Json:
		err = c.ShouldBindJSON(obj)
	case Query:
		err = c.ShouldBindQuery(obj)
	case URI:
		err = c.ShouldBindUri(obj)
	case Form:
		err = c.ShouldBindWith(obj, binding.Form)
	}

	// Handle binding error
	if err != nil {
		return obj, err
	}

	// Validate parameters
	err = validate.Struct(c, obj)

	return obj, err
}

func BindJson[T any](c *gin.Context, obj T) (T, error) {
	return Bind(c, Json, obj)
}

func BindQuery[T any](c *gin.Context, obj T) (T, error) {
	return Bind(c, Query, obj)
}

func BindUri[T any](c *gin.Context, obj T) (T, error) {
	return Bind(c, URI, obj)
}

func BindForm[T any](c *gin.Context, obj T) (T, error) {
	return Bind(c, Form, obj)
}

func BindCSV(c *gin.Context) ([][]string, error) {
	// Get file pointer
	file_ptr, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	// Open file
	file, err := file_ptr.Open()
	if err != nil {
		return nil, err
	}

	// Close the file on exit
	defer file.Close()

	// Parse CSV
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
