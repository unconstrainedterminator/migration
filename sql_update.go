package migration

import (
	"context"
	"github.com/unconstrainedterminator/os"
	"path/filepath"
)

func init() {
	path := os.GetCurrentAbs()
	dirname := "update"
	dirname = filepath.Join(path, dirname)
	Run(context.Background(), dirname)
}
