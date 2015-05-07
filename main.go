package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var csvData = false
var dynamoData = false

type GTINSource interface {
	Start() error
	Get(string) (*GTIN, error)
}

var gtins map[int]GTIN

func getopt(name string, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}

func main() {

	port := flag.String("p", "80", "Port")
	csvPath := flag.String("c", "", "Path to a csv file to use for GTIN data.")
	mockData := flag.Bool("m", false, "Return mock GTIN data.")
	flag.Parse()

	var source GTINSource

	if *mockData {
		source = &MockGTINSource{}
	} else if *csvPath != "" {
		// source = &CSVGTINSource{}
		// src.DataPath = *csvPath
	} else {
		// Dynamo by default
		source = &DynamoGTINSource{}
		src := source.(*DynamoGTINSource)
		src.AccessKey = getopt("AWS_ACCESS_KEY", "")
		src.SecretKey = getopt("AWS_SECRET_KEY", "")
	}

	err := source.Start()
	if err != nil {
		log.Fatalf("Could not initialize data source: %v\n", err)
	}

	r := gin.Default()

	r.GET("/stores/:storeId/gtins/:gtin", func(c *gin.Context) {

		storeId := c.Params.ByName("storeId")
		gtinParam := c.Params.ByName("gtin")

		if storeId == "health" && gtinParam == "check" {
			c.String(200, "Healthy")
			return
		}

		gtin, err := source.Get(gtinParam)
		if err != nil {
			if err == ErrGTINNotFound {
				c.String(404, err.Error())
			} else {
				c.String(500, err.Error())
			}
			return
		}
		if gtin.GTIN == "" {
			c.String(404, "GTIN not found.")
			return
		}

		gtin.GTIN = gtinParam
		gtin.DisplayGTIN = gtinParam
		gtin.DisplayGTINType = "UPC-A"
		if len(gtinParam) == 13 {
			gtin.DisplayGTINType = "EAN-13"
		}

		c.JSON(200, gtin)
	})
	r.Run(fmt.Sprintf(":%s", *port))
}

var ErrGTINNotFound = errors.New("GTIN not found.")

type StoreDetailPrice struct {
	RegularRetail string `json:"regularRetail"`
	SpecialRetail string `json:"specialRetail"`
}

type GTINStoreDetail struct {
	StoreNumber string           `json:"storeNumber"`
	Price       StoreDetailPrice `json:"price"`
}

type GTIN struct {
	GTIN                string            `json:"gtin"`
	DisplayGTIN         string            `json:"displayGTIN"`
	DisplayGTINType     string            `json:"displayGTINType"` // UPC-A or EAN-13
	SKU                 string            `json:"sku"`
	Size1               string            `json:"size1"`
	Size2               string            `json:"size2"`
	Style               string            `json:"style"`
	OriginalRetail      string            `json:"originalRetail"`
	RecommendRetail     string            `json:"recommendRetail"`
	SubdivisionName     string            `json:"subdivisionName"`
	ClassName           string            `json:"className"`
	SubclassName        string            `json:"subclassName"`
	Description         string            `json:"description"`
	Color               string            `json:"color"`
	SupplierName        string            `json:"supplierName"`
	VendorProductNumber string            `json:"vendorProductNumber"`
	StoreDetail         []GTINStoreDetail `json:"storeDetail"`
}
