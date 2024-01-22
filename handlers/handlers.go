package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/condur/matrix/log"
	"github.com/condur/matrix/matrix"
	"github.com/condur/matrix/param"
	"github.com/condur/matrix/types"
)

func Echo(c *gin.Context) {
	// Parse request as CSV
	records, err := param.BindCSV(c)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Initialize matrix
	m := matrix.NewMatrix[float64]()

	// Populate matrix with parsed CSV records
	err = m.Populate(records)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get matrix string representation
	response := m.String()

	// Return response
	c.String(http.StatusOK, response)
}

func Invert(c *gin.Context) {
	// Parse request as CSV
	records, err := param.BindCSV(c)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Initialize matrix
	m := matrix.NewMatrix[float64]()

	// Populate matrix with parsed CSV records
	err = m.Populate(records)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get matrix invert representation
	inverted := m.Invert()

	// Get inverted matrix string representation
	response := inverted.String()

	// Return response
	c.String(http.StatusOK, response)
}

func Flatten(c *gin.Context) {
	// Parse request as CSV
	records, err := param.BindCSV(c)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Initialize matrix
	m := matrix.NewMatrix[float64]()

	// Populate matrix with parsed CSV records
	err = m.Populate(records)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get matrix flatten string representation
	response := m.FlattenString(matrix.Delimiter)

	// Return response
	c.String(http.StatusOK, response)
}

func Sum(c *gin.Context) {
	// Parse request as CSV
	records, err := param.BindCSV(c)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Initialize matrix
	m := matrix.NewMatrix[float64]()

	// Populate matrix with parsed CSV records
	err = m.Populate(records)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get matrix sum
	sum := m.Sum()

	// Get response
	response := types.ToString(sum)

	// Return response
	c.String(http.StatusOK, response)
}

func Multiply(c *gin.Context) {
	// Parse request as CSV
	records, err := param.BindCSV(c)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	// Initialize matrix
	m := matrix.NewMatrix[float64]()

	// Populate matrix with parsed CSV records
	err = m.Populate(records)
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get matrix multiplication
	multiply := m.Multiply()

	// Get response
	response := types.ToString(multiply)

	// Return response
	c.String(http.StatusOK, response)
}
