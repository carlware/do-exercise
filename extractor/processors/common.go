package processors

import "geoextractor-go/extractor"

func GetFileName(info extractor.FileInfo) (string, error) {
	return info.Filename, nil
}
