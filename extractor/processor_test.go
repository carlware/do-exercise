package extractor

import (
	"bytes"
	"testing"

	"geoextractor-go/extractor/writers"
	"github.com/stretchr/testify/assert"
)

func TestTagProcessor_AddTagHandler(t *testing.T) {
	processor := NewTagProcessor([]FileInfo{})
	processor = processor.AddTagHandler("myname", func(info FileInfo) (string, error) {
		return "", nil
	})

	assert.Equal(t, []string{"myname"}, processor.columns)
	assert.Len(t, processor.processorRows, 1)
}

func TestTagProcessor_Write(t *testing.T) {
	processor := NewTagProcessor([]FileInfo{
		{
			Filename: "foo.jpg",
		},
	})
	processor = processor.AddTagHandler("name", func(info FileInfo) (string, error) {
		return info.Filename, nil
	})

	var b bytes.Buffer
	csvWriter := writers.NewCsvWriter(&b)

	err := processor.Write(csvWriter)
	assert.NoError(t, err)
	assert.Equal(t, "name\nfoo.jpg\n", b.String())
}
