package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

// ReadFileLineByLine - reads file line by line in string slice
func ReadFileLineByLine(filename string) (lines []string, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}

	defer func() {
		cErr := file.Close()
		if cErr != nil {
			Log().Warn().Err(cErr).Send()
		}
	}()

	s := bufio.NewScanner(file)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	err = s.Err()
	return
}

// WriteCSV - Создает и заполняет указанный .csv файл. Если заголовки не нужны, то указать headers = nil
func WriteCSV(filename string, headers []string, records [][]string) (err error) {
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		return
	}

	defer func() {
		cErr := file.Close()
		if cErr != nil {
			Log().Warn().Err(cErr).Send()
		}
	}()

	w := csv.NewWriter(file)
	defer w.Flush()

	if headers != nil {
		if err = w.Write(headers); err != nil {
			return
		}
	}

	for i := range records {
		if err = w.Write(records[i]); err != nil {
			return
		}
	}

	return
}
