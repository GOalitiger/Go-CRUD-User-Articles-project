package controllers

import (
	"crudOperations/models"
	"crudOperations/shared"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// get all articles
func GetAllArticles(c echo.Context) error {
	a := []models.Article{}
	// a := []interface{}{}
	result := shared.Db.Find(&a)

	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, a)
}

// get article by id api
func GetArticleById(c echo.Context) error {
	a := models.Article{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	result := shared.Db.Find(&a, id)

	if result.Error != nil {
		return result.Error
	}

	if a.ArticleId == 0 {
		return c.String(http.StatusNotFound, "Article not found!")
	}

	return c.JSON(http.StatusOK, a)
}

// for creating a new article
func CreateArticle(c echo.Context) error {
	a := models.Article{}
	err := c.Bind(&a)
	if err != nil {
		return err
	}

	a.Created_DateTime = time.Now()
	result := shared.Db.Create(&a)

	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, a)
}

// update article
func UpdateArticle(c echo.Context) error {
	a := models.Article{}
	err := c.Bind(&a)
	if err != nil {
		return err
	}
	var dbArticle = models.Article{}

	result := shared.Db.Where("article_id = ?", a.ArticleId).Find(&dbArticle)

	if result.Error != nil {
		return result.Error
	}

	// article id is compared with the given id to update
	if dbArticle.ArticleId == a.ArticleId {
		result := shared.Db.Where("article_id = ?", a.ArticleId).Save(&a)
		if result.Error != nil {
			return result.Error
		}
	} else {
		return c.String(http.StatusNotFound, "Article id  not found!")
	}
	return c.JSON(http.StatusOK, a)
}

// delete article
func DeleteArticleById(c echo.Context) error {
	a_id := c.Param("id")

	result := shared.Db.Where("article_id = ?", a_id).Delete(&models.Article{})
	if result.Error != nil {
		return result.Error
	}

	return c.String(http.StatusOK, "Successfully Deleted")

}
