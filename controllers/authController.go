package controllers

import (
	"net/http"
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE username=golang", input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credential"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged in"})
}
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
		return
	}
	user.Password = string(hash)
	_, err = database.DB.NamedExec(`INSERT INTO USERS(USERNAME, USER_PASSWORD) VALUES (:username, :password)`, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in user creation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created sucessfully"})
}
