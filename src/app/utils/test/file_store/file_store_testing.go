package utilstestingfilestore

import (
	"io/ioutil"
	"os"
	"testing"
)

func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
