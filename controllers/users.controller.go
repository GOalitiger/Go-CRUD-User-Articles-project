package controllers

import (
	"crudOperations/models"
	"crudOperations/shared"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// For getting all the users
func GetAllUsers(c echo.Context) error {
	users := []models.Users{}
	result := shared.Db.Find(&users)
	if result.Error != nil {
		// c.Logger().Print(err)
		return result.Error
	}

	// fmt.Println(users)
	if len(users) != 0 {
		return c.JSON(http.StatusOK, users)
	} else {
		return c.String(http.StatusNotFound, "No users found")
	}
}

// create user
func CreateUser(c echo.Context) error {
	user := models.Users{}

	// shared.EchoNew.Validator = &(models.UserValidator{Validator: shared.ValidatorNew})

	err := c.Bind(&user)
	if err != nil {
		return err
	}

	// if err = c.Validate(user); err != nil {
	// 	return err
	// }

	err1 := shared.ValidatorNew.Struct(user)
	if err != nil {
		return err1
	}

	user.Created_at = time.Now()
	result := shared.Db.Omit("Id").Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusOK, result.Error)
	}
	return c.JSON(http.StatusOK, user)
}

// Getting user by passing ID in params
func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	user := models.Users{}
	result := shared.Db.Find(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusOK, result.Error)
	}
	return c.JSON(http.StatusOK, user)
}

// Updating user
func UpdateUser(c echo.Context) error {
	currUser := models.Users{}

	if err := c.Bind(&currUser); err != nil {
		return err
	}
	if err := shared.ValidatorNew.Struct(currUser); err != nil {
		return err
	}

	users := []models.Users{}
	result := shared.Db.Find(&users)
	if result.Error != nil {
		// c.Logger().Print(err)
		return result.Error
	}

	idFound := -1

	if len(users) > 0 {
		for _, u := range users {
			if currUser.Id == u.Id {
				idFound = u.Id
				break
			}
		}

		if idFound != -1 {
			// shared.Db.Save(&currUser)
			res := shared.Db.Model(&currUser).Select("*").Omit("Id").Updates(currUser)
			if res.Error != nil {
				// c.Logger().Print(err)
				return res.Error
			}
			return c.JSON(http.StatusOK, currUser)
		} else {
			//will return that such id doesnot exist
			return c.String(http.StatusNotFound, fmt.Sprintf("No id =%v exists", currUser.Id))
		}

	} else {
		return c.String(http.StatusNotFound, "users List empty")
	}

}

// deleting user
func DeleteUserById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	var user models.Users
	// result := shared.Db.Delete(&models.Users{}, id)

	if result := shared.Db.Find(&user, id); result.Error != nil {
		return result.Error
	} else {
		fmt.Println(result.RowsAffected)
	}

	if user.Id != 0 {
		result := shared.Db.Delete(&user)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(http.StatusOK, user)
	} else {
		return c.String(http.StatusNotFound, fmt.Sprintf("No user found with id = %v", id))
	}

}

// Using Preloading
func GetUsersWithThereArticles(c echo.Context) error {
	usersWithArticles := []models.Users{}
	result := shared.Db.Preload("Articles").Find(&usersWithArticles)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, usersWithArticles)
}

// Getting users with articles using logic after getting values from join
func GetUsersWithThereArticlesUsingJOINWithLogic(c echo.Context) error {

	var userJoinArtices []models.UserArticles
	result := shared.Db.Model(models.Users{}).Select("*").
		Joins("JOIN articles ON users.id = articles.user_Id").Find(&userJoinArtices)

	if result.Error != nil {
		return result.Error
	}

	// after getting join mapping the user and article api
	if len(userJoinArtices) > 0 {
		usersWithArticles := []models.APIUser{}
		usersWithArticlesMap := map[int]*models.APIUser{}
		for _, v := range userJoinArtices {
			var ViewUser models.APIUser
			var ViewArticle models.APIArticle
			ViewUser.MapAPIUsers(&v)
			ViewArticle.MapAPIArticles(&v)

			// using map to push them in and then get them out
			apiUserTempCheck, ok := usersWithArticlesMap[ViewUser.Id]
			if ok {
				apiUserTempCheck.Articles = append(apiUserTempCheck.Articles, ViewArticle)
			} else {
				ViewUser.Articles = append(ViewUser.Articles, ViewArticle)
				usersWithArticlesMap[ViewUser.Id] = &ViewUser
			}
		}

		// traversing the map to get the userAPI from it
		for _, v := range usersWithArticlesMap {
			usersWithArticles = append(usersWithArticles, *v)
		}

		return c.JSON(http.StatusOK, usersWithArticles)
	}

	return c.String(http.StatusNotFound, "No Data found ")
	//

}

func GetUsersWithThereArticlesUsingJOIN(c echo.Context) error {
	usersWithArticles := []models.APIUserForQuery{}

	result := shared.Db.Model(models.Users{}).Select(`users.id, users.name,users.created_at,users.email,
		 JSON_ARRAYAGG(JSON_OBJECT('article_id',articles.article_id,
		 'article_name',articles.article_name,
		 'content',articles.content,
		 'is_public', articles.is_public,
		 'created_date_time', articles.created_at)) as articles_json`).
		Joins("JOIN articles ON users.id = articles.user_Id").Group("users.id").Find(&usersWithArticles)
	// Find(&usersWithArticles)

	if result.Error != nil {
		return result.Error
	}

	//
	fmt.Println(usersWithArticles)
	return c.JSON(http.StatusOK, usersWithArticles)
}

// We get these by embedding the APIUser and APIArticle
func GettingUserArticlesTogetherFromJoin(c echo.Context) error {
	var userJoinArtices []models.APIUsersArticles
	result := shared.Db.Model(models.Users{}).Select("*").
		Joins("JOIN articles ON users.id = articles.user_Id").Find(&userJoinArtices)

	if result.Error != nil {
		return result.Error
	}
	type response struct {
		Status   int64       `json:"status"`
		Response interface{} `json:"response"`
		Length   int         `json:"length"`
	}

	var res response

	res.Response = userJoinArtices
	res.Length = len(userJoinArtices)
	res.Status = http.StatusOK

	return c.JSON(http.StatusOK, res)
	// return c.JSON(http.StatusOK, userJoinArtices)
}
