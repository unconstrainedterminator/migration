package migration

import (
	"github.com/unconstrainedterminator/os"
	"path/filepath"
)

func NewPath(dirname string) string {
	path := os.GetCurrentPath()
	dbDir := "db"
	return filepath.Join(path, dbDir, dirname)
}
