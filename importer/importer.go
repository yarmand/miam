package importer

import (
	"context"

	"github.com/spf13/afero"
)

// Processor is a function that process an image read from disk
type Processor func(ctx context.Context, importer Importer, filename string) (processedFilename string)

type Importer struct {
	process     Processor
	source      string
	destination string
	appFS       afero.Fs
}

func New(processor Processor, src string, dest string, fs afero.Fs) (importer Importer) {
	importer = Importer{process: processor, source: src, destination: dest, appFS: fs}
	return
}

// After creates a new Importer that will execute this importer processor on the result
// of the processor passed as argument.
func (i Importer) After(processor Processor) Importer {
	return New(func(ctx context.Context, importer Importer, filename string) string {
		f := processor(ctx, i, filename)
		return i.process(ctx, i, f)
	}, i.source, i.destination, i.appFS)
}

// WrapedIn creates a new Importer that will execute this importer and pass itsresult to
// the Processor passed as argument.
func (i Importer) WrapedIn(processor Processor) Importer {
	return New(func(ctx context.Context, importer Importer, filename string) string {
		f := i.process(ctx, i, filename)
		return processor(ctx, i, f)
	}, i.source, i.destination, i.appFS)
}

// Start importing files from source
func (i Importer) Start() {
}

// ImportFiles get a list of files and run processor on them.
func (i Importer) ImportFiles(fileNames []string) {
	for _, fileName := range fileNames {
		i.process(context.TODO(), i, fileName)
	}
}
