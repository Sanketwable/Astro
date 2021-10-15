package controllers

import (
	"astro/database"
	"astro/models"
	"astro/repository"
	"astro/repository/crud"
	"astro/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateBannerOffer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	banner := models.Banner{}
	json.Unmarshal(body, &banner)
	Data := banner.Datas
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := crud.NewrepositoryBannerOfferCRUD(db)
	func(bannerOfferRepository repository.BannerOfferRepository) {
				Data, err = bannerOfferRepository.Save(Data)
				if err != nil {
					responses.ERROR(w, http.StatusUnprocessableEntity, err)
					return
				}
				w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	}(repo)
	banner.Datas = Data;
	banner.ApiName = "Fetch Banner"
	banner.HttpStatusCode = 200
	banner.HttpStatus = "OK"
	banner.Message = "Banners fetched successfully"
	responses.JSON(w, http.StatusCreated, banner)
}

func GetBannerOffer(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(pid)
	banner := models.Banner{}
	Data := []models.Data{}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := crud.NewrepositoryBannerOfferCRUD(db)
	func(bannerOfferRepository repository.BannerOfferRepository) {
		Data, err = bannerOfferRepository.FindById(int64(id))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	}(repo)
	
	banner.Datas = Data;
	banner.ApiName = "Fetch Banner"
	banner.HttpStatusCode = 200
	banner.HttpStatus = "OK"
	banner.Message = "Banners fetched successfully"
	responses.JSON(w, http.StatusCreated, banner)

}