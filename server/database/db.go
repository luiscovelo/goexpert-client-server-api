package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type connDB struct {
	DB *gorm.DB
}

func (c *connDB) Close() {
	if c != nil {
		c.Close()
	}
}

func New() (*connDB, error) {
	db, err := gorm.Open(sqlite.Open("database/quotation.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &connDB{DB: db}, nil
}

func (c *connDB) AutoMigrate(objs ...interface{}) error {
	return c.DB.AutoMigrate(objs...)
}
