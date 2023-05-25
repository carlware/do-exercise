package extractor

import (
	"github.com/dsoprea/go-exif/v3"
)

type TagHandleFunc func(info FileInfo) (string, error)

type RecordWriter interface {
	Write(record []string) error
	Flush()
}

type FileInfo struct {
	Filename string
	IfdIndex exif.IfdIndex
}
