package bucketdtos

import (
	"encoding/xml"
	"strings"
)

const (
    SingleAvailabilityZone = "SingleAvailabilityZone"
    Directory              = "Directory"
	AvailabilityZone 	 = "AvailabilityZone"
)

type CreateBucketConfiguration struct {
	XMLName            xml.Name `xml:"CreateBucketConfiguration"`
	XMLNS              string   `xml:"xmlns,attr"`
	LocationConstraint *string   `xml:"LocationConstraint"`
	Location           *LocationInfo `xml:"Location"`
	Bucket             *BucketInfo   `xml:"Bucket"`
}

type LocationInfo struct {
	Name string `xml:"Name"`
	Type string `xml:"Type"`
}

type  BucketInfo struct {
	DataRedundancy string `xml:"DataRedundancy"`
	Type           string `xml:"Type"`
}

// Fonction de validation pour DataRedundancy
func IsValidDataRedundancy(value string) bool {
    return value == SingleAvailabilityZone
}

// Fonction de validation pour Type
func IsValidBucketType(value string) bool {
    return value == Directory
}

func IsValidLocationType(value string) bool {
	return value == AvailabilityZone
}

/**
 * Fonction de validation pour LocationConstraint
       AfSouth1, ApEast1, ApNortheast1, ApNortheast2, ApNortheast3,
        ApSouth1, ApSouth2, ApSoutheast1, ApSoutheast2, ApSoutheast3,
        CaCentral1, CnNorth1, CnNorthwest1, EU, EuCentral1, EuNorth1,
        EuSouth1, EuSouth2, EuWest1, EuWest2, EuWest3, MeSouth1,
        SaEast1, UsEast2, UsGovEast1, UsGovWest1, UsWest1, UsWest2,
		*/
func IsValidLocationConstraint(value string) bool {
	trimmedValue := strings.TrimSpace(value)
	validValues := []string{
		"AfSouth1", "ApEast1", "ApNortheast1", "ApNortheast2", "ApNortheast3",
		"ApSouth1", "ApSouth2", "ApSoutheast1", "ApSoutheast2", "ApSoutheast3",
		"CaCentral1", "CnNorth1", "CnNorthwest1", "EU", "EuCentral1", "EuNorth1",
		"EuSouth1", "EuSouth2", "EuWest1", "EuWest2", "EuWest3", "MeSouth1",
		"SaEast1", "UsEast2", "UsGovEast1", "UsGovWest1", "UsWest1", "UsWest2",
	}
	for _, validValue := range validValues {
		if trimmedValue == validValue {
			return true
		}
	}
	return false
}