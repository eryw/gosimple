package config

import (
	"os"
	"path/filepath"
)

var TemplatesPath string

func init() {
	dir, _ := os.Getwd()
	TemplatesPath = filepath.Join(dir, "templates")
}
