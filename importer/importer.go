package importer

import (
	"context"
	"io/fs"
	"log"

	"github.com/spf13/afero"
)

// Processor is a function that process an image read from disk
type Processor func(ctx context.Context, importer Importer, filename string) (processedFilename string, err error)

type Importer struct {
	process Processor
	source  string
	appFS   afero.Fs
	logger  *log.Logger
}

func New(fs afero.Fs, src string, processor Processor) (importer Importer) {
	importer = Importer{process: processor, source: src, appFS: fs}
	return
}

// Then creates a new Importer that will execute this importer processor on the result
// of the processor passed as argument.
func (i Importer) Then(processor Processor) Importer {
	return New(i.appFS, i.source, func(ctx context.Context, importer Importer, filename string) (processedFilename string, err error) {
		firstOut, err := i.process(ctx, i, filename)
		if err != nil {
			return
		}
		return processor(ctx, i, firstOut)
	})
}

func (i Importer) Logger() *log.Logger {
	if i.logger == nil {
		i.logger = log.Default()
	}
	return i.logger
}

// Start importing files from source in go routines.
// Import should be called only once, after having build a chain of
// processors using WrapedIn() and Then().
func (i Importer) Import() error {
	return afero.Walk(i.appFS, i.source, func(path string, info fs.FileInfo, err error) error {
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
