package importer

import (
	"context"
)

// Processor is a function that process an image read from disk
type Processor func(ctx context.Context, filename string)

// ImportFiles get a list of files and run processor on them.
func ImportFiles(fileNames []string, process Processor) {
	for _, fileName := range fileNames {
		process(nil, fileName)
	}
}
