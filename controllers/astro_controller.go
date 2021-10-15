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

//CreateAstro is a func
func CreateAstro(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	astro := models.Astro{}
	
	err = json.Unmarshal(body, &astro)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var id int
	repo := crud.NewrepositoryAstroCRUD(db)
	func(astroRepository repository.AstroRepository) {
				id, err = astroRepository.Save(astro)
				if err != nil {
					responses.ERROR(w, http.StatusUnprocessableEntity, err)
					return
				}
				w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	}(repo)
	astro.ID = uint32(id)
	
	responses.JSON(w, http.StatusCreated, astro)
}
func GetAstro(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	astro := models.Astro{}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := crud.NewrepositoryAstroCRUD(db)
	func(astroRepository repository.AstroRepository) {
				astro, err = astroRepository.FindById(uint64(id))
				if err != nil {
					responses.ERROR(w, http.StatusUnprocessableEntity, err)
					return
				}
				w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	}(repo)
	responses.JSON(w, http.StatusOK, astro)
 }
func UpdateAstro(w http.ResponseWriter, r *http.Request) {

}

func DeleteAstro(w http.ResponseWriter, r *http.Request) {

}