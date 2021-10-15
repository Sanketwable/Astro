package models

//Profession is a struct
type Astro struct {
	ID                         uint32           `json:"id"`
	UrlSlug                    string           `json:"urlSlug"`
	NamePrefix                 string           `json:"namePrefix"`
	FirstName                  string           `json:"firstName"`
	LastName                   string           `json:"lastName"`
	AboutMe                    string           `json:"aboutMe"`
	ProfliePicUrl              string           `json:"profliePicUrl"`
	Experience                 float32          `json:"experience"`
	Rating                     uint32           `json:"rating"`
	MinimumCallDuration        float32          `json:"minimumCallDuration"`
	MinimumCallDurationCharges float32          `json:"minimumCallDurationCharges"`
	AdditionalPerMinuteCharges float32          `json:"additionalPerMinuteCharges"`
	IsAvailable                bool             `json:"isAvailable"`
	IsOnCall                   bool             `json:"isOnCall"`
	Languages                  []string         `json:"languages"`
	Skills                     []string         `json:"skills"`
	Images                     Image            `json:"images"`
	Availability               AvailabilityInfo `json:"availability"`
}

// Image is a struct
type Image struct {
	Large  ImageInfo `json:"large"`
	Medium ImageInfo `json:"medium"`
}

// ImageInfo is a Struct
type ImageInfo struct {
	ImageURL string `json:"imageUrl"`
	ID       uint32 `json:"id"`
}

// AvailabilityInfo is a Struct
type AvailabilityInfo struct {
	Days []string `json:"days"`
	Slot SlotInfo `json:"slot"`
}

// SlotInfo is a struct
type SlotInfo struct {
	ToFormat   string `json:"toFormat"`
	Fromformat string `json:"fromFormat"`
	From       string `json:"from"`
	To         string `json:"to"`
}
