package cloudfront

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Parser struct {
	delimiter rune // default: '\t'
}

func NewParser() *Parser {
	return &Parser{delimiter: '\t'}
}

func (p *Parser) ParseFile(in *os.File) ([]*Line, error) {
	var lines []*Line

	r := csv.NewReader(in)
	r.Comma = '\t'
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		line, err := NewLine(record)

		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, line)
	}

	return lines, nil
}
