package controllers

import (
	"crudOperations/models"
	"crudOperations/shared"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

// omitempty tag is added here it will omit then field it from json response if empty
type LoginUser struct {
	EmailId  string `json:"email"`
	Password string `json:"password,omitempty"`
}

// login with generation of jwt and returns token in header
func Login(c echo.Context) error {
	userCred := LoginUser{}
	if err := c.Bind(&userCred); err != nil {
		c.Logger().Error(err)
		return err
	}
	// for validation of email and password if they exist in DB

	ok, err := validateEmailPassword(&userCred)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	if ok {
		jwtToken, err := GenerateJWT(&userCred)
		if err != nil {
			fmt.Println(err)
			return err
		}
		c.Response().Header().Set("x-auth-token", "Bearer "+jwtToken)
		return c.JSON(http.StatusOK, LoginUser{EmailId: userCred.EmailId})
	} else {
		return c.String(http.StatusNotFound, "Email id or password not valid")
	}

}

//Generates the JWT token
/*
	-first we assign it header to define the algo to generate jwt
	-add claims to the token includes expiration time and etc
	-then we add the secret signature , we need to get from env variable
*/

func GenerateJWT(user *LoginUser) (string, error) {
	cnfig := models.ConfigDatabase{}
	if err := cleanenv.ReadEnv(&cnfig); err != nil {
		shared.EchoNew.Logger.Fatal(err)
		return "", err
	}
	//generating a token with siginging algo HMAC
	token := jwt.New(jwt.SigningMethodHS256)

	//generation of claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["EmailId"] = (*user).EmailId

	tokenStr, err := token.SignedString([]byte(cnfig.JWTSecretSigningKey))
	if err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}

}

// validating if the emailID and password given is present in DB or not
func validateEmailPassword(user *LoginUser) (bool, error) {
	dbUser := models.Users{}

	result := shared.Db.Where("email = ? AND password = ?", user.EmailId, user.Password).Find(&dbUser)
	if result.Error != nil {
		return false, result.Error
	}

	if (dbUser.Email == user.EmailId) && (dbUser.Password == user.Password) {
		return true, nil
	} else {
		return false, nil
	}

}
