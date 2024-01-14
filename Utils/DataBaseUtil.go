package Utils

type Job struct {
	Id          uint   `gorm:"primaryKey"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	PostedBy    uint   `json:"PostedBy"`
	Status      string `json:"Status"`
}

type User struct {
	UId           uint   `gorm:"primaryKey"`
	Password      string `json:"Password"`
	Email         string `json:"Email"`
	ResumeAddress string `json:"ResumeAddress"`
	JobId         uint   `json:"JobId"`
}

type HRUser struct {
	HId      uint   `gorm:"primaryKey"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
