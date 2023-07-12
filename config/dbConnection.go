package config

import (
	"crudOperations/shared"
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var cnfigDB models.ConfigDatabase

const (
	username = "root"
	password = "admin"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "test"
)

// this function runs before any other function in a particular package
func init() {
	// if err := godotenv.Load(".env"); err != nil {
	// 	shared.EchoNew.Logger.Fatal(err)
	// }
	// if err := cleanenv.ReadEnv(&cnfigDB); err != nil {
	// 	shared.EchoNew.Logger.Fatal(err)
	// }
	// fmt.Println(cnfigDB)

	// fmt.Println("Port = ", os.Getenv("PORT"))

}

func ConnectDB() (*gorm.DB, error) {
	fmt.Println(shared.CnfigDB)

	fmt.Println("Port = ", os.Getenv("PORT"))
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// from constants
	// dataConncectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	// 	username, password, host, port, dbname)

	// from env file
	// dataConncectionString := "root:admin@tcp(mysqldb:3304)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dataConncectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		shared.CnfigDB.User, shared.CnfigDB.Password, shared.CnfigDB.Host, shared.CnfigDB.Port, shared.CnfigDB.DbName)

	sqlDB, err1 := sql.Open("mysql", dataConncectionString)
	if err1 != nil {
		return nil, err1
	}

	// logger config line added to log the query generated

	gormDB, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err2 != nil {
		return nil, err2
	}

	shared.Db = gormDB
	shared.SqlDB = sqlDB

	// defer sqlDB.Close()
	return gormDB, nil

}
