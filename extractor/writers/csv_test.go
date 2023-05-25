package writers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsvWriter(t *testing.T) {
	record := []string{"a", "b", "c"}

	var b bytes.Buffer
	csvWriter := NewCsvWriter(&b)

	err := csvWriter.Write(record)
	assert.NoError(t, err)
	csvWriter.Flush()

	assert.NoError(t, err)
	assert.Equal(t, "a,b,c\n", b.String())
}
