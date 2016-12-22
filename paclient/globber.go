package paclient

import (
	"os"
	"path/filepath"
)

func NewLocalGlobber() *GlobFinder {
	return &GlobFinder{
		Glob:      filepath.Glob,
		Separator: filepath.Separator,
		Abs:       filepath.Abs,
	}
}

type GlobFinder struct {
	Glob      func(string) ([]string, error)
	Find      func(string) ([]string, error)
	Separator uint8
	Abs       func(string) (string, error)
}

type LocalGlobFinder struct {
}

func (l *LocalGlobFinder) Glob(pre string) ([]string, error) {
	return filepath.Glob(pre)
}

func (l *LocalGlobFinder) Abs(input string) (string, error) {
	return filepath.Abs(input)
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
