package importer

import (
	"context"
	"fmt"
	"os"
	"path"
)

func PLog(ctx context.Context, i Importer, filename string) (processedFilename string, err error) {
	i.Logger().Printf("[PLog] PRocessing file %s\n", filename)
	return filename, err
}

func PMoveToDest(ctx context.Context, i Importer, filename string) (processedFilename string, err error) {
	date := GetCreationDate(i.appFS, filename)
	destdatePath := fmt.Sprintf("%s/%d/%02d/%02d", i.destination, date.Year(), date.Month(), date.Day())
	destfilename := fmt.Sprintf("%s/%s", destdatePath, path.Base(filename))
	i.appFS.MkdirAll(destdatePath, os.ModePerm)
	i.appFS.Rename(filename, destfilename)
	return processedFilename, err
}
