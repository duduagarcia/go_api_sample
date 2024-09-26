package services

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type SampleService struct {
    Repository *repositories.SampleRepository
}

func NewSampleService(repo *repositories.SampleRepository) *SampleService {
    return &SampleService{Repository: repo}
}

func (sampleService *SampleService) GetAllSamples() ([]models.Sample, error) {
    return sampleService.Repository.GetAll()
}

func (sampleService *SampleService) CreateSample(sample models.Sample) (models.Sample, error) {
    return sampleService.Repository.Create(sample)
}
