package sbvj

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/types"
)

const (
	segmentBits int64 = 0x7f
	continueBit int64 = 0x80
)

type SBVJWriter struct {
	*bufio.Writer
}

type SBVJOptions struct {
	Name      string
	Versioned bool
	Version   int32
}

func NewWriter(w io.Writer, opt *SBVJOptions) (*SBVJWriter, error) {
	if opt == nil {
		opt = &SBVJOptions{
			Name:      "",
			Versioned: false,
			Version:   0,
		}
	}

	writer := SBVJWriter{bufio.NewWriter(w)}
	if _, err := writer.WriteString("SBVJ01"); err != nil {
		return nil, err
	}

	if err := writer.WriteVarint(int64(len(opt.Name))); err != nil {
		return nil, err
	}

	if _, err := writer.WriteString(opt.Name); err != nil {
		return nil, err
	}

	if err := writer.WriteBoolean(opt.Versioned); err != nil {
		return nil, err
	}

	if opt.Versioned {
		if err := binary.Write(writer, binary.BigEndian, opt.Version); err != nil {
			return nil, err
		}
	}

	return &writer, nil
}

func (w *SBVJWriter) WriteVarint(value int64) error {
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

func (w *SBVJWriter) WriteBoolean(b bool) error {
	if b {
		return w.WriteByte(1)
	} else {
		return w.WriteByte(0)
	}
}

func (w *SBVJWriter) PackNil() error {
	err := w.WriteByte(byte(types.NIL))
	return err
}

func (w *SBVJWriter) PackDouble(d float64) error {
	if err := w.WriteByte(byte(types.DOUBLE)); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, d)
}

func (w *SBVJWriter) PackBoolean(b bool) error {
	if err := w.WriteByte(byte(types.BOOLEAN)); err != nil {
		return err
	}

	return w.WriteBoolean(b)
}

func (w *SBVJWriter) PackVarint(value int64) error {
	if err := w.WriteByte(byte(types.VARINT)); err != nil {
		return err
	}

	return w.WriteVarint(value)
}

func (w *SBVJWriter) PackString(s string) error {
	if err := w.WriteByte(byte(types.STRING)); err != nil {
		return err
	}

	if err := w.WriteVarint(int64(len(s))); err != nil {
		return err
	}

	if _, err := w.WriteString(s); err != nil {
		return err
	}

	return nil
}

func (w *SBVJWriter) PackList(l types.SBVJList) error {
	if err := w.WriteByte(byte(types.LIST)); err != nil {
		return err
	}

	size := len(l)
	if err := w.WriteVarint(int64(size)); err != nil {
		return err
	}

	for _, obj := range l {
		if err := w.PackObject(obj); err != nil {
			return err
		}
	}

	return nil
}

func (w *SBVJWriter) PackMap(m types.SBVJMap) error {
	if err := w.WriteByte(byte(types.MAP)); err != nil {
		return err
	}

	size := len(m)
	if err := w.WriteVarint(int64(size)); err != nil {
		return err
	}

	for key, value := range m {
		if err := w.PackString(key); err != nil {
			return err
		}

		if err := w.PackObject(value); err != nil {
			return err
		}
	}

	return nil
}

func (w *SBVJWriter) PackObject(o types.SBVJObject) (err error) {
	switch o.Type {
	case types.NIL:
		w.PackNil()
	case types.DOUBLE:
		v := o.Value.(float64)
		err = w.PackDouble(v)
	case types.BOOLEAN:
		v := o.Value.(bool)
		err = w.PackBoolean(v)
	case types.VARINT:
		v := o.Value.(int64)
		err = w.PackVarint(v)
	case types.STRING:
		v := o.Value.(string)
		err = w.PackString(v)
	case types.LIST:
		v := o.Value.(types.SBVJList)
		err = w.PackList(v)
	case types.MAP:
		v := o.Value.(types.SBVJMap)
		err = w.PackMap(v)
	}

	return err
}
