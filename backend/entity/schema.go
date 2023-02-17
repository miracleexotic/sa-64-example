package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name" valid:"required~Name cannot be blank"`
	Email     string             `bson:"email,omitempty" json:"email" valid:"email"`
	StudentID string             `bson:"student_id,omitempty" json:"student_id" valid:"matches(^[BMD]\\d{7}$)"`
	Password  string             `bson:"password,omitempty" json:"password"`
}

type Video struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name,omitempty" json:"name"`
	Url     string             `bson:"url,omitempty" json:"url"`
	OwnerID primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id"`
}

type Resolution struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Value string             `bson:"value,omitempty" json:"value"`
}

type WatchVideo struct {
	WatchedTime  time.Time          `bson:"watched_time,omitempty" json:"watched_time" valid:"past~Watched time must be a past date"`
	VideoID      primitive.ObjectID `bson:"video_id,omitempty" json:"video_id"`
	ResolutionID primitive.ObjectID `bson:"resolution_id,omitempty" json:"resolution_id"`
}

type WatchVideoData struct {
	ID          int       `json:"id"`
	WatchedTime time.Time `bson:"watched_time,omitempty" json:"watched_time" valid:"past~Watched time must be a past date"`

	ResolutionID primitive.ObjectID `bson:"resolution_id,omitempty" json:"resolution_id"`
	Resolution   Resolution         `bson:"resolution,omitempty" json:"resolution"`

	PlaylistID primitive.ObjectID `bson:"playlist_id,omitempty" json:"playlist_id"`
	Playlist   Playlist           `bson:"playlist,omitempty" json:"playlist"`

	VideoID primitive.ObjectID `bson:"video_id,omitempty" json:"video_id"`
	Video   Video              `bson:"video,omitempty" json:"video"`
}

type Playlist struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	OwnerID     primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id"`
	WatchVideos []WatchVideo       `bson:"watch_videos,omitempty" json:"watch_videos"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.After(t)
	})
	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.Before(time.Time(t))
	})
}
