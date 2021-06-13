package importer

import (
	"fmt"
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
)

func TestWrapedIn(t *testing.T) {
	var result string
	i := New(func(c context.Context, f string) string {
		return "bob"
	})
	i2 := i.WrapedIn(func(c context.Context, f string) string {
		result = "wrapped_" + f
		return result
	})
	i2.ImportFiles([]string{"hello"})
	assert.Equal(t, "wrapped_bob", result, "failed wrapping other Importer")
}

func TestAfter(t *testing.T) {
	var result string
	i := New(func(c context.Context, f string) string {
		result = "wrapped_" + f
		return result
	})
	i2 := i.After(func(c context.Context, f string) string {
		return "bob"
	})
	i2.ImportFiles([]string{"hello"})
	assert.Equal(t, "wrapped_bob", result, "failed wrapping other Importer")
}
func TestImportFiles(t *testing.T) {
	var processedFiles []string
	files := []string{"a", "b", "c"}
	process := func(ctx context.Context, f string) string {
		processedFiles = append(processedFiles, f)
		return ""
	}
	i := New(process)
	i.ImportFiles(files)
	assert.Equal(t, files, processedFiles, fmt.Sprintf("processedFiles do not equal input files\n processFiles=%v\ninputFiles=%v", processedFiles, files))
}
