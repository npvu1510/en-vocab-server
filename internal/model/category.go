package model

import (
	"fmt"

	baseRepo "github.com/saoladigital/sld-go-common/repository"
)

type Category struct {
	Id   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`

	Dictionaries []Dictionary `gorm:"many2many:dictionaries_categories"`

	baseRepo.BaseEntity
}

// Constructors
func NewCategory(id int, name string) *Category {
	return &Category{Id: id, Name: name}
}

// Methods
func (c *Category) DisplayCategory() {
	fmt.Printf("Category Id: %d\nCategory Name: %s\n", c.Id, c.Name)
}

func (c Category) TableName() string {
	return "categories"
}
