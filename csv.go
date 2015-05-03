package main

import (
	"github.com/atotto/encoding/csv"
	"io"
	"os"
)

type CSVGTINSource struct {
	DataPath string
	GTINs    map[string]GTIN
}

func (s *CSVGTINSource) Start() error {
	s.GTINs = make(map[string]GTIN)

	file, err := os.Open(s.DataPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		gtin := GTIN{}
		err2 := reader.ReadStruct(&gtin)
		if err2 == io.EOF {
			break
		} else if err2 != nil {
			return err2
		}
		s.GTINs[gtin.GTIN] = gtin
	}
	return err
}

func (s *CSVGTINSource) Get(gtinId string) (*GTIN, error) {
	gtin := new(GTIN)
	*gtin, _ = s.GTINs[gtinId]
	return gtin, nil
}
