package importer

import (
	sql "github.com/jmoiron/sqlx"
)

// Processor is a function that process the image.
// it can be wrapped or followed by another processor.
type Processor func(importer Importer, nextProcessor *Processor)

// Importer get images from a **source**, move them in a date based **ImportRoot** and
// insert reference in a **importDB**
type Importer struct {
	source     string
	importRoot string
	importDB   sql.DB
	processor  Processor
}
