package controllers

// import (
// 	"astro/database"
// 	"astro/models"
// 	"astro/repository"
// 	"astro/repository/crud"
// 	"astro/responses"
// 	"astro/utils/channels"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"

// 	"github.com/jinzhu/gorm"
// )

// // GetLicense is a func
// func GetLicense(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	licenses := []models.License{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		licenses, err = ProfileRepository.FindLicense(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, licenses)
// 	}(repo)

// }

// // PostLicense is a func
// func PostLicense(w http.ResponseWriter, r *http.Request) {

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	license := models.License{}
// 	err = json.Unmarshal(body, &license)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	license.UserID = pid
// 	license.FileURL = "upload/user" + strconv.Itoa(int(pid)) + "/license" + license.LicenseNumber

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveLicense(license)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, license)
// 	}(repo)
// }

// // PutLicense is a func
// func PutLicense(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	license := models.License{}
// 	err = json.Unmarshal(body, &license)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateLicense(int64(pid), int64(license.LicenseID), license)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, license)
// 	}(repo)
// }

// // DeleteLicense is a func
// func DeleteLicense(w http.ResponseWriter, r *http.Request) {

// 	urlParams := r.URL.Query()
// 	licenseID, err := strconv.ParseUint(urlParams.Get("license_id"), 10, 64)
// 	fmt.Println(licenseID)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteLicense(int64(licenseID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)

// }

// // GetSpecialization is a func
// func GetSpecialization(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	specializations := []models.Specialization{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		specializations, err = ProfileRepository.FindSpecialization(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, specializations)
// 	}(repo)
// }

// // PostSpecialization is a func
// func PostSpecialization(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	specialization := models.Specialization{}
// 	err = json.Unmarshal(body, &specialization)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	specialization.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveSpecialization(specialization)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, specialization)
// 	}(repo)
// }

// // PutSpecialization is a func
// func PutSpecialization(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	specialization := models.Specialization{}
// 	err = json.Unmarshal(body, &specialization)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateSpecialization(int64(pid), int64(specialization.SpecializationID), specialization)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, specialization)
// 	}(repo)
// }

// // DeleteSpecialization is a func
// func DeleteSpecialization(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	specializationID, err := strconv.ParseUint(urlParams.Get("specialization_id"), 10, 64)
// 	fmt.Println(specializationID)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteSpecialization(int64(specializationID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetCertification is a func
// func GetCertification(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	certifications := []models.Certification{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		certifications, err = ProfileRepository.FindCertification(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, certifications)
// 	}(repo)
// }

// // PostCertification is a func
// func PostCertification(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	certificate := models.Certification{}
// 	err = json.Unmarshal(body, &certificate)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	certificate.UserID = pid
// 	certificate.FileURL = "upload/user" + strconv.Itoa(int(pid)) + "/certification" + certificate.Certification

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveCertification(certificate)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, certificate)
// 	}(repo)
// }

// // PutCertification is a func
// func PutCertification(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	certificate := models.Certification{}
// 	err = json.Unmarshal(body, &certificate)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateCertification(int64(pid), int64(certificate.CertificateID), certificate)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, certificate)
// 	}(repo)
// }

// // DeleteCertification is a func
// func DeleteCertification(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	certificateID, err := strconv.ParseUint(urlParams.Get("certificate_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteCertification(int64(certificateID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetReference is a func
// func GetReference(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	references := []models.Reference{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		references, err = ProfileRepository.FindReference(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, references)
// 	}(repo)
// }

// // PostReference is a func
// func PostReference(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	reference := models.Reference{}
// 	err = json.Unmarshal(body, &reference)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	reference.UserID = pid
// 	reference.LORURL = "upload/user" + strconv.Itoa(int(pid)) + "/reference" + reference.Name

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveReference(reference)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, reference)
// 	}(repo)
// }

// // PutReference is a func
// func PutReference(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	reference := models.Reference{}
// 	err = json.Unmarshal(body, &reference)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateReference(int64(pid), int64(reference.ReferenceID), reference)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, reference)
// 	}(repo)
// }

// // DeleteReference is a func
// func DeleteReference(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	referenceID, err := strconv.ParseUint(urlParams.Get("reference_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteReference(int64(referenceID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetWorkExperience is a func
// func GetWorkExperience(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	workExperiences := []models.WorkExperience{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		workExperiences, err = ProfileRepository.FindWorkExperience(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		for j, i := range workExperiences {
// 			wid := i.WorkExperienceID
// 			fmt.Println("wid  = ", wid)
// 			repo := crud.NewRepositoryProfileCRUD(db)
// 			var positions []models.Position
// 			func(ProfileRepository repository.ProfileRepository) {
// 				positions, _ = ProfileRepository.FindPosition(uint64(pid), int64(wid))
// 			}(repo)
// 			workExperiences[j].Positions = positions
// 			positions = nil
// 		}
// 		responses.JSON(w, http.StatusOK, workExperiences)
// 	}(repo)
// }

// // PostWorkExperience is a func
// func PostWorkExperience(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	workExperience := models.WorkExperience{}
// 	err = json.Unmarshal(body, &workExperience)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	workExperience.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveWorkExperience(workExperience)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, workExperience)
// 	}(repo)
// }

// // PutWorkExperience is a func
// func PutWorkExperience(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	workExperience := models.WorkExperience{}
// 	err = json.Unmarshal(body, &workExperience)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateWorkExperience(int64(pid), int64(workExperience.WorkExperienceID), workExperience)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, workExperience)
// 	}(repo)
// }

// // DeleteWorkExperience is a func
// func DeleteWorkExperience(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	workExperienceID, err := strconv.ParseUint(urlParams.Get("work_experience_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		var rs *gorm.DB
// 		done := make(chan bool)
// 		go func(ch chan<- bool) {
// 			rs = db.Debug().Model(models.Position{}).Where("user_id = ? and work_experience_id = ?", pid, workExperienceID).Take(&models.Position{}).Delete(&models.Position{})
// 			ch <- true
// 		}(done)
// 		if channels.OK(done) {
// 			if rs.Error != nil {
// 				if gorm.IsRecordNotFoundError(rs.Error) {
// 					fmt.Println(errors.New("position not found"))
// 				}
// 			}
// 		}
// 		err = ProfileRepository.DeleteWorkExperience(int64(workExperienceID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // PostPosition is a func
// func PostPosition(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	position := models.Position{}
// 	err = json.Unmarshal(body, &position)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	position.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SavePosition(position)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, position)
// 	}(repo)
// }

// // PutPosition is a func
// func PutPosition(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	position := models.Position{}
// 	err = json.Unmarshal(body, &position)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdatePosition(int64(pid), int64(position.PositionID), int64(position.WorkExperienceID), position)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, position)
// 	}(repo)
// }

// // DeletePosition is a func
// func DeletePosition(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	positionID, err := strconv.ParseUint(urlParams.Get("position_id"), 10, 64)
// 	workExperienceID, err := strconv.ParseUint(urlParams.Get("work_experience_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeletePosition(int64(positionID), int64(pid), int64(workExperienceID))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetEducation is a func
// func GetEducation(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	educations := []models.Education{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		educations, err = ProfileRepository.FindEducation(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, educations)
// 	}(repo)
// }

// // PostEducation is a func
// func PostEducation(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	education := models.Education{}
// 	err = json.Unmarshal(body, &education)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	education.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveEducation(education)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, education)
// 	}(repo)
// }

// // PutEducation is a func
// func PutEducation(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	education := models.Education{}
// 	err = json.Unmarshal(body, &education)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateEducation(int64(pid), int64(education.EducationID), education)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, education)
// 	}(repo)
// }

// // DeleteEducation is a func
// func DeleteEducation(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	educationID, err := strconv.ParseUint(urlParams.Get("education_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteEducation(int64(educationID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetLanguage is a func
// func GetLanguage(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	languages := []models.Language{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		languages, err = ProfileRepository.FindLanguage(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, languages)
// 	}(repo)
// }

// // PostLanguage is a func
// func PostLanguage(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	language := models.Language{}
// 	err = json.Unmarshal(body, &language)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	language.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveLanguage(language)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, language)
// 	}(repo)
// }

// // PutLanguage is a func
// func PutLanguage(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	language := models.Language{}
// 	err = json.Unmarshal(body, &language)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateLanguage(int64(pid), int64(language.LanguageID), language)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, language)
// 	}(repo)
// }

// // DeleteLanguage is a func
// func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	languageID, err := strconv.ParseUint(urlParams.Get("language_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteLanguage(int64(languageID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetAward is a func
// func GetAward(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	awards := []models.Award{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		awards, err = ProfileRepository.FindAward(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, awards)
// 	}(repo)
// }

// // PostAward is a func
// func PostAward(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	award := models.Award{}
// 	err = json.Unmarshal(body, &award)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	award.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveAward(award)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, award)
// 	}(repo)
// }

// // PutAward is a func
// func PutAward(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	award := models.Award{}
// 	err = json.Unmarshal(body, &award)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateAward(int64(pid), int64(award.AwardID), award)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, award)
// 	}(repo)
// }

// // DeleteAward is a func
// func DeleteAward(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	awardID, err := strconv.ParseUint(urlParams.Get("award_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteAward(int64(awardID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetVolunteering is a func
// func GetVolunteering(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	volunteerings := []models.Volunteering{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		volunteerings, err = ProfileRepository.FindVolunteering(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, volunteerings)
// 	}(repo)
// }

// // PostVolunteering is a func
// func PostVolunteering(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	volunteering := models.Volunteering{}
// 	err = json.Unmarshal(body, &volunteering)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	volunteering.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveVolunteering(volunteering)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, volunteering)
// 	}(repo)
// }

// // PutVolunteering is a func
// func PutVolunteering(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	volunteering := models.Volunteering{}
// 	err = json.Unmarshal(body, &volunteering)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateVolunteering(int64(pid), int64(volunteering.VolunteeringID), volunteering)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, volunteering)
// 	}(repo)
// }

// // DeleteVolunteering is a func
// func DeleteVolunteering(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	volunteeringID, err := strconv.ParseUint(urlParams.Get("volunteering_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteVolunteering(int64(volunteeringID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }

// // GetProject is a func
// func GetProject(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	projects := []models.Project{}

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		projects, err = ProfileRepository.FindProject(uint64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, projects)
// 	}(repo)
// }

// // PostProject is a func
// func PostProject(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	project := models.Project{}
// 	err = json.Unmarshal(body, &project)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	project.UserID = pid

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.SaveProject(project)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, project)
// 	}(repo)
// }

// // PutProject is a func
// func PutProject(w http.ResponseWriter, r *http.Request) {
// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	project := models.Project{}
// 	err = json.Unmarshal(body, &project)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.UpdateProject(int64(pid), int64(project.ProjectID), project)
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, project)
// 	}(repo)
// }

// // DeleteProject is a func
// func DeleteProject(w http.ResponseWriter, r *http.Request) {
// 	urlParams := r.URL.Query()
// 	projectID, err := strconv.ParseUint(urlParams.Get("project_id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	pid, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repo := crud.NewRepositoryProfileCRUD(db)

// 	func(ProfileRepository repository.ProfileRepository) {
// 		err = ProfileRepository.DeleteProject(int64(projectID), int64(pid))
// 		if err != nil {
// 			responses.ERROR(w, http.StatusInternalServerError, err)
// 			return
// 		}
// 		responses.JSON(w, http.StatusOK, "deleted successfully")
// 	}(repo)
// }
