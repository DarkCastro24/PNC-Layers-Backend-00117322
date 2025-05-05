package csv

import (
	"encoding/csv"
	"os"
)

func ReadAllFromFile(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return csv.NewReader(f).ReadAll()
}

func WriteAllToFile(path string, records [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	w := csv.NewWriter(f)
	defer w.Flush()
	return w.WriteAll(records)
}
