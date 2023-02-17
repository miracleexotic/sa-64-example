package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/miracleexotic/sa-64-example/entity"
	"go.mongodb.org/mongo-driver/bson"
)

// POST /watch_videos
func CreateWatchVideo(c *gin.Context) {
	email, _ := c.Get("email")

	var user entity.User
	if err := entity.DB().Collection("Users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var watchvideo entity.WatchVideo
	var playlist entity.Playlist

	if err := c.ShouldBindJSON(&watchvideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(watchvideo)

	if err := entity.DB().Collection("Playlists").FindOne(context.TODO(), bson.M{"$and": []interface{}{bson.M{"owner_id": user.ID}, bson.M{"title": "Watched"}}}).Decode(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlist.WatchVideos = append(playlist.WatchVideos, watchvideo)

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(watchvideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
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

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// GET /watchvideo
func ListWatchVideos(c *gin.Context) {
	email, _ := c.Get("email")

	var user entity.User
	if err := entity.DB().Collection("Users").FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var playlists []entity.Playlist
	cur, err := entity.DB().Collection("Playlists").Find(context.TODO(), bson.M{"owner_id": user.ID})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cur.All(context.TODO(), &playlists)

	var watchvideodatas []entity.WatchVideoData
	for i := 0; i < len(playlists); i++ {
		for j := 0; j < len(playlists[i].WatchVideos); j++ {
			var resolution entity.Resolution
			if err := entity.DB().Collection("Resolutions").FindOne(context.TODO(), bson.M{"_id": playlists[i].WatchVideos[j].ResolutionID}).Decode(&resolution); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var video entity.Video
			if err := entity.DB().Collection("Videos").FindOne(context.TODO(), bson.M{"_id": playlists[i].WatchVideos[j].VideoID}).Decode(&video); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			watchvideodata := entity.WatchVideoData{
				ID:           ((i + 1) * 10) + j,
				WatchedTime:  playlists[i].WatchVideos[j].WatchedTime,
				ResolutionID: playlists[i].WatchVideos[j].ResolutionID,
				Resolution:   resolution,
				PlaylistID:   playlists[i].ID,
				Playlist:     playlists[i],
				VideoID:      playlists[i].WatchVideos[j].VideoID,
				Video:        video,
			}

			watchvideodatas = append(watchvideodatas, watchvideodata)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": watchvideodatas})
}
