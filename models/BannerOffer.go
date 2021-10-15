package models

type Banner struct {
	HttpStatus     string `json:"httpStatus"`
	HttpStatusCode int    `json:"httpStatusCode"`
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	ApiName        string `json:"apiName"`
	Datas          []Data `json:"data"`
}

type Data struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	LandingUrl string   `json:"landingUrl"`
	OrderType  string   `json:"orderType"`
	MetaData   MetaData `json:"metaData"`
	Image      Image    `json:"images"`
}

type MetaData struct {
	ProductCode      string `json:"productCode"`
	ProductName      string `json:"productName"`
	ProfessionalSlug string `json:"professionalSlug"`
}
