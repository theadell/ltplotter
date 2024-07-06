package parser

import (
	"io"

	"github.com/theadell/ltspice"
)

func Parse(reader io.Reader) (*ltspice.SimData, error) {
	return ltspice.ParseFromReader(reader)
}
