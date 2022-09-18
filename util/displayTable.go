package util

import (
	"os"

	"github.com/lensesio/tableprinter"
)

func DisplayTable(data interface{}) {
	printer := tableprinter.New(os.Stdout)

	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"

	printer.Print(data)
}
