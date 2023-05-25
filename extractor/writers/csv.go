package writers

import (
	"encoding/csv"
	"os"
)

type CsvWriter struct {
	csvWriter *csv.Writer
	file *os.File
}

func NewCsvWriter(filename string) (*CsvWriter, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &CsvWriter{
		csvWriter: csv.NewWriter(f),
		file: f,
	}, nil
}

func (w *CsvWriter) Write(record []string) error {
	return w.csvWriter.Write(record)
}

func (w *CsvWriter) Flush() {
	w.csvWriter.Flush()
	w.file.Close()
}
