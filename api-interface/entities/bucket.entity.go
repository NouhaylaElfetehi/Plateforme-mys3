package entity

import (
	"encoding/json"
)

type Bucket struct {
    Name         string          `xml:"name"`
    CreationDate *string          `xml:"creationDate"`
    Owner        Owner           `xml:"owner"`
    URI          string          `xml:"uri"`
    Type         BucketType      `xml:"type"`
    StorageClass *StorageClassType `xml:"storageClass"`
    Versioning   VersioningStatus `xml:"versioning"`
    ObjectCount  *int64           `xml:"objectCount"`
    Size         *int64           `xml:"size"`
    LastModified *string          `xml:"lastModified"`
}

type CreateBucketRequestStruct struct {
    Name         string          `xml:"name"`
    Owner        Owner           `xml:"owner"`
    Type         BucketType      `xml:"type"`
    Versioning   VersioningStatus `xml:"versioning"`
}

func (b *Bucket) GetKey() string {
    return b.Name
}

func (b *Bucket) Serialize() ([]byte, error) {
    return json.Marshal(b)
}

func (b *Bucket) Deserialize(data []byte) error {
    return json.Unmarshal(data, b)
}