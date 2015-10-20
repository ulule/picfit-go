package picfit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGeometry(t *testing.T) {
	assert := assert.New(t)

	geometry, err := ParseGeometry("22x32")
	assert.Nil(err)

	assert.Equal(geometry.X, 22)
	assert.Equal(geometry.Y, 32)

	geometry, err = ParseGeometry("x")
	assert.NotNil(err)

	geometry, err = ParseGeometry("x32")
	assert.Nil(err)

	geometry, err = ParseGeometryWithRatio("20x", float64(2.0))
	assert.Nil(err)
	assert.Equal(geometry.X, 20)
	assert.Equal(geometry.Y, 10)

	geometry, err = ParseGeometryWithRatio("x20", float64(2.0))
	assert.Nil(err)
	assert.Equal(geometry.X, 40)
	assert.Equal(geometry.Y, 20)
}
