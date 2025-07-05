package table

import (
	"github.com/alexeyco/simpletable"
)

type Table struct {
	table *simpletable.Table
}

// NewTable creates a new Table instance
func NewTable() *Table {
	t := simpletable.New()
	return &Table{table: t}
}

// SetHeaders sets the headers for the table
func (t *Table) SetHeaders(headers ...string) {
	cells := make([]*simpletable.Cell, len(headers))
	for i, h := range headers {
		cells[i] = &simpletable.Cell{Align: simpletable.AlignCenter, Text: h}
	}
	t.table.Header = &simpletable.Header{Cells: cells}
}

// AddRow adds a row to the table
func (t *Table) AddRow(row ...string) {
	cells := make([]*simpletable.Cell, len(row))
	for i, cell := range row {
		cells[i] = &simpletable.Cell{Text: cell}
	}
	t.table.Body.Cells = append(t.table.Body.Cells, cells)
}

// Print prints the table to stdout
func (t *Table) Print() {
	t.table.SetStyle(simpletable.StyleUnicode)
	t.table.Println()
}
