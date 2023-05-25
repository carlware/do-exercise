package extractor

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"io/ioutil"
	"log"
	"os"
)

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

func ExtractExifDataFromImages(imageByName map[string][]byte) ([]FileInfo, error) {
	imagesInfo := make([]FileInfo, 0, len(imageByName))
	for imageName, data := range imageByName {
		exifRaw, err := exif.SearchAndExtractExif(data)
		if err != nil {
			log.Printf("%s: No EXIF data\n", imageName)
			continue
		}
		im, err := exifcommon.NewIfdMappingWithStandard()
		ti := exif.NewTagIndex()

		_, index, err := exif.Collect(im, ti, exifRaw)
		imagesInfo = append(imagesInfo, FileInfo{
			Filename: imageName,
			IfdIndex: index,
		})
	}

	return imagesInfo, nil
}
