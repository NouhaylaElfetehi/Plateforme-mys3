package s3_operations

import (
	"api-interface/models"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ListBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	KeyCount    int      `xml:"KeyCount"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	Contents    []struct {
		Key          string `xml:"Key"`
		LastModified string `xml:"LastModified"`
		ETag         string `xml:"ETag"`
		Size         int    `xml:"Size"`
		StorageClass string `xml:"StorageClass"`
	} `xml:"Contents"`
	CommonPrefixes []struct {
		Prefix string `xml:"Prefix"`
	} `xml:"CommonPrefixes"`
}

type S3ObjectLister struct {
	client             *s3.Client
	awsAccessKeyId     string
	awsSecretAccessKey string
	baseURL            string
}

func NewS3ObjectLister(client *s3.Client, accessKeyId, secretAccessKey string, baseURL string) *S3ObjectLister {
	return &S3ObjectLister{
		client:             client,
		awsAccessKeyId:     accessKeyId,
		awsSecretAccessKey: secretAccessKey,
		baseURL:            baseURL,
	}
}

func (lister *S3ObjectLister) ListObjects(ctx context.Context, bucketName string, prefix string, delimiter string, maxKeys int32) ([]models.BucketObject, error) {
	// Construire l'URL de la requête
	baseURL := fmt.Sprintf("%s/%s?list-type=2", lister.baseURL, bucketName)
	query := url.Values{}
	query.Add("list-type", "2")
	if prefix != "" {
		query.Add("prefix", prefix)
	}
	if delimiter != "" {
		query.Add("delimiter", delimiter)
	}
	if maxKeys > 0 {
		query.Add("max-keys", fmt.Sprintf("%d", maxKeys))
	}
	requestURL := baseURL + "?" + query.Encode()

	// Créer la chaîne à signer (StringToSign)
	httpMethod := "GET"
	canonicalURI := "/" + bucketName + "/"
	canonicalQuery := query.Encode()
	canonicalResource := canonicalURI
	if canonicalQuery != "" {
		canonicalResource += "?" + canonicalQuery
	}

	date := time.Now().UTC().Format(time.RFC1123)
	stringToSign := fmt.Sprintf("%s\n\n\n%s\n%s", httpMethod, "", "", date, canonicalResource)

	// Calculer la signature
	signature := calculateSignature(stringToSign, lister.awsSecretAccessKey)

	// Créer la requête HTTP
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	// Ajouter les en-têtes nécessaires
	req.Header.Add("Authorization", fmt.Sprintf("AWS %s:%s", lister.awsAccessKeyId, signature))
	req.Header.Add("x-amz-date", date)

	// Envoyer la requête
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Lire la réponse XML
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Déchiffrer le XML
	var result ListBucketResult
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// Convertir en BucketObject
	var bucketObjects []models.BucketObject
	for _, obj := range result.Contents {
		bucketObject := models.BucketObject{
			Key:          obj.Key,
			LastModified: obj.LastModified,
			ETag:         obj.ETag,
			Size:         int64(obj.Size),
			StorageClass: models.StorageClassType(obj.StorageClass),
			Type:         "OBJECT",
			URI:          obj.Key,
			BucketName:   bucketName,
		}
		bucketObjects = append(bucketObjects, bucketObject)
	}

	return bucketObjects, nil
}

func calculateSignature(stringToSign, secretKey string) string {
	key := []byte(secretKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}
