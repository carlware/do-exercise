package extractor

import (
	"fmt"
	"log"
)

type TagProcessor struct {
	columns       []string
	imagesInfo    []FileInfo
	processorRows []TagHandleFunc
}

func NewTagProcessor(imagesInfo []FileInfo) TagProcessor {
	return TagProcessor{
		imagesInfo: imagesInfo,
	}
}

func (w TagProcessor) AddTagHandler(name string, processor TagHandleFunc) TagProcessor {
	w.columns = append(w.columns, name)
	w.processorRows = append(w.processorRows, processor)
	return w
}

func (w TagProcessor) Write(writer RecordWriter) error {
	if err := writer.Write(w.columns); err != nil {
		return fmt.Errorf("error while writing %w", err)
	}

Exit:
	for _, imageInfo := range w.imagesInfo {
		record := make([]string, len(w.columns))
		for idx, processor := range w.processorRows {
			data, err := processor(imageInfo)
			if err != nil {
				log.Printf("cannot process tag=%s for image=%s", w.columns[idx], imageInfo.Filename)
				continue Exit
			}
			record[idx] = data
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error while writing %w", err)
		}
	}
	writer.Flush()

	return nil
}
