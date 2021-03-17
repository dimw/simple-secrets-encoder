package fileutils

import (
	"github.com/gobwas/glob"
	"os"
	"path/filepath"
)

func Glob(workdir string, pattern string) ([]string, error) {
	fileSearchGlob := glob.MustCompile(pattern)

	var files []string

	err := filepath.Walk(workdir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if fileSearchGlob.Match(filepath.ToSlash(path)) {
					files = append(files, path)
				}
			}
			return nil
		},
	)

	return files, err
}
