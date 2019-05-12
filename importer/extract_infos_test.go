package importer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractInfo(t *testing.T) {
	t.Run("GetCreationDate", func(t *testing.T) {
		assert := assert.New(t)
		fname := "../testAssets/testImage.jpeg"
		date := GetCreationDate(fname)
		assert.Equal(date.Year(), 2018)
		assert.Equal(int(date.Month()), 6)
		assert.Equal(date.Day(), 25)
	})
}
