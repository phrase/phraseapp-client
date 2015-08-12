package cli

import (
	"fmt"
	"strings"
)

type row []string

type table struct {
	rows     []row
	colWidth []int
}

func (t *table) addRow(r row) {
	t.rows = append(t.rows, r)
	for colIdx := range r {
		l := len(r[colIdx])
		if colIdx < len(t.colWidth) {
			if t.colWidth[colIdx] < l {
				t.colWidth[colIdx] = l
			}
		} else {
			t.colWidth = append(t.colWidth, l)
		}
	}
}

func (t *table) String() string {
	lines := make([]string, len(t.rows))
	for rowIdx := range t.rows {
		line := ""
		for colIdx := range t.rows[rowIdx] {
			col := t.rows[rowIdx][colIdx]
			line += fmt.Sprintf("%s%-*s", col, t.colWidth[colIdx]+1-len(col), " ")
		}
		lines[rowIdx] = line
	}
	return strings.Join(lines, "\n")
}
