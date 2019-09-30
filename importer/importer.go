package importer

import (
	"context"
)

// Importer is responsible to feed a chain of processors
type Importer struct {
	processor Processor
}

// Processor is a function that process an image read from disk
type Processor func(ctx context.Context, filename string)

// ImportFile improt a file by aplying the processor
func (importer Importer) ImportFile(fileName string) {
	importer.processor(context.Background(), fileName)
}

func nothing(ctx context.Context, filename string) {

}

// NewImporter create an base importer that does nothing
func NewImporter() (importer Importer) {
	importer = Importer{processor: nothing}
	return
}

// AddProcessor add a *Processor* to the chain of processors used by the *Importer*
func (importer Importer) AddProcessor(processor Processor) {
	importer.processor = func(ctx context.Context, filename string) {
		importer.processor(ctx, filename)
		processor(ctx, filename)
	}
}
