package writers

import (
	"encoding/csv"
	"io"
)

type CsvWriter struct {
	csvWriter *csv.Writer
}

func NewCsvWriter(writer io.Writer) *CsvWriter {
	return &CsvWriter{
		csvWriter: csv.NewWriter(writer),
	}
}

func (w *CsvWriter) Write(record []string) error {
	return w.csvWriter.Write(record)
}

func (w *CsvWriter) Flush() {
	w.csvWriter.Flush()
}
