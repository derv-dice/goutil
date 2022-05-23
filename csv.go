package main

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/jszwec/csvutil"
)

var ErrWriterCSVNotInited = errors.New("WriterCSV not initialized. It's important to use NewWriterCSV()")

type WriterCSV struct {
	init bool
	st   interface{}

	out *os.File
	enc *csvutil.Encoder
	wr  *csv.Writer
}

func NewWriterCSV(filename string, structType interface{}, headers bool) (wrCSV *WriterCSV, err error) {
	wrCSV = new(WriterCSV)
	if wrCSV.out, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return
	}

	wrCSV.st = structType
	wrCSV.wr = csv.NewWriter(wrCSV.out)
	wrCSV.enc = csvutil.NewEncoder(wrCSV.wr)
	wrCSV.init = true

	if headers {
		if err = wrCSV.writeHeaders(); err != nil {
			return
		}
	}

	return
}

func (w *WriterCSV) writeHeaders() (err error) {
	if !w.init {
		return ErrWriterCSVNotInited
	}

	if err = w.enc.EncodeHeader(w.st); err != nil {
		return
	}

	w.wr.Flush()
	return
}

func (w *WriterCSV) WriteRow(v interface{}) (err error) {
	if !w.init {
		return ErrWriterCSVNotInited
	}

	if err = w.enc.Encode(v); err != nil {
		return
	}

	w.wr.Flush()
	return
}

func (w *WriterCSV) Close() (err error) {
	if !w.init {
		return ErrWriterCSVNotInited
	}

	return w.out.Close()
}

func UnmarshalCSV(data []byte, dest interface{}) (err error) {
	return csvutil.Unmarshal(data, dest)
}

func MarshalCSV(v interface{}) ([]byte, error) {
	return csvutil.Marshal(v)
}
