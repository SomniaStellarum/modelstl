package modelstl

import (
	"encoding/binary"
)

type Writer struct {
	header []byte
	num    uint32
	w      *bufio.Writer
}

func NewWriter(w io.Writer, header []byte, num uint32) *Writer (
	writer := &Writer {
		header: header,
		num:    num,
		w:      bufio.NewWriter(w),
	}
	
)
