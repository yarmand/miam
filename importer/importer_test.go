package importer

import (
	"fmt"
	"io"
	"testing"
	"time"

	"os"
	"path"

	"context"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestProcessorFunctions(t *testing.T) {
	t.Run("Then run the parameter funtion after the left processor function", func(t *testing.T) {
		result := ""
		i := New(nil, "", func(c context.Context, i Importer, f string) (string, error) {
			return "before_", nil
		})
		i2 := i.Then(func(c context.Context, i Importer, f string) (string, error) {
			result = f + "after"
			return result, nil
		})
		i2.ImportFiles([]string{"hello"})
		assert.Equal(t, "before_after", result, "failed wrapping other Importer")
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
		i := New(nil, "", process)
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
		i := New(appFS, src, func(c context.Context, i Importer, f string) (string, error) {
			processedFiles = append(processedFiles, path.Base(f))
			return "", nil
		})
		i.Import()
		assert.Equal(t, []string{"a", "b", "c", "d"}, processedFiles, "processedFiles do not contains all files created in src folder")
	})
}

func TestProcessors(t *testing.T) {
	var appFS afero.Fs = afero.NewMemMapFs()
	var osFS afero.Fs = afero.NewOsFs()
	src := fmt.Sprintf("%s/%v", "/tmp/src", time.Now().Unix())
	appFS.MkdirAll(src, os.ModePerm)
	defer appFS.Remove(src)
	copyFile(osFS, "../testAssets/testImage.jpg", appFS, fmt.Sprintf("%s/testImage.jpg", src))
	copyFile(osFS, "../testAssets/testImage2.jpg", appFS, fmt.Sprintf("%s/testImage2.jpg", src))

	dest := fmt.Sprintf("%s/%v", "/tmp/dest", time.Now().Unix())
	defer appFS.Remove(dest)
	t.Run("PMoveToDest move file to date bas folder", func(t *testing.T) {
		i := New(appFS, src, PLogFilename("Processing: ", "")).
			Then(PMoveToDateFolder(dest)).
			Then(PLogFilename("processed file:", ""))
		i.Import()
		stats, err := appFS.Stat(fmt.Sprintf("%s/2018/06/25/testImage.jpg", dest))
		assert.Nil(t, err, "error checking moved file: %v", err)
		assert.Equal(t, "testImage.jpg", stats.Name(), "wrong moved image name")
		stats, err = appFS.Stat(fmt.Sprintf("%s/2019/06/23/testImage2.jpg", dest))
		assert.Nil(t, err, "error checking moved file: %v", err)
		assert.Equal(t, "testImage2.jpg", stats.Name(), "wrong moved image name")
	})
}

func copyFile(srcFS afero.Fs, srcPath string, destFS afero.Fs, destPath string) error {
	// copy test image in src
	from, err := srcFS.Open(srcPath)
	if err != nil {
		panic(fmt.Sprintf("error opening src file %s", srcPath))
	}
	defer from.Close()
	to, _ := destFS.OpenFile(destPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer to.Close()
	io.Copy(to, from)
	return nil
}
