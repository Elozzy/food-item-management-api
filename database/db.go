package database

import (

	"gorm.io/gorm"
	_"github.com/mattn/go-sqlite3"

)

var (
	DBConn *gorm.DB
)