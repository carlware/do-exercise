package extractor

import (
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type GPS struct {
	Degrees float64
	Minutes float64
	Seconds float64
	Direction string
}

func rationalToFloat(rational exifcommon.Rational) float64 {
	if rational.Denominator == 0 {
		return float64(rational.Numerator)
	}
	return float64(rational.Numerator) / float64(rational.Denominator)
}

func ConvertToGPS(degrees, minutes, seconds exifcommon.Rational, direction string) GPS {
	return GPS{
		Degrees:   rationalToFloat(degrees),
		Minutes:   rationalToFloat(minutes),
		Seconds:   rationalToFloat(seconds),
		Direction: direction,
	}
}

func DegreesToDecimal(gps GPS) float64 {
	decimal := gps.Degrees + gps.Minutes/60 + gps.Seconds/3600

	if gps.Direction == "S" || gps.Direction == "W" {
		decimal = -decimal
	}

	return decimal
}

func ExtractLongitude(longitude, longitudeRef exif.ExifTag) {

}