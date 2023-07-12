package utils

import (
	"fmt"
	"log"
	"github.com/Bassiouni/money-mate-api/contracts"
	"github.com/Bassiouni/money-mate-api/db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	u, err := models.All()

	if err != nil {
		log.Printf("Problem with fetching the users: %v\n", err)
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusFound, u)
}

func CreateUser(c *gin.Context) {
	var u contracts.User

	if err := c.Bind(&u); err != nil {
		log.Printf("Couldn't bind object: %v\n", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "provide the correct JSON object",
		})

		return
	}

	fmt.Println("User Creation", u)

	u, err := models.CreateNewUser(&u)

	if err != nil {
		log.Printf("Couldn't create the user %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Couldn't create the user",
		})

		return
	}

	c.JSON(http.StatusCreated, u)
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Invalid ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid ID",
		})
		return
	}

	u, err := models.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusFound, u)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	
	u, err := models.FindByEmail(email)
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusFound, u)
}
