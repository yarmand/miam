package importer

import "testing"

func TestGetCreationDate(t *testing.T) {
	fname := "../testAssets/testImage.jpeg"

	date := GetCreationDate(fname)
	if !(date.Year() == 2018 && date.Month() == 6 && date.Day() == 24) {
		t.Error("Incorect date\nexpected: 2018-06-25 ...\ngot:", date)
	}
}
