package main

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
	"log"
	"reflect"
)

type DynamoGTINSource struct {
	AccessKey string
	SecretKey string
	creds     aws.CredentialsProvider
	Client    *dynamodb.DynamoDB
}

func (dyn *DynamoGTINSource) Start() error {
	dyn.Client = dynamodb.New(&aws.Config{Region: "us-west-2"})
	return nil
}

func (dyn *DynamoGTINSource) Get(gtinId string) (*GTIN, error) {

	gtin := &GTIN{GTIN: gtinId}

	// Get SKU from GTIN
	req := &dynamodb.GetItemInput{
		TableName: aws.String("gtins"),
		Key: &map[string]*dynamodb.AttributeValue{
			"GTIN": &dynamodb.AttributeValue{S: aws.String(gtinId)},
		},
	}
	resp, err := dyn.Client.GetItem(req)
	if err != nil {
		return gtin, err
	} else if resp.Item == nil {
		return gtin, ErrGTINNotFound
	}
	sku := awsutil.StringValue(resp)
	if sku == "" {
		return new(GTIN), ErrGTINNotFound
	}
	log.Printf("SKU:%v\n", sku)

	// Now get SKU product info
	req = &dynamodb.GetItemInput{
		TableName: aws.String("skus"),
		Key: &map[string]*dynamodb.AttributeValue{
			"SKU": &dynamodb.AttributeValue{S: aws.String(sku)},
		},
	}
	resp, err = dyn.Client.GetItem(req)
	if err != nil {
		return gtin, err
	} else if resp.Item == nil {
		return gtin, ErrGTINNotFound
	}

	// We have our data, load up the GTIN
	item := resp.Item
	productInfo := buildProductInfo(item)
	productInfo.GTIN = gtinId
	productInfo.BundleItemCount = 1

	return productInfo, nil

}

func buildProductInfo(item *map[string]*dynamodb.AttributeValue) *GTIN {
	keys := []string{"SKU", "Size1", "Size2", "Style", "OriginalRetail", "RecommendRetail",
		"DivisionCode", "SubdivisionName", "SubdivisionName", "DepartmentCode", "DepartmentName",
		"ClassCode", "ClassName", "SubclassCode", "SubclassName", "Description", "LongDescription",
		"Color", "ColorCode", "SupplierCode", "SupplierName", "VendorProductName"}
	p := &GTIN{}
	for _, key := range keys {
		itm := *item
		if _, ok := itm[key]; ok {
			reflect.ValueOf(p).Elem().FieldByName(key).SetString(awsutil.StringValue(itm))
		}
	}
	// StoreDetail: []GTINStoreDetail{GTINStoreDetail{
	// StoreNumber: "3",
	// Price: StoreDetailPrice{
	// RegularRetail:     "120.00",
	// SpecialRetail:     "15.00",
	// IsClearancePriced: true,
	// },
	// }},
	return p
}
