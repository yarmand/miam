package importer

import (
	"fmt"
	"testing"

	"os"
	"path"

	"context"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestProcessorFunctions(t *testing.T) {
	t.Run("WrapIn use result of left processor function as input to the wrap function", func(t *testing.T) {
		var result string
		i := New(nil, "", "", func(c context.Context, i Importer, f string) (string, error) {
			return "bob", nil
		})
		i2 := i.WrapedIn(func(c context.Context, i Importer, f string) (string, error) {
			result = "wrapped_" + f
			return result, nil
		})
		i2.ImportFiles([]string{"hello"})
		assert.Equal(t, "wrapped_bob", result, "failed wrapping other Importer")
	})

	t.Run("After run the parameter funtion after the left processor function", func(t *testing.T) {
		var result string
		i := New(nil, "", "", func(c context.Context, i Importer, f string) (string, error) {
			result = "before_" + f
			return result, nil
		})
		i2 := i.After(func(c context.Context, i Importer, f string) (string, error) {
			return "bob", nil
		})
		i2.ImportFiles([]string{"hello"})
		assert.Equal(t, "before_bob", result, "failed wrapping other Importer")
	})
}

func TestImportFiles(t *testing.T) {
	t.Run("ImportFiles apply the processor function to all imported files", func(t *testing.T) {
		var processedFiles []string
		files := []string{"a", "b", "c"}
		process := func(ctx context.Context, i Importer, f string) (string, error) {
			processedFiles = append(processedFiles, f)
			return "", nil
		}
		i := New(nil, "", "", process)
		i.ImportFiles(files)
		assert.Equal(t, files, processedFiles, fmt.Sprintf("processedFiles do not equal input files\n processFiles=%v\ninputFiles=%v", processedFiles, files))
	})
}

func TestStart(t *testing.T) {
	var appFS afero.Fs = afero.NewMemMapFs()
	src := "/tmp/src"
	dest := "/tmp/dest"
	appFS.MkdirAll(src, os.ModePerm)
	appFS.Create(fmt.Sprintf("%s/a", src))
	appFS.Create(fmt.Sprintf("%s/b", src))
	appFS.Create(fmt.Sprintf("%s/c", src))
	sub := fmt.Sprintf("%s/subdir", src)
	appFS.MkdirAll(sub, os.ModePerm)
	appFS.Create(fmt.Sprintf("%s/d", sub))
	appFS.MkdirAll(dest, os.ModePerm)

	t.Run("Import process all files", func(t *testing.T) {
		var processedFiles []string
		i := New(appFS, src, dest, func(c context.Context, i Importer, f string) (string, error) {
			processedFiles = append(processedFiles, path.Base(f))
			return "", nil
		})
		i.Import()
		assert.Equal(t, []string{"a", "b", "c", "d"}, processedFiles, "processedFiles do not contains all files created in src folder")
	})
}
