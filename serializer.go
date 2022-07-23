package sbvj

import (
	"bufio"
	"encoding/binary"
	"io"
)

const (
	segmentBits int64 = 0x7f
	continueBit int64 = 0x80
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

func (w *SBVJWriter) PackVarint(value int64) error {
	if err := w.WriteByte(byte(VARINT)); err != nil {
		return err
	}

	return w.writeVarint(value)
}

func (w *SBVJWriter) writeVarint(value int64) error {
	for {
		if (value & ^segmentBits) == 0 {
			return w.WriteByte(byte(value))
		}

		if err := w.WriteByte(byte((value & segmentBits) | continueBit)); err != nil {
			return err
		}
		value >>= 7
	}
}

func (w *SBVJWriter) PackString(s string) error {
	if err := w.WriteByte(byte(STRING)); err != nil {
		return err
	}

	if err := w.writeVarint(int64(len(s))); err != nil {
		return err
	}

	if _, err := w.WriteString(s); err != nil {
		return err
	}

	return nil
}

func (w *SBVJWriter) PackList(l SBVJList) error {
	if err := w.WriteByte(byte(LIST)); err != nil {
		return err
	}

	size := len(l.Items)
	if err := w.writeVarint(int64(size)); err != nil {
		return err
	}

	for _, obj := range l.Items {
		if err := w.PackObject(obj); err != nil {
			return err
		}
	}

	return nil
}
