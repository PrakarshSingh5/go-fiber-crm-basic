package database

import (
	
		"github.com/jinzhu/gorm"
		"github.com/jinzhu/gorm/dialects/sqlite"
		"github.com/gofiber/fiber"
)

var (
	DBConn *gorm.DB
)
