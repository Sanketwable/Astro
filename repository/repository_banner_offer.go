package repository

import "astro/models"

type BannerOfferRepository interface {
	Save([]models.Data) ([]models.Data, error)
	FindById(int64) ([]models.Data, error)
}
