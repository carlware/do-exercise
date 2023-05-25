package main

import (
	"flag"
	"fmt"
	"geoextractor-go/extractor"
	"geoextractor-go/extractor/writers"
	"log"
	"os"
)

func main() {
	var inputName string
	var outputName string

	flag.StringVar(&inputName, "input", "", "image filepath or directory")
	flag.StringVar(&outputName, "output", "output.csv", "filename of output")
	flag.Parse()

	if inputName == "" {
		fmt.Println("missing --input argument")
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

	fileWriter, err := os.Create(outputName)
	if err != nil {
		log.Fatalf("error while creating the file")
	}

	recordWriter := writers.NewCsvWriter(fileWriter)
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
