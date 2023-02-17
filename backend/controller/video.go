package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miracleexotic/sa-64-example/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /videos
func CreateVideo(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Videos").InsertOne(context.TODO(), &video)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.InsertedID})
}

// GET /video/:id
func GetVideo(c *gin.Context) {
	var video entity.Video
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Collection("Videos").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": video})
}

// GET /videos
func ListVideos(c *gin.Context) {
	var videos []entity.Video
	cur, err := entity.DB().Collection("Videos").Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cur.All(context.TODO(), &videos)

	c.JSON(http.StatusOK, gin.H{"data": videos})
}

// DELETE /videos/:id
func DeleteVideo(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Videos").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.DeletedCount})
}

// PATCH /videos
func UpdateVideo(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Videos").UpdateOne(context.TODO(), bson.M{"_id": video.ID}, bson.M{
		"$set": bson.M{
			"name":     video.Name,
			"url":      video.Url,
			"owner_id": video.OwnerID,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.UpsertedID})
}
