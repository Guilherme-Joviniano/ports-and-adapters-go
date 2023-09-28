package orm

import (
	"gorm.io/gorm"
)

type ProductOrmAdapter struct {
	db gorm.DB
}

// TODO: Finish this adapter
// func NewProductOrmAdapter() *ProductOrmAdapter {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	db.AutoMigrate(&application.Product{})

// 	return &ProductOrmAdapter{
// 		db: *db,
// 	}
// }
