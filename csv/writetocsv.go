// Package csv provides helpers to write CSV files.
package csv

import (
	"encoding/csv"
	"os"

	"github.com/jszwec/csvutil"
)

// CSVEncoder is a custom encoder for marshalling types to CSV files.
type CSVEncoder[T any] func(t T) ([]byte, error)

// WriteToCSV writes the given data to a CSV file, optionally using the given encoders to marshal custom types.
// Will write a header row if the given file is empty.
// Wraps github.com/jszwec/csvutil
func WriteToCSV(filename string, data interface{}, encoder ...any) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	// nolint:errcheck // Explicit err check via sync at end of function
	defer f.Close()

	// Wrap file in CSV struct encoder
	w := csv.NewWriter(f)
	enc := csvutil.NewEncoder(w)
	enc.Tag = "csv"
	for _, e := range encoder {
		enc.Register(e)
	}

	// Only write header if new file
	info, err := f.Stat()
	if err != nil {
		return err
	}
	if info.Size() > 0 {
		enc.AutoHeader = false
	}

	// Write report
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	// Flush and check for errors
	w.Flush()
	if w.Error() != nil {
		return w.Error()
	}
	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
