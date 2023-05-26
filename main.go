package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"geoextractor-go/extractor"
)

func main() {
	var inputName string
	var outputName string
	var format string

	flag.StringVar(&inputName, "input", "", "image filepath or directory")
	flag.StringVar(&outputName, "output", "output", "filename of output")
	flag.StringVar(&format, "format", "csv", "output format")
	flag.Parse()

	if inputName == "" {
		fmt.Println("missing --input argument")
		os.Exit(0)
	}

	if format != "csv" && format != "html" {
		fmt.Printf("unsupported format %s\n", format)
		os.Exit(0)
	}

	files, err := extractor.ReadFiles(inputName)
	if err != nil {
		log.Fatalf("error while reading the files %s", err)
	}

	exifImagesData, err := extractor.ExtractExifDataFromImages(files)
	if err != nil {
		log.Fatalf("error while extracting %s", err)
	}

	fileWriter, err := os.Create(fmt.Sprintf("%s.%s", outputName, format))
	if err != nil {
		log.Fatalf("error while creating the file")
	}

	recordWriter := extractor.CreateRecordWriter(format, fileWriter)
	if recordWriter == nil {
		log.Fatalf("error while creting the record writer")
	}
	processor := extractor.NewTagProcessor(exifImagesData)

	err = processor.
		AddTagHandler("filename", extractor.GetFileName).
		AddTagHandler("latitude", extractor.GetLatitude).
		AddTagHandler("longitude", extractor.GetLongitude).
		Write(recordWriter)
	if err != nil {
		log.Fatalf("error while creating output file %s", err)
	}

}
