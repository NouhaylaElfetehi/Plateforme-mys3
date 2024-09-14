package entity

import (
	"encoding/xml"
)

type Bucketobject struct {
	XMLName      xml.Name         `xml:"BucketObject"`
	Key          string           `xml:"Key"`
	LastModified string           `xml:"LastModified"`
	ETag         string           `xml:"ETag"`
	Size         int64            `xml:"Size"`
	StorageClass StorageClassType `xml:"StorageClass"`
	Owner        Owner            `xml:"Owner"`
	Type         string           `xml:"Type"`
	URI          string           `xml:"URI"`
	BucketName   string           `xml:"BucketName"`
}
