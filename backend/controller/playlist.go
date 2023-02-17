package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miracleexotic/sa-64-example/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /playlists
func CreatePlaylist(c *gin.Context) {
	var playlist entity.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Playlists").InsertOne(context.TODO(), &playlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.InsertedID})
}

// GET /playlist/:id
func GetPlaylist(c *gin.Context) {
	var playlist entity.Playlist
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Collection("Playlists").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

// GET /playlist/watched/user/:id
func GetPlaylistWatchedByUser(c *gin.Context) {
	var playlist entity.Playlist
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Collection("Playlists").FindOne(context.TODO(), bson.M{"$and": []interface{}{bson.M{"owner_id": id}, bson.M{"title": "Watched"}}}).Decode(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": playlist})
}

// GET /playlists
func ListPlaylists(c *gin.Context) {
	var playlists []entity.Playlist
	cur, err := entity.DB().Collection("Playlists").Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cur.All(context.TODO(), &playlists)

	c.JSON(http.StatusOK, gin.H{"data": playlists})
}

// DELETE /playlists/:id
func DeletePlaylist(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Playlists").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.DeletedCount})
}

// PATCH /playlists
func UpdatePlaylist(c *gin.Context) {
	var playlist entity.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := entity.DB().Collection("Playlists").UpdateOne(context.TODO(), bson.M{"_id": playlist.ID}, bson.M{
		"$set": bson.M{
			"title":        playlist.Title,
			"owner_id":     playlist.OwnerID,
			"watch_videos": playlist.WatchVideos,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res.UpsertedID})
}
