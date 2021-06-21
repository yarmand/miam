package importer

import (
	"context"
	"io/fs"

	"github.com/spf13/afero"
)

// Processor is a function that process an image read from disk
type Processor func(ctx context.Context, importer Importer, filename string) (processedFilename string, err error)

type Importer struct {
	process     Processor
	source      string
	destination string
	appFS       afero.Fs
}

func New(fs afero.Fs, src string, dest string, processor Processor) (importer Importer) {
	importer = Importer{process: processor, source: src, destination: dest, appFS: fs}
	return
}

// After creates a new Importer that will execute this importer processor on the result
// of the processor passed as argument.
func (i Importer) After(processor Processor) Importer {
	return New(i.appFS, i.source, i.destination, func(ctx context.Context, importer Importer, filename string) (procesedFilename string, err error) {
		procesedFilename, err = processor(ctx, i, filename)
		if err != nil {
			return
		}
		return i.process(ctx, i, procesedFilename)
	})
}

// WrapedIn creates a new Importer that will execute this importer and pass itsresult to
// the Processor passed as argument.
func (i Importer) WrapedIn(processor Processor) Importer {
	return New(i.appFS, i.source, i.destination, func(ctx context.Context, importer Importer, filename string) (processedFilename string, err error) {
		processedFilename, err = i.process(ctx, i, filename)
		if err != nil {
			return
		}
		return processor(ctx, i, processedFilename)
	})
}

// Start importing files from source
func (i Importer) Import() {
	afero.Walk(i.appFS, i.source, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			_, err := i.process(context.TODO(), i, path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// ImportFiles get a list of files and run processor on them.
func (i Importer) ImportFiles(fileNames []string) {
	for _, fileName := range fileNames {
		i.process(context.TODO(), i, fileName)
	}
}
