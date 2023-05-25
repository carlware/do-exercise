package main

import (
	"flag"
	"fmt"
	"geoextractor-go/extractor"
	"geoextractor-go/extractor/processors"
	"geoextractor-go/extractor/writers"
	"log"
	"os"
)

func main() {
	var input string
	var output string

	flag.StringVar(&input, "input", "", "image filepath or directory")
	flag.StringVar(&output, "output", "output.csv", "filename of output")
	flag.Parse()

	if input == "" {
		fmt.Println("missing --input argument")
		os.Exit(0)
	}

	files, err := extractor.ReadFiles(input)
	if err != nil {
		log.Fatalf("error while reading the files %s", err)
	}

	exifImagesData, err := extractor.ExtractExifDataFromImages(files)
	if err != nil {
		log.Fatalf("error while extracting %s", err)
	}

	writer, err := writers.NewCsvWriter(output)
	defer writer.Flush()
	if err != nil {
		log.Fatalf("error while creting output file %s", err)
	}

	processor := extractor.NewTagProcessor(exifImagesData)

	err = processor.
		AddTagHandler("filename", processors.GetFileName).
		AddTagHandler("latitude", processors.GetLatitude).
		AddTagHandler("longitude", processors.GetLongitude).
		Write(writer)
	if err != nil {
		log.Fatalf("error while creting output file %s", err)
	}

}
