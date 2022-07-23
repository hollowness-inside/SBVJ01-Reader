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