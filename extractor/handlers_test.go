package extractor

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileName(t *testing.T) {
	info := FileInfo{
		Filename: "foo.jpg",
	}

	name, err := GetFileName(info)
	assert.NoError(t, err)
	assert.Equal(t, "foo.jpg", name)
}

func getIfdIndex(t *testing.T, filename string) exif.IfdIndex {
	t.Helper()

	rawExif, err := exif.SearchFileAndExtractExif(fmt.Sprintf("assets/%s", filename))
	assert.NoError(t, err)

	im, err := exifcommon.NewIfdMappingWithStandard()
	assert.NoError(t, err)

	ti := exif.NewTagIndex()
	_, index, err := exif.Collect(im, ti, rawExif)
	assert.NoError(t, err)

	return index
}

func TestGetLatitude(t *testing.T) {
	fileWithGPS := FileInfo{
		Filename: "gps.jpg",
		IfdIndex: getIfdIndex(t, "gps.jpg"),
	}
	fileWithoutGPS := FileInfo{
		Filename: "NDM_8901.jpg",
		IfdIndex: getIfdIndex(t, "NDM_8901.jpg"),
	}

	t.Run("should get the latitude", func(t *testing.T) {
		latitude, err := GetLatitude(fileWithGPS)
		assert.NoError(t, err)
		assert.Equal(t, "26.586667", latitude)
	})

	t.Run("should return error if file doesn't contain latitude", func(t *testing.T) {
		latitude, err := GetLatitude(fileWithoutGPS)
		assert.Error(t, err)
		assert.Equal(t, "", latitude)
	})

	t.Run("should get the longitude", func(t *testing.T) {
		longitude, err := GetLongitude(fileWithGPS)
		assert.NoError(t, err)
		assert.Equal(t, "26.586667", longitude)
	})

	t.Run("should return error if file doesn't contain longitude", func(t *testing.T) {
		longitude, err := GetLongitude(fileWithoutGPS)
		assert.Error(t, err)
		assert.Equal(t, "", longitude)
	})
}
