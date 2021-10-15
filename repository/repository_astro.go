package repository

import "astro/models"

type AstroRepository interface {
	Save(models.Astro) (int, error)
	FindById(uint64) (models.Astro, error)
}
