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
		DivisionCode:        "10",
		SubdivisionCode:     "100",
		SubdivisionName:     "WOMEN'S SWEATERS",
		DepartmentCode:      "012",
		DepartmentName:      "SWEATERS",
		ClassCode:           "20",
		ClassName:           "SWEATERS",
		SubclassCode:        "222",
		SubclassName:        "FUZZY (2-4MM)",
		Description:         "A warm fuzzy sweater.",
		LongDescription:     "A warm fuzzy sweater for cold days.",
		Color:               "FUSCHIA SA",
		ColorCode:           "20",
		SupplierCode:        "123456789",
		SupplierName:        "iSWEATR",
		VendorProductNumber: "ELECTRA-LS",
		BundleItemCount:     1,
		StoreDetail: []GTINStoreDetail{GTINStoreDetail{
			StoreNumber: "3",
			Price: StoreDetailPrice{
				RegularRetail:     "120.00",
				SpecialRetail:     "15.00",
				IsClearancePriced: true,
			},
		}},
	}
	return productInfo, nil
}
