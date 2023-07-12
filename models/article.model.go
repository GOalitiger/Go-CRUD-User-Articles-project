package models

import "time"

//remember by default the gorm converts the name into snake case
//ArticleId = article_id
type Article struct {
	// ArticleInputId int       `json:"article_Id,omitempty" gorm:"column:article_id"`
	ArticleId        int       `json:"article_id,omitempty" gorm:"column:article_id"`
	UserId           int       `json:"user_id" `
	ArticleName      string    `json:"article_name"`
	Content          string    `json:"content"`
	IsPublic         bool      `json:"is_public"`
	Created_DateTime time.Time `json:"created_at" gorm:"column:created_at"`
}

type APIArticle struct {
	ArticleId        int       `json:"article_id"`
	ArticleName      string    `json:"article_name"`
	Content          string    `json:"content"`
	IsPublic         bool      `json:"is_public"`
	Created_DateTime time.Time `json:"created_date_time" gorm:"column:created_at"`
}

// mapping article api struct built on run time from embedded struct UserArticles
func (apiArticle *APIArticle) MapAPIArticles(userArticle *UserArticles) {
	apiArticle.ArticleId = userArticle.ArticleId
	apiArticle.ArticleName = userArticle.ArticleName
	apiArticle.Content = userArticle.Content
	apiArticle.IsPublic = userArticle.IsPublic
	apiArticle.Created_DateTime = userArticle.Created_DateTime
}
