package extractor

import (
	"fmt"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

func GetFileName(info FileInfo) (string, error) {
	return info.Filename, nil
}

func GetLatitude(info FileInfo) (string, error) {
	ifdTag, err := info.IfdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return "", err
	}
	gpsInfo, err := ifdTag.GpsInfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", gpsInfo.Latitude.Decimal()), nil
}

func GetLongitude(info FileInfo) (string, error) {
	ifdTag, err := info.IfdIndex.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return "", err
	}
	gpsInfo, err := ifdTag.GpsInfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", gpsInfo.Latitude.Decimal()), nil
}
