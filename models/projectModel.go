package models

type Project struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"project_name"`
	Summary string `json:"project_summary"`
}
