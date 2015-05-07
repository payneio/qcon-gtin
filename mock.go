package main

import (
	"fmt"
)

type MockGTINSource struct {
}

func (s *MockGTINSource) Start() error { return nil }

func (s *MockGTINSource) Get(gtinId string) (*GTIN, error) {
	gtin := gtinId
	if gtin != "123456789012" {
		return &GTIN{GTIN: gtin}, ErrGTINNotFound
	}
	productInfo := &GTIN{
		GTIN:                fmt.Sprintf("0%s", gtin),
		DisplayGTIN:         gtin,
		DisplayGTINType:     "UPC-A",
		SKU:                 "1234567890",
		Size1:               "SMALL",
		Size2:               "4",
		Style:               "0123456",
		OriginalRetail:      "140.00",
		RecommendRetail:     "160.00",
		SubdivisionName:     "WOMEN'S SWEATERS",
		ClassName:           "SWEATERS",
		SubclassName:        "FUZZY (2-4MM)",
		Description:         "A warm fuzzy sweater.",
		Color:               "FUSCHIA SA",
		SupplierName:        "iSWEATR",
		VendorProductNumber: "ELECTRA-LS",
		StoreDetail: []GTINStoreDetail{GTINStoreDetail{
			StoreNumber: "3",
			Price: StoreDetailPrice{
				RegularRetail: "120.00",
				SpecialRetail: "15.00",
			},
		}},
	}
	return productInfo, nil
}
