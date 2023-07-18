package services

import (
	"archive/zip"
	"bytes"
	_ "embed"

	"github.com/spf13/afero"
	"github.com/spf13/afero/zipfs"
)

//go:generate bash -x generate-prod.sh

//go:embed ui.zip
var uiData []byte

var ui afero.Fs

func init() {
	zipReader, err := zip.NewReader(bytes.NewReader(uiData), int64(len(uiData)))
	if err != nil {
		panic(err)
	}
	ui = zipfs.New(zipReader)
}
