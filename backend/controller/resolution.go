package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miracleexotic/sa-64-example/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /resolutions
func CreateResolution(c *gin.Context) {
	var resolution entity.Resolution
	if err := c.ShouldBindJSON(&resolution); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Resolutions").InsertOne(context.TODO(), &resolution)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.InsertedID})
}

// GET /resolution/:id
func GetResolution(c *gin.Context) {
	var resolution entity.Resolution
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Collection("Resolutions").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&resolution); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resolution})
}

// GET /resolutions
func ListResolutions(c *gin.Context) {
	var resolutions []entity.Resolution
	cur, err := entity.DB().Collection("Resolutions").Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cur.All(context.TODO(), &resolutions)

	c.JSON(http.StatusOK, gin.H{"data": resolutions})
}

// DELETE /resolutions/:id
func DeleteResolution(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Resolutions").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.DeletedCount})
}

// PATCH /resolutions
func UpdateResolution(c *gin.Context) {
	var resolution entity.Resolution
	if err := c.ShouldBindJSON(&resolution); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Resolutions").UpdateOne(context.TODO(), bson.M{"_id": resolution.ID}, bson.M{
		"$set": bson.M{
			"value": resolution.Value,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.UpsertedID})
}
