package crud

import (
	"astro/models"
	"database/sql"
	"fmt"
)

type repositoryAstroCRUD struct {
	db *sql.DB
}

//NewRepositoryBasicInfoCRUD is
func NewrepositoryAstroCRUD(db *sql.DB) *repositoryAstroCRUD {
	return &repositoryAstroCRUD{db}
}

func (r *repositoryAstroCRUD) Save(astro models.Astro) (int, error) {
	astroQuery, err := r.db.Exec("INSERT INTO Astro (urlSlug, namePrefix, firstName, lastName, aboutMe, profliePicUrl, experience, rating, minimumCallDuration, minimumCallDurationCharges, additionalPerMinuteCharges, isAvailable, isOnCall) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", astro.UrlSlug, astro.NamePrefix, astro.FirstName, astro.LastName, astro.AboutMe, astro.ProfliePicUrl, astro.Experience, astro.Rating, astro.MinimumCallDuration, astro.MinimumCallDurationCharges, astro.AdditionalPerMinuteCharges, astro.IsAvailable, astro.IsOnCall)
	if err != nil {
		return -1, err
	}
	id, err := astroQuery.LastInsertId()
	if err != nil {
		return -1, err
	}
	ImageInfoQuery, err := r.db.Exec("INSERT INTO ImageInfo (pid, id, imageUrl, imageSize, imageType) VALUES (?, ?, ?, ?, ?)", id, astro.Images.Large.ID, astro.Images.Large.ImageURL, "medium", "astro")
	if err != nil {
		return -1, err
	}
	_, err = ImageInfoQuery.LastInsertId()
	if err != nil {
		return -1, err
	}
	ImageInfoQuery, err = r.db.Exec("INSERT INTO ImageInfo (pid, id, imageUrl, imageSize, imageType) VALUES (?, ?, ?, ?, ?)", id, astro.Images.Medium.ID, astro.Images.Medium.ImageURL, "large", "astro")
	if err != nil {
		return -1, err
	}
	_, err = ImageInfoQuery.LastInsertId()
	if err != nil {
		return -1, err
	}

	for _, element := range astro.Availability.Days {
		daysQuery, err := r.db.Exec("INSERT INTO Days (id, day) VALUES (?, ?)", id, element)
		if err != nil {
			return -1, err
		}
		_, err = daysQuery.LastInsertId()
		if err != nil {
			return -1, err
		}
	}

	slotInfoQuery, err := r.db.Exec("INSERT INTO SlotInfo (id, toFormat, fromFormat, fromTime, toTime) VALUES (?, ?, ?, ?, ?)", id, astro.Availability.Slot.ToFormat, astro.Availability.Slot.Fromformat, astro.Availability.Slot.From, astro.Availability.Slot.To)
	if err != nil {
		return -1, err
	}
	_, err = slotInfoQuery.LastInsertId()
	if err != nil {
		return -1, err
	}

	for _, element := range astro.Skills {
		skillQuery, err := r.db.Exec("INSERT INTO Skills (id, skill) VALUES (?, ?)", id, element)
		if err != nil {
			return -1, err
		}
		_, err = skillQuery.LastInsertId()
		if err != nil {
			return -1, err
		}
	}
	for _, element := range astro.Languages {
		languageQuery, err := r.db.Exec("INSERT INTO Language (id, language) VALUES (?, ?)", id, element)
		if err != nil {
			return -1, err
		}
		_, err = languageQuery.LastInsertId()
		if err != nil {
			return -1, err
		}
	}

	return int(id), nil
}

func (r *repositoryAstroCRUD) FindById(id uint64) (models.Astro, error) {
	var alb models.Astro

	//Quering form Astro Table
	astro := r.db.QueryRow("SELECT * FROM Astro WHERE id = ?", id)
	if err := astro.Scan(&alb.ID, &alb.UrlSlug, &alb.NamePrefix, &alb.FirstName, &alb.LastName, &alb.AboutMe, &alb.ProfliePicUrl, &alb.Experience, &alb.Rating, &alb.MinimumCallDuration, &alb.MinimumCallDurationCharges, &alb.AdditionalPerMinuteCharges, &alb.IsAvailable, &alb.IsOnCall); err != nil {
		if err == sql.ErrNoRows {
			return alb, err
		}
		return alb, err
	}

	days, err := r.db.Query("SELECT * FROM Days WHERE id = ?", id)
	if err != nil {
		return alb, err
	}

	for days.Next() {
		var day string
		if err := days.Scan(&id, &day); err != nil {
			return alb, err
		}
		alb.Availability.Days = append(alb.Availability.Days, day)
	}
	defer days.Close()

	slotinfo := r.db.QueryRow("SELECT * FROM SlotInfo WHERE id = ?", id)
	if err := slotinfo.Scan(&id, &alb.Availability.Slot.ToFormat, &alb.Availability.Slot.Fromformat, &alb.Availability.Slot.From, &alb.Availability.Slot.To); err != nil {
		if err == sql.ErrNoRows {
			return alb, err
		}
		return alb, err
	}
	skills, err := r.db.Query("SELECT * FROM Skills WHERE id = ?", id)
	if err != nil {
		return alb, err
	}

	for skills.Next() {
		var skill string
		if err := skills.Scan(&id, &skill); err != nil {
			return alb, err
		}
		alb.Skills = append(alb.Skills, skill)
	}
	defer skills.Close()
	languages, err := r.db.Query("SELECT * FROM Language WHERE id = ?", id)
	if err != nil {
		return alb, err
	}

	for languages.Next() {
		var language string
		if err := languages.Scan(&id, &language); err != nil {
			return alb, err
		}
		alb.Languages = append(alb.Languages, language)
	}
	defer languages.Close()
	var imageSize string
	var imageType string
	mediumimageinfo := r.db.QueryRow("SELECT * FROM ImageInfo WHERE pid = ? AND imageSize = ? AND imageType = ?", id, "medium", "astro")
	if err := mediumimageinfo.Scan(&alb.Images.Medium.ID, &alb.Images.Medium.ImageURL, &id, &imageSize, &imageType); err != nil {
		if err == sql.ErrNoRows {
			return alb, err
		}
		return alb, err
	}
	largeimageinfo := r.db.QueryRow("SELECT * FROM ImageInfo WHERE pid = ? AND imageSize = ? AND imageType = ?", id, "large", "astro")
	if err := largeimageinfo.Scan(&alb.Images.Large.ID, &alb.Images.Large.ImageURL, &id, &imageSize, &imageType); err != nil {
		if err == sql.ErrNoRows {
			return alb, err
		}
		fmt.Println(err)
		return alb, err
	}
	return alb, nil
}

