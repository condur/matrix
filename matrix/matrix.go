package matrix

import (
	"fmt"
	"strings"
	"sync"

	"github.com/condur/matrix/types"
)

const (
	Delimiter string = ","
)

type Matrix[T types.Number] struct {
	lock      sync.RWMutex // Locking
	container [][]T        // Matrix
}

func NewMatrix[T types.Number]() *Matrix[T] {
	return &Matrix[T]{
		lock:      sync.RWMutex{},
		container: make([][]T, 0),
	}
}

func (m *Matrix[T]) Populate(records [][]string) error {
	// Lock
	m.lock.Lock()
	defer m.lock.Unlock()

	// Get the records size
	size := len(records)

	// Initialize matrix
	m.container = make([][]T, size)
	for i := 0; i < size; i++ {
		m.container[i] = make([]T, len(records[i]))
	}

	// Define reusable variables
	var val T
	var err error

	// Loop over records
	for i := 0; i < size; i++ {
		for j := 0; j < len(records[i]); j++ {

			// Convert record value
			val, err = types.Parse[T](records[i][j])
			if err != nil {
				return err
			}

			// Assign value
			m.container[i][j] = val
		}
	}

	return nil
}

func (m *Matrix[T]) Invert() *Matrix[T] {
	// Lock
	m.lock.Lock()
	defer m.lock.Unlock()

	// Get the matrix size
	size := len(m.container)

	// Initialize inverted matrix
	inverted := make([][]T, size)
	for i := 0; i < size; i++ {
		inverted[i] = make([]T, len(m.container[i]))
	}

	// Loop and assign
	for i := 0; i < size; i++ {
		for j := 0; j < len(m.container[i]); j++ {
			inverted[i][j] = m.container[j][i]
		}
	}

	// Create a new matrix from inverted representation
	return &Matrix[T]{
		lock:      sync.RWMutex{},
		container: inverted,
	}
}

func (m *Matrix[T]) Flatten() []T {
	// Lock
	m.lock.Lock()
	defer m.lock.Unlock()

	// Initialize flatten slice
	var flatten []T

	// Loop over matrix values
	for _, row := range m.container {
		flatten = append(flatten, row...)
	}

	return flatten
}

func (m *Matrix[T]) FlattenString(delimiter string) string {
	// Convert matrix to flat version
	flatten := m.Flatten()

	// Initialize flatten result slice
	flat := make([]string, len(flatten))

	// Loop over flatten values
	for i := range flatten {
		flat[i] = types.ToString(flatten[i])
	}

	// Get the string representation
	return strings.Join(flat, delimiter)
}

func (m *Matrix[T]) Sum() T {
	// Lock
	m.lock.Lock()
	defer m.lock.Unlock()

	// Define sum variable
	var sum T

	// Loop over matrix value
	for i := 0; i < len(m.container); i++ {
		for j := 0; j < len(m.container[i]); j++ {
			if i == 0 && j == 0 {
				sum = m.container[i][j]
			} else {
				sum += m.container[i][j]
			}
		}
	}

	return sum
}

func (m *Matrix[T]) Multiply() T {
	// Lock
	m.lock.Lock()
	defer m.lock.Unlock()

	// Define multiply variable
	var multiply T

	// Loop over matrix value
	for i := 0; i < len(m.container); i++ {
		for j := 0; j < len(m.container[i]); j++ {
			if i == 0 && j == 0 {
				multiply = m.container[i][j]
			} else {
				multiply *= m.container[i][j]
			}

		}
	}

	return multiply
}

func (m *Matrix[T]) String() string {
	// Lock
	m.lock.Lock()
	defer m.lock.Unlock()

	// Get the matrix size
	size := len(m.container)

	// Initialize strings based matrix
	converted := make([][]string, size)
	for i := 0; i < size; i++ {
		converted[i] = make([]string, len(m.container[i]))
	}

	// Convert matrix values to string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			converted[i][j] = types.ToString(m.container[i][j])
		}
	}

	// Prepare response
	var response string
	for _, row := range converted {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, Delimiter))
	}

	return response
}
