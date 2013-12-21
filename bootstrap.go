package gosimple

import (
	"os"
	"path/filepath"
)

var templatesPath = "templates"

func init() {
	dir, _ := os.Getwd()
	templatesPath = filepath.Join(dir, templatesPath)
}

func Bootstrap() (err error) {
	Dispatch()
	return nil
}
