package models

import (
	entity "api-interface/entities"
	repository "api-interface/repositories"
)

type BucketObjectModel struct {
    bucketObjectRepository *repository.QueryBuilder[*entity.BucketObject]
}

// NewBucketModel initialise un BucketObjectModel avec le repository approprié
func UseBucketObjectModel() (*BucketObjectModel, error) {
    queryBuilder, err := UseRepository[*entity.BucketObject]("BucketObject")
    if err != nil {
        return nil, err
    }

    return &BucketObjectModel{
        bucketObjectRepository: queryBuilder,
    }, nil
}

func (bm *BucketObjectModel) Insert(bucket *entity.BucketObject) error {
	return bm.bucketObjectRepository.Insert(bucket)
}

func (bm *BucketObjectModel) GetAllBucketObjects() ([]entity.BucketObject, error) {
    buckets, err := bm.bucketObjectRepository.Find(func(b *entity.BucketObject) bool {
        return true
    })
    if err != nil {
        return nil, err
    }

    // Convertir []*entity.Bucket en []entity.Bucket
    var result []entity.BucketObject
    for _, bucket := range buckets {
        result = append(result, *bucket)
    }
    return result, nil
}

func (bm *BucketObjectModel) GetBucketByName(name string) (*entity.BucketObject, error) {
    bucket := new(entity.BucketObject)
    err := bm.bucketObjectRepository.Get(name, bucket)
    if err != nil {
        return nil, err
    }
    return bucket, nil
}

