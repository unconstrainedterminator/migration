package migration

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Run(ctx context.Context, dirname string) {
	files, err := getFiles(dirname)
	if err != nil {
		return
	}

	start, err := g.DB().Begin(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range files {
		var read []byte
		read, err = readAll(filepath.Join(dirname, file.Name()))
		if err != nil {
			log.Fatal(err)
			return
		}

		err = call(read, start)
		if err != nil {
			err = start.Rollback()
			log.Println(err)
			return
		}
	}

	err = start.Commit()
	if err != nil {
		msg := fmt.Sprintf("DB %s failed", dirname)
		log.Println(msg)
		return
	}

	msg := fmt.Sprintf("DB %s successfully", dirname)
	log.Println(msg)
	return
}

func getFiles(dirname string) (files []os.FileInfo, err error) {
	var file *os.File
	file, err = os.Open(dirname)
	if err != nil {
		log.Fatal(err)
		return
	}

	files, err = file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func readAll(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

func convert(str string) string {
	sql, _ := gregex.ReplaceStringFunc("\\?", str, func(s string) string {
		return fmt.Sprintf("%s", "?")
	})
	return sql
}

func call(script []byte, start gdb.TX) (err error) {
	sql := strings.Split(string(script), ";")
	for _, s := range sql {
		if strings.Contains(s, "?") {
		}

		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		_, err = start.Exec(s)
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func ReadFileAll(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}
