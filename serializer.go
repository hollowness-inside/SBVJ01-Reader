package sbvj

import (
	"bufio"
	"encoding/binary"
	"io"
)

type SBVJWriter struct {
	*bufio.Writer
}

func NewWriter(w io.Writer) SBVJWriter {
	writer := SBVJWriter{bufio.NewWriter(w)}
	writer.WriteString("SBVJ01")

	return writer
}

func (w *SBVJWriter) PackDouble(d float64) error {
	if err := w.WriteByte(byte(DOUBLE)); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, d)
}

func (w *SBVJWriter) PackBoolean(b bool) error {
	if err := w.WriteByte(byte(BOOLEAN)); err != nil {
		return err
	}

	if b {
		return w.WriteByte(1)
	} else {
		return w.WriteByte(0)
	}
}
