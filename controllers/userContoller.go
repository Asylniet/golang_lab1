package controllers

import (
	"github.com/gin-gonic/gin"
	"lab1/models"
	"lab1/services"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	users, result := services.GetUsers()
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, result := services.GetUserByName(name)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func AddUser(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	result := services.AddUser(name, password)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't add user, please try again",
		})
		return
	}

	c.Status(http.StatusOK)
}

func LoginUser(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	user, result := services.LoginUser(name, password)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
		return
	}

	if user == (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")
	result := services.RemoveUser(id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.Status(http.StatusOK)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	password := c.Query("password")
	result := services.UpdateUser(id, name, password)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}
	c.Status(http.StatusOK)
}

func RateItem(c *gin.Context) {
	id := c.Param("id")
	ratingStr := c.Param("rating")

	rating, ratingErr := strconv.ParseFloat(ratingStr, 32)

	if ratingErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Provide valid rating",
		})
	}

	result, err := services.RateItem(id, float32(rating))
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find item",
		})
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
	}

	c.Status(http.StatusOK)
}

func SaveItem(c *gin.Context) {
	userId := c.Param("id")
	itemId := c.Query("item")

	user, _ := strconv.Atoi(userId)
	item, _ := strconv.Atoi(itemId)

	result := services.SaveItem(user, item)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.Status(http.StatusOK)
}

func RemoveSavedItem(c *gin.Context) {
	userId := c.Param("id")
	itemId := c.Query("item")
	result := services.RemoveSavedItem(userId, itemId)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.Status(http.StatusOK)
}

func GetSavedItems(c *gin.Context) {
	userId := c.Param("id")
	result, items := services.GetSavedItems(userId)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}
