package writers

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

//go:embed template.html.tpl
var htmlTemplate string

type HtmlWriter struct {
	writer  io.Writer
	columns []string
	records []map[string]interface{}
}

func NewHtmlWriter(writer io.Writer) *HtmlWriter {
	return &HtmlWriter{
		writer: writer,
	}
}

func (w *HtmlWriter) Write(record []string) error {
	if len(w.columns) == 0 {
		w.columns = record
	} else if len(record) == len(w.columns) {
		details := make(map[string]interface{})
		for idx, value := range record {
			details[w.columns[idx]] = value
		}
		w.records = append(w.records, details)
	}
	return nil
}

func (w *HtmlWriter) Flush() {
	columns := make([]string, len(w.columns))
	for idx, name := range w.columns {
		columns[idx] = fmt.Sprintf("'%s'", name)
	}
	columnsStr := strings.Join(columns, ",")
	payload, _ := json.Marshal(w.records)
	output := fmt.Sprintf(htmlTemplate, columnsStr, payload)
	w.writer.Write([]byte(output))
}
