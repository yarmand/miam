package importer

import (
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/afero"
)

// GetCreationDate extract creation date from exif infos
func GetCreationDate(fs afero.Fs, fname string) (time.Time, error) {
	x, err := decode(fs, fname)
	if err != nil {
		return time.Now(), err
	}
	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	return tm, nil
}

func decode(fs afero.Fs, fname string) (*exif.Exif, error) {
	f, err := fs.Open(fname)
	if err != nil {
		return nil, err
	}

	x, err := exif.Decode(f)
	if err != nil {
		return nil, err
	}

	return x, nil
}
