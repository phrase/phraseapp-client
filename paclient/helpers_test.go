package paclient

import (
	"fmt"
	"sort"
)

type ByPath []*LocaleFile

func (a ByPath) Len() int           { return len(a) }
func (a ByPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPath) Less(i, j int) bool { return a[i].Path < a[j].Path }

func compareLocaleFiles(actualFiles LocaleFiles, expectedFiles LocaleFiles) error {
	sort.Sort(ByPath(actualFiles))
	sort.Sort(ByPath(expectedFiles))
	for idx, actualFile := range actualFiles {
		expected := expectedFiles[idx]
		actual := actualFile
		if expected.Path != actual.Path {
			return fmt.Errorf("Expected Path %s should eql %s", expected.Path, actual.Path)
		}
		if expected.Name != actual.Name {
			return fmt.Errorf("Expected Name %s should eql %s", expected.Name, actual.Name)
		}
		if expected.Code != actual.Code {
			return fmt.Errorf("Expected Code %s should eql %s", expected.Code, actual.Code)
		}
		if expected.ID != actual.ID {
			return fmt.Errorf("Expected ID %s should eql %s", expected.ID, actual.ID)
		}
	}
	return nil
}
