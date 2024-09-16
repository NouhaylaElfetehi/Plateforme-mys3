package entity

import (
	"encoding/json"
)

type Bucket struct {
    Name         string          `xml:"name" json:"name"`
    CreationDate *string          `xml:"creationDate" json:"creationDate"`
    Owner        Owner           `xml:"owner" json:"owner"`
    URI          string          `xml:"uri" json:"uri"`
    Type         BucketType      `xml:"type" json:"type"`
    StorageClass *StorageClassType `xml:"storageClass" json:"storageClass"`
    Versioning   VersioningStatus `xml:"versioning" json:"versioning"`
    ObjectCount  *int64           `xml:"objectCount" json:"objectCount"`
    Size         *int64           `xml:"size" json:"size"`
    LastModified *string          `xml:"lastModified" json:"lastModified"`
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