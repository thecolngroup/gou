package csv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	Name       string
	Properties map[string]any
}

func TestWriteToCSV(t *testing.T) {

	data := []MyStruct{
		{
			Name:       "Row 1",
			Properties: map[string]any{"Param 1": "0"},
		},
		{
			Name:       "Row 2",
			Properties: map[string]any{"Param 1": 0},
		},
	}

	encMap := func(m map[string]any) ([]byte, error) {
		return []byte(fmt.Sprint(m)), nil
	}

	err := WriteToCSV("testdata/writetocsv_test.csv", data, encMap)
	assert.NoError(t, err)
}
