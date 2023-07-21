package backend

import (
	"io/fs"
	"os"
)

func ListFolder(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}
