package models

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type UserValidator struct {
	Validator *validator.Validate
}

func (uv *UserValidator) Validate(i interface{}) error {
	return uv.Validator.Struct(i)
}

type Users struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"  validate:"required,min=4"`
	Password   string    `json:"password,omitempty"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_on"`
	Articles   []Article `json:"articles" gorm:"foreignKey:UserId"`
}

type UserArticles struct {
	Users
	Article
}

// `gorm:"embedded"`
// ;references:id
type APIUsersArticles struct {
	APIUserWithoutArticles
	// Article `gorm:"foreignKey:UserId"`
	APIArticle
}

// gorm:"foreignKey:UserId"
type APIUser struct {
	Id         int          `json:"id"`
	Name       string       `json:"name"`
	Created_at time.Time    `json:"created_at"`
	Email      string       `json:"email"`
	Articles   []APIArticle `json:"articles" `
	// ArticlesJson interface{} `gorm:"column:articles_json" json:"articles_json"`
}

type APIUserForQuery struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Email      string    `json:"email"`
	// Articles   []APIArticle `json:"articles" `
	ArticlesJson *interface{} `gorm:"column:articles_json" json:"articles_json"`
}

type APIUserWithoutArticles struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Email      string    `json:"email"`
}

// mapping user api struct built on run time from embedded struct UserArticles
func (apiUser *APIUser) MapAPIUsers(userArticle *UserArticles) {
	apiUser.Id = userArticle.Id
	apiUser.Email = userArticle.Email
	apiUser.Name = userArticle.Name
	apiUser.Created_at = userArticle.Created_at
}
