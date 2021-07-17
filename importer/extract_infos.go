package importer

import (
	"log"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/afero"
)

// GetCreationDate extract creation date from exif infos
func GetCreationDate(fs afero.Fs, fname string) time.Time {
	x := decode(fs, fname)
	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	return tm
}

func decode(fs afero.Fs, fname string) *exif.Exif {
	f, err := fs.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return x
}
