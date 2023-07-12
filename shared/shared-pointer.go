package shared

import (
	"crudOperations/models"
	"database/sql"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

var Db *gorm.DB
var ValidatorNew *validator.Validate
var SqlDB *sql.DB
var EchoNew *echo.Echo
var CnfigDB models.ConfigDatabase

// func (db *DB) Debug() (tx *DB) {
// 	return db.Session(&Session{
// 	  Logger:db.Logger.LogMode(logger.Info),
// 	})
//   }
// var e

// func Start() {

// }
