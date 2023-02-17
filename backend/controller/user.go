package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miracleexotic/sa-64-example/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/asaskevich/govalidator"
)

// GET /users
// List all users
func ListUsers(c *gin.Context) {
	var users []entity.User
	cur, err := entity.DB().Collection("Users").Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cur.All(context.TODO(), &users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /user/:id
// Get user by id
func GetUser(c *gin.Context) {
	var user entity.User
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Collection("Users").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users
func CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Users").InsertOne(context.TODO(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.InsertedID})
}

// PATCH /users
func UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Users").UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"student_id": user.StudentID,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.UpsertedID})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Users").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.DeletedCount})
}
