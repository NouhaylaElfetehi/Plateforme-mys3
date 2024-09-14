package entity

import (
	"encoding/xml"
)

type UploadFileRequest struct {
	XMLName     xml.Name `xml:"UploadFileRequest"`
	FileName    string   `xml:"FileName"`
	FileContent string   `xml:"FileContent"`
}

// Serialize convertit une requête UploadFileRequest en XML.
func (ufr *UploadFileRequest) Serialize() ([]byte, error) {
	return xml.Marshal(ufr)
}

// Deserialize convertit un flux XML en une requête UploadFileRequest.
func (ufr *UploadFileRequest) Deserialize(data []byte) error {
	return xml.Unmarshal(data, ufr)
}
