package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/npvu1510/en-vocab-server/internal/model"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "Create dictionaries table",
			Migrate: func(db *gorm.DB) error {
				return db.Migrator().CreateTable(&model.Dictionary{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(&model.Dictionary{})
			},
		},
		{
			ID: "Create categories table",
			Migrate: func(db *gorm.DB) error {
				return db.Migrator().CreateTable(&model.Category{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(&model.Category{})
			},
		},
		{
			ID: "Add gorm model to categies model",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&model.Category{})
			},
		},
		{
			ID: "Upgrade dictionary and category models (field-level perrmissions, size, index, ...)",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&model.Dictionary{}, &model.Category{})
			},
		},
	})
}
