package gotable

// Table holds the table state
type Table struct {
	columns []Column
	output  Output
}

// NewTable creates a new table
func NewTable(columns []Column) Table {
	return Table{
		columns: columns,
	}
}

// Print outputs the table to the output object
func (t *Table) Print(output Output) {

}
