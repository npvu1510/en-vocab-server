package model

import "gorm.io/gorm"

type Dictionary struct {
	Id         uint   `gorm:"column:id;primaryKey;<-:create" json:"id"`
	Definition string `gorm:"column:definition;size:50;not null;index" json:"definition"`
	Image      string `gorm:"column:image;size:255" json:"image"`
	AudioGB    string `gorm:"column:audio_gb;size:255" json:"audio_gb"`
	AudioUS    string `gorm:"column:audio_us;size:255" json:"audio_us"`
	Source     string `gorm:"size:255;not null;<-:create" json:"source"`

	Categories []Category `gorm:"many2many:dictionaries_categories"`
	gorm.Model
}

// CONSTRUCTOR
func NewDictionary(definition string, categories []Category, imgSrc string, source string) *Dictionary {
	return &Dictionary{Definition: definition, Categories: categories, Image: imgSrc, Source: source}
}
