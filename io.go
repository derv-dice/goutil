package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"os"
)

// ReadFileLineByLine - reads file line by line in string slice
func ReadFileLineByLine(filename string, unique bool) (lines []string, err error) {
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

	var uniqMap map[string]bool
	if unique {
		uniqMap = map[string]bool{}
	}

	for s.Scan() {
		if unique {
			hash := hashFromString(s.Text())

			if uniqMap[hash] {
				continue
			}

			uniqMap[hash] = true
		}

		lines = append(lines, s.Text())
	}

	err = s.Err()
	return
}

func hashFromString(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
