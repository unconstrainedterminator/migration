package migration

import (
	"context"
	"github.com/unconstrainedterminator/os"
	"path/filepath"
)

func init() {
	path := os.GetCurrentAbs()
	dirname := "init"
	dirname = filepath.Join(path, "init")
	Run(context.Background(), dirname)
}
