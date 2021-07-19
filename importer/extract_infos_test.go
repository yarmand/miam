package importer

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestExtractInfo(t *testing.T) {
	t.Run("GetCreationDate", func(t *testing.T) {
		assert := assert.New(t)
		fname := "../testAssets/testImage.jpg"
		date, _ := GetCreationDate(afero.NewOsFs(), fname)
		assert.Equal(date.Year(), 2018)
		assert.Equal(int(date.Month()), 6)
		assert.Equal(date.Day(), 25)
	})
}
