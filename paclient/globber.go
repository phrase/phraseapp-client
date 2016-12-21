package paclient

import (
	"os"
	"path/filepath"
)

type GlobFinder interface {
	Glob(pre string) ([]string, error)
	Find(candidate string) ([]string, error)
	Separator() uint8
}

type LocalGlobFinder struct {
}

func (l *LocalGlobFinder) Glob(pre string) ([]string, error) {
	return filepath.Glob(pre)
}

func (l *LocalGlobFinder) Find(candidate string) (files []string, err error) {
	err = filepath.Walk(candidate, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.Mode().IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func (l *LocalGlobFinder) Separator() uint8 {
	return os.PathSeparator
}
