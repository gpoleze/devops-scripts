package tablerenderer

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"reflect"
)

type getRowFromItem[T any] func(T) table.Row

func BuildTable[T any](listOfItems []T, fn getRowFromItem[T]) {
	var header table.Row

	for _, field := range reflect.VisibleFields(reflect.TypeOf(listOfItems[0])) {
		header = append(header, field.Name)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(header)

	for _, item := range listOfItems {
		t.AppendRow(fn(item))
	}

	t.Render()
	return
}
