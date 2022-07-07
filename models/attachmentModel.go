package models

type Attachment struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Extend string
}
