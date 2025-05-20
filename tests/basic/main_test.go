package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("Add One(%d), Input = %d , actual = %d ", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be equal be 3")
}
