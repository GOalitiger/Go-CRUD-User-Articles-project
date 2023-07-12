package main

import (
	"crudOperations/config"
	"crudOperations/shared"
	"fmt"
	"os"

	"crudOperations/routes"
	// "crudOperations/shared"
	"gopkg.in/go-playground/validator.v9"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		shared.EchoNew.Logger.Fatal(err)
	}
	if err := cleanenv.ReadEnv(&shared.CnfigDB); err != nil {
		shared.EchoNew.Logger.Fatal(err)
	}
}

func main() {
	// e := echo.New()
	// v := validator.New()
	// shared.EchoNew = e
	// shared.ValidatorNew = v

	fmt.Println("Port = ", os.Getenv("PORT"))
	fmt.Println("config Port = ", shared.CnfigDB)

	shared.EchoNew = echo.New()
	shared.ValidatorNew = validator.New()

	_, err := config.ConnectDB()

	if err != nil {
		fmt.Println("No connection with DB established!")
		panic(err)
	}

	// shared.Db = Db
	// routes.SettingRoutes(e)
	routes.SettingRoutes(shared.EchoNew)
	defer shared.SqlDB.Close()

	// e.Logger.Print("Server is starting on 8080")
	// e.Logger.Fatal(e.Start(":8080"))

	shared.EchoNew.Logger.Print("Server is starting on 8080")
	shared.EchoNew.Logger.Fatal(shared.EchoNew.Start(":8080"))

}
