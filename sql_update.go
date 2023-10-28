package migration

import (
	"context"
)

func init() {
	dirname := "update"
	dirname = NewPath(dirname)
	Run(context.Background(), dirname)
}
