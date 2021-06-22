package importer

import "context"

func PLog(ctx context.Context, i Importer, filename string) (processedFilename string, err error) {
	i.Logger().Printf("[PLog] PRocessing file %s\n", filename)
	return filename, nil
}
