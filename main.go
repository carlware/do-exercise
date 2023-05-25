package main

import (
	"flag"
	"fmt"
	"geoextractor-go/extractor"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"io/ioutil"
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

	writer, err := extractor.NewCsvWriter(output)
	defer writer.Flush()
	if err != nil {
		fmt.Printf("err %s\n", err)
	}

	processor := extractor.Processor(exifImagesData)

	err = processor.
		AddField("filename", func(info extractor.ImageInfo) string {
			return info.Filename
		}).
		AddField("latitude", func(info extractor.ImageInfo) string {
			ifdTag, _ := info.IfdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
			gpsInfo, _ := ifdTag.GpsInfo()
			if gpsInfo.Latitude.Decimal() == 0 {
				return "N/A"
			}
			return fmt.Sprintf("%f", gpsInfo.Latitude.Decimal())
		}).
		AddField("longitude", func(info extractor.ImageInfo) string {
			ifdTag, _ := info.IfdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
			gpsInfo, _ := ifdTag.GpsInfo()
			if gpsInfo.Longitude.Decimal() == 0 {
				return "N/A"
			}
			return fmt.Sprintf("%f", gpsInfo.Longitude.Decimal())
		}).
		Write(writer)
	if err != nil {
		fmt.Printf("error write %s\n", err)
	}

}

func main1() {
	filename := "./images/more_images/wax-card.jpg"

	imageBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("err", err)
	}

	exifRaw, err := exif.SearchAndExtractExif(imageBytes)
	if err != nil {
		fmt.Printf("No EXIF data.\n")
		os.Exit(1)
	}

	//fmt.Println("raw", string(exifRaw))

	entries, _, err := exif.GetFlatExifDataUniversalSearch(exifRaw, nil, true)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		//if strings.HasPrefix(entry.TagName, "GPS") {
		fmt.Printf("entry %+v\n", entry)
		//fmt.Printf("entry: name=%s  value=%+v T=%T \n", entry.TagName, entry.Value, entry.Value)
		//}
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, exifRaw)

	for _, leaf := range index.Tree {
		fmt.Printf("leaft %+v\n", leaf)
	}

	ifdTag, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	gpsInfo, err := ifdTag.GpsInfo()
	fmt.Printf("index %+v %+v\n", gpsInfo.Latitude.Decimal(), gpsInfo.Longitude.Decimal())
}
