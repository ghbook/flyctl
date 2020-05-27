package presenters

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
)

// Presentable - Records (and field names) which may be presented by a Presenter
type Presentable interface {
	FieldNames() []string
	Records() []map[string]string
}

// Presenter - A self managing presenter which can be rendered in multiple ways
type Presenter struct {
	Item Presentable
	Out  io.Writer
	Opts Options
}

// Options - Presenter options
type Options struct {
	Vertical   bool
	HideHeader bool
}

// Render - Renders a presenter as a field list or table
func (p *Presenter) Render() error {
	if p.Opts.Vertical {
		return p.renderFieldList()
	}

	return p.renderTable()
}

func (p *Presenter) renderTable() error {
	table := tablewriter.NewWriter(p.Out)

	cols := p.Item.FieldNames()

	if !p.Opts.HideHeader {
		table.SetHeader(cols)
	}
	table.SetBorder(false)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnSeparator(" ")
	table.SetNoWhiteSpace(true)
	table.SetTablePadding(" ") // pad with tabs

	for _, kv := range p.Item.Records() {
		fields := []string{}
		for _, col := range cols {
			fields = append(fields, kv[col])
		}
		table.Append(fields)
	}

	table.Render()

	fmt.Fprintln(p.Out)

	return nil
}

func (p *Presenter) renderFieldList() error {
	table := tablewriter.NewWriter(p.Out)

	cols := p.Item.FieldNames()

	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetColumnSeparator("=")
	table.SetColumnAlignment([]int{tablewriter.ALIGN_DEFAULT, tablewriter.ALIGN_LEFT})

	for _, kv := range p.Item.Records() {
		for _, col := range cols {
			table.Append([]string{col, kv[col]})
		}
		table.Render()

		fmt.Fprintln(p.Out)
	}

	return nil
}
