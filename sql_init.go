package migration

import (
	"context"
)

func init() {
	dirname := "init"
	dirname = NewPath(dirname)
	Run(context.Background(), dirname)
}
