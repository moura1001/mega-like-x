package store

import (
	"io/ioutil"
	utilstestingfilestore "moura1001/mega_like_x/src/app/utils/test/file_store"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := utilstestingfilestore.CreateTempFile(t, "abcdef")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("123"))

	file.Seek(0, 0)
	newFileContent, _ := ioutil.ReadAll(file)

	got := string(newFileContent)
	want := "123"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
