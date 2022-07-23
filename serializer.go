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

type SBVJOptions struct {
	Name      string
	Versioned bool
	Version   int32
}

func NewWriter(w io.Writer, opt *SBVJOptions) (*SBVJWriter, error) {
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
	err := w.WriteByte(byte(NIL))
	return err
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

	return w.WriteBoolean(b)
}

func (w *SBVJWriter) PackVarint(value int64) error {
	if err := w.WriteByte(byte(VARINT)); err != nil {
		return err
	}

	return w.WriteVarint(value)
}

func (w *SBVJWriter) PackString(s string) error {
	if err := w.WriteByte(byte(STRING)); err != nil {
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

func (w *SBVJWriter) PackList(l SBVJList) error {
	if err := w.WriteByte(byte(LIST)); err != nil {
		return err
	}

	size := len(l.Items)
	if err := w.WriteVarint(int64(size)); err != nil {
		return err
	}

	for _, obj := range l.Items {
		if err := w.PackObject(obj); err != nil {
			return err
		}
	}

	return nil
}

func (w *SBVJWriter) PackMap(m SBVJMap) error {
	if err := w.WriteByte(byte(MAP)); err != nil {
		return err
	}

	size := len(m.Items)
	if err := w.WriteVarint(int64(size)); err != nil {
		return err
	}

	for _, pair := range m.Items {
		if err := w.PackString(pair.Key); err != nil {
			return err
		}

		if err := w.PackObject(pair.Value); err != nil {
			return err
		}
	}

	return nil
}

func (w *SBVJWriter) PackObject(o SBVJObject) (err error) {
	switch o.Type {
	case NIL:
		w.PackNil()
	case DOUBLE:
		v := o.Value.(float64)
		err = w.PackDouble(v)
	case BOOLEAN:
		v := o.Value.(bool)
		err = w.PackBoolean(v)
	case VARINT:
		v := o.Value.(int64)
		err = w.PackVarint(v)
	case STRING:
		v := o.Value.(string)
		err = w.PackString(v)
	case LIST:
		v := o.Value.(SBVJList)
		err = w.PackList(v)
	case MAP:
		v := o.Value.(SBVJMap)
		err = w.PackMap(v)
	}

	return err
}
