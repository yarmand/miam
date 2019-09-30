package importer

import (
	"testing"

	"context"

	"github.com/google/go-cmp/cmp"
)

func TestImportFiles(t *testing.T) {
	var processedFiles []string
	files := []string{"a", "b", "c"}
	process := func(ctx context.Context, f string) { processedFiles = append(processedFiles, f) }
	ImportFile(files[0])
	if !cmp.Equal(files, processedFiles) {
		t.Errorf("processedFiles do not equal input files\n processFiles=%v\ninputFiles=%v", processedFiles, files)
	}
}
