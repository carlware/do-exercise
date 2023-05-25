package processors

import (
	"fmt"
	"geoextractor-go/extractor"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

func GetLatitude(info extractor.FileInfo) (string, error) {
	ifdTag, err := info.IfdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return "", err
	}
	gpsInfo, err := ifdTag.GpsInfo()
	if err != nil {
		return "", err
	}
	if gpsInfo.Latitude.Decimal() == 0 {
		return "N/A", nil
	}
	return fmt.Sprintf("%f", gpsInfo.Latitude.Decimal()), nil
}

func GetLongitude(info extractor.FileInfo) (string, error) {
	ifdTag, err := info.IfdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return "", err
	}
	gpsInfo, err := ifdTag.GpsInfo()
	if err != nil {
		return "", err
	}
	if gpsInfo.Latitude.Decimal() == 0 {
		return "N/A", nil
	}
	return fmt.Sprintf("%f", gpsInfo.Latitude.Decimal()), nil
}