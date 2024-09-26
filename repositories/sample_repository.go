package repositories

import (
	"ticket-groupg-gin/models"
	"gorm.io/gorm"
)

type SampleRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *SampleRepository {
    return &SampleRepository{DB: db}
}

func (r *SampleRepository) GetAll() ([]models.Sample, error) {
    var samples []models.Sample
    if err := r.DB.Find(&samples).Error; err != nil {
        return nil, err
    }
    return samples, nil
}

func (r *SampleRepository) Create(sample models.Sample) (models.Sample, error) {
    if err := r.DB.Create(&sample).Error; err != nil {
        return models.Sample{}, err
    }
    return sample, nil
}
