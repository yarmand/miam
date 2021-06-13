package importer

import (
	"context"
)

// Processor is a function that process an image read from disk
type Processor func(ctx context.Context, filename string) (processedFilename string)

type Importer struct {
	process Processor
}

func New(processor Processor) (importer Importer) {
	importer = Importer{process: processor}
	return
}

// After creates a new Importer that will execute this importer processor on the result
// of the processor passed as argument.
func (i Importer) After(processor Processor) Importer {
	return New(func(ctx context.Context, filename string) string {
		f := processor(ctx, filename)
		return i.process(ctx, f)
	})
}

// WrapedIn creates a new Importer that will execute this importer and pass itsresult to
// the Processor passed as argument.
func (i Importer) WrapedIn(processor Processor) Importer {
	return New(func(ctx context.Context, filename string) string {
		f := i.process(ctx, filename)
		return processor(ctx, f)
	})
}

// ImportFiles get a list of files and run processor on them.
func (i Importer) ImportFiles(fileNames []string) {
	for _, fileName := range fileNames {
		i.process(context.TODO(), fileName)
	}
}
