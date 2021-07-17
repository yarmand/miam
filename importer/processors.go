package importer

import (
	"context"
	"fmt"
	"os"
	"path"
)

// PLogFilename print to the importer logger the beforeString filename and afterString
func PLogFilename(beforeString string, afterString string) Processor {
	return func(ctx context.Context, i Importer, filename string) (processedFilename string, err error) {
		i.Logger().Printf("%s%s%s\n", beforeString, filename, afterString)
		return filename, err
	}
}

// PMoveToDateFolder move file to destination directory, into a folder
// correspoding to the photo creation date.
func PMoveToDateFolder(destPath string) Processor {
	return func(ctx context.Context, i Importer, filename string) (processedFilename string, err error) {
		date := GetCreationDate(i.appFS, filename)
		destdatePath := fmt.Sprintf("%s/%d/%02d/%02d", destPath, date.Year(), date.Month(), date.Day())
		destfilename := fmt.Sprintf("%s/%s", destdatePath, path.Base(filename))
		i.appFS.MkdirAll(destdatePath, os.ModePerm)
		i.appFS.Rename(filename, destfilename)
		return processedFilename, err
	}
}
