package matrix

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("matrix", func() {

	var (
		records = [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}
	)

	Context("stringify", func() {

		It("should populate and get matrix string representation", func() {
			// Arrange
			m := NewMatrix[int]()
			err := m.Populate(records)

			// Act
			stringified := m.String()

			// Assert
			Expect(err).To(BeNil())
			Expect(m).NotTo(BeNil())
			Expect(stringified).To(Equal("1,2,3\n4,5,6\n7,8,9\n"))
		})
	})

	Context("invert", func() {

		It("should populate and get inverted matrix", func() {
			// Arrange
			m := NewMatrix[int]()
			err := m.Populate(records)

			// Act
			inverted := m.Invert()

			// Assert
			Expect(err).To(BeNil())
			Expect(inverted).NotTo(BeNil())
			Expect(inverted.String()).To(Equal("1,4,7\n2,5,8\n3,6,9\n"))
		})
	})

	Context("flatten", func() {

		It("should populate and get flat matrix", func() {
			// Arrange
			m := NewMatrix[int]()
			err := m.Populate(records)

			// Act
			flat := m.Flatten()

			// Assert
			Expect(err).To(BeNil())
			Expect(flat).NotTo(BeNil())
			Expect(flat).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})

	Context("sum", func() {

		It("should populate and get matrix sum", func() {
			// Arrange
			m := NewMatrix[int]()
			err := m.Populate(records)

			// Act
			sum := m.Sum()

			// Assert
			Expect(err).To(BeNil())
			Expect(sum).To(Equal(45))
		})
	})

	Context("multiply", func() {

		It("should populate and get matrix multiply", func() {
			// Arrange
			m := NewMatrix[int]()
			err := m.Populate(records)

			// Act
			multiply := m.Multiply()

			// Assert
			Expect(err).To(BeNil())
			Expect(multiply).To(Equal(362880))
		})
	})
})
