package importer

import (
	"fmt"
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
)

func TestProcessorFunctions(t *testing.T) {
	t.Run("WrapIn use result of left processor function as input to the wrap function", func(t *testing.T) {
		var result string
		i := New(func(c context.Context, i Importer, f string) string {
			return "bob"
		}, "", "", nil)
		i2 := i.WrapedIn(func(c context.Context, i Importer, f string) string {
			result = "wrapped_" + f
			return result
		})
		i2.ImportFiles([]string{"hello"})
		assert.Equal(t, "wrapped_bob", result, "failed wrapping other Importer")
	})

	t.Run("After run the parameter funtion after the left processor function", func(t *testing.T) {
		var result string
		i := New(func(c context.Context, i Importer, f string) string {
			result = "before_" + f
			return result
		}, "", "", nil)
		i2 := i.After(func(c context.Context, i Importer, f string) string {
			return "bob"
		})
		i2.ImportFiles([]string{"hello"})
		assert.Equal(t, "before_bob", result, "failed wrapping other Importer")
	})
}

func TestImportFiles(t *testing.T) {
	t.Run("ImportFiles apply the processor function to all imported files", func(t *testing.T) {
		var processedFiles []string
		files := []string{"a", "b", "c"}
		process := func(ctx context.Context, i Importer, f string) string {
			processedFiles = append(processedFiles, f)
			return ""
		}
		i := New(process, "", "", nil)
		i.ImportFiles(files)
		assert.Equal(t, files, processedFiles, fmt.Sprintf("processedFiles do not equal input files\n processFiles=%v\ninputFiles=%v", processedFiles, files))
	})
}
