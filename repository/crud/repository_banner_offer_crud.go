package crud

import (
	"astro/models"
	"database/sql"
)

type repositoryBannerOfferCRUD struct {
	db *sql.DB
}

//NewRepositoryBasicInfoCRUD is
func NewrepositoryBannerOfferCRUD(db *sql.DB) *repositoryBannerOfferCRUD {
	return &repositoryBannerOfferCRUD{db}
}

func (r *repositoryBannerOfferCRUD) Save(bannerOffers []models.Data) ([]models.Data, error) {
	for i, element := range bannerOffers {
		bannerData, err := r.db.Exec("INSERT INTO BannerData (name, landingUrl, orderType, productCode, productName, professionalSlug) VALUES (?, ?, ?, ?, ?, ?)", element.Name, element.LandingUrl, element.OrderType, element.MetaData.ProductCode, element.MetaData.ProductName, element.MetaData.ProfessionalSlug)
		if err != nil {
			return bannerOffers, err
		}
		ppid, err := bannerData.LastInsertId()
		bannerOffers[i].ID = int(ppid)
		if err != nil {
			return bannerOffers, err
		}
		_, err = r.db.Exec("INSERT INTO ImageInfo (pid, id, imageUrl, imageSize, imageType) VALUES (?, ?, ?, ?, ?)", ppid, bannerOffers[i].Image.Large.ID, bannerOffers[i].Image.Large.ImageURL, "medium", "banner")
		if err != nil {
			return bannerOffers, err
		}
		_, err = r.db.Exec("INSERT INTO ImageInfo (pid, id, imageUrl, imageSize, imageType) VALUES (?, ?, ?, ?, ?)", ppid, bannerOffers[i].Image.Medium.ID, bannerOffers[i].Image.Medium.ImageURL, "large", "banner")
		if err != nil {
			return bannerOffers, err
		}
	}
	return bannerOffers, nil
}

func (r *repositoryBannerOfferCRUD) FindById(id int64) ([]models.Data, error) {
	var datas []models.Data
	var imageSize string
	var imageType string
	if id == 0 {
		rows, err := r.db.Query("SELECT * FROM BannerData")
		if err != nil {
			return datas, err
		}
		for rows.Next() {
			var data models.Data
			if err := rows.Scan(&data.ID, &data.Name, &data.LandingUrl, &data.OrderType, &data.MetaData.ProductCode, &data.MetaData.ProductName, &data.MetaData.ProfessionalSlug); err != nil {
				return datas, err
			}
			datas = append(datas, data)
		}
		defer rows.Close()
		
		for i, d := range datas {
			iid := d.ID
			row := r.db.QueryRow("SELECT * FROM ImageInfo WHERE pid = ? AND imageSize = ? AND imageType = ?", iid, "large", "banner")
			if err := row.Scan(&datas[i].Image.Large.ID, &datas[i].Image.Large.ImageURL, &iid, &imageSize, &imageType); err != nil {
				if err == sql.ErrNoRows {
					return datas, err
				}
				return datas, err
			}
			row = r.db.QueryRow("SELECT * FROM ImageInfo WHERE pid = ? AND imageSize = ? AND imageType = ?", iid, "medium", "banner")
			if err := row.Scan(&datas[i].Image.Medium.ID, &datas[i].Image.Medium.ImageURL, &iid, &imageSize, &imageType); err != nil {
				if err == sql.ErrNoRows {
					return datas, err
				}
				return datas, err
			}
		}
	} else {
		var data models.Data
		row := r.db.QueryRow("SELECT * FROM BannerData WHERE id = ?", id)
		if err := row.Scan(&data.ID, &data.Name, &data.LandingUrl, &data.OrderType, &data.MetaData.ProductCode, &data.MetaData.ProductName, &data.MetaData.ProfessionalSlug); err != nil {
			if err == sql.ErrNoRows {
				return datas, err
			}
			return datas, err
		}
		iid := data.ID
		row = r.db.QueryRow("SELECT * FROM ImageInfo WHERE pid = ? AND imageSize = ? AND imageType = ?", iid, "large", "banner")
		if err := row.Scan(&data.Image.Large.ID, &data.Image.Large.ImageURL, &iid, &imageSize, &imageType); err != nil {
			if err == sql.ErrNoRows {
				return datas, err
			}
			return datas, err
		}
		row = r.db.QueryRow("SELECT * FROM ImageInfo WHERE pid = ? AND imageSize = ? AND imageType = ?", iid, "medium", "banner")
		if err := row.Scan(&data.Image.Medium.ID, &data.Image.Medium.ImageURL, &iid, &imageSize, &imageType); err != nil {
			if err == sql.ErrNoRows {
				return datas, err
			}
			return datas, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}
