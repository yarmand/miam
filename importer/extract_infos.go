package importer

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// GetCreationDate extract creation date from exif infos
func GetCreationDate(fname string) time.Time {
	x := decode(fname)
	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	return tm
}

func decode(fname string) *exif.Exif {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return x
}
