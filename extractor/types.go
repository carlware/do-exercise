package extractor

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type handleRowFunc func(info ImageInfo) string

type RecordWriter interface {
	Write(record []string) error
	Flush()
}

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

type ImageInfo struct {
	Filename string
	IfdIndex exif.IfdIndex
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func ReadFiles(filepath string) (map[string][]byte, error) {
	fileDataByName := make(map[string][]byte)
	isDir, err := isDirectory(filepath)
	if err != nil {
		return nil, err
	}

	if !isDir {
		imageBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			return fileDataByName, err
		}
		fileDataByName[filepath] = imageBytes
	} else  {
		files, err := ioutil.ReadDir(filepath)
		if err != nil {
			return fileDataByName, err
		}

		for _, file := range files {
			filename := fmt.Sprintf("%s/%s", filepath, file.Name())
			filesByName, err := ReadFiles(filename)
			if err != nil {
				return nil, err
			}
			for name, data := range filesByName {
				fileDataByName[name] = data
			}
		}
	}

	return fileDataByName, nil
}

func ExtractExifDataFromImages(imageByName map[string][]byte) ([]ImageInfo, error) {
	imagesInfo := make([]ImageInfo, 0, len(imageByName))
	for imageName, data := range imageByName {
		exifRaw, err := exif.SearchAndExtractExif(data)
		if err != nil {
			log.Printf("%s: No EXIF data\n", imageName)
			continue
		}
		im, err := exifcommon.NewIfdMappingWithStandard()
		ti := exif.NewTagIndex()

		_, index, err := exif.Collect(im, ti, exifRaw)
		imagesInfo = append(imagesInfo, ImageInfo{
			Filename: imageName,
			IfdIndex: index,
		})
	}

	return imagesInfo, nil
}

type CSVWriter struct {
	columns       []string
	imagesInfo    []ImageInfo
	processorRows []handleRowFunc
}

func Processor(imagesInfo []ImageInfo) CSVWriter {
	return CSVWriter{
		columns:    []string{},
		imagesInfo: imagesInfo,
	}
}

func (w CSVWriter) AddField(name string, processor handleRowFunc) CSVWriter {
	w.columns = append(w.columns, name)
	w.processorRows = append(w.processorRows, processor)
	return w
}

func (w CSVWriter) Write(writer RecordWriter) error {
	if err := writer.Write(w.columns); err != nil {
		return err
	}

	for _, imageInfo := range w.imagesInfo {
		record := make([]string, len(w.columns))
		for idx, processor := range w.processorRows {
			data := processor(imageInfo)
			record[idx] = data
		}
		if err := writer.Write(record); err!= nil {
			return err
		}
	}

	return nil
}
