package entity

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

var db *mongo.Database

func DB() *mongo.Database {
	return db
}

func ConnectDatabase() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	database_name := os.Getenv("DB_MONGO_NAME")
	uri := "mongodb://" + os.Getenv("DB_MONGO_USER") + ":" + os.Getenv("DB_MONGO_PASS") + "@" + os.Getenv("DB_MONGO_URL") + ":" + os.Getenv("DB_MONGO_PORT") + "/" + database_name

	fmt.Printf("[>] Connecting to MongoDB... ")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("[Failed]")
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("[Failed]")
		panic(err)
	}
	fmt.Println("[Connected]")

	db = client.Database(database_name)
	return client
}

func ClearDatabase() {
	fmt.Printf("[-] Clearing All Collections... ")

	collections := []string{"Users", "Videos", "Resolutions", "Playlists"}

	for i := 0; i < len(collections); i++ {
		if err := db.Collection(collections[i]).Drop(context.TODO()); err != nil {
			fmt.Println("[Failed]")
			log.Fatal(err)
		}
	}

	fmt.Println("[Completed]")
}

func CreateDatabase() {
	fmt.Println("[+] Creating init database... ")
	// Users
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	chanwit := User{
		ID:        primitive.NewObjectID(),
		Name:      "Chanwit",
		Email:     "chanwit@gmail.com",
		StudentID: "B6200000",
		Password:  string(password),
	}
	name := User{
		ID:        primitive.NewObjectID(),
		Name:      "Name",
		Email:     "name@example.com",
		StudentID: "B6200001",
		Password:  string(password),
	}

	res, err := db.Collection("Users").InsertMany(context.TODO(), []interface{}{
		&chanwit,
		&name,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" |- Users: %v\n", res.InsertedIDs)

	// Videos
	saLecture4 := Video{
		ID:      primitive.NewObjectID(),
		Name:    "SA Lecture 4",
		Url:     "https://youtu.be/123",
		OwnerID: chanwit.ID,
	}
	howTo := Video{
		ID:      primitive.NewObjectID(),
		Name:    "How to ...",
		Url:     "https://youtu.be/456",
		OwnerID: chanwit.ID,
	}
	helloWorld := Video{
		ID:      primitive.NewObjectID(),
		Name:    "Hello World with C",
		Url:     "https://youtu.be/789",
		OwnerID: name.ID,
	}

	res, err = db.Collection("Videos").InsertMany(context.TODO(), []interface{}{
		&saLecture4,
		&howTo,
		&helloWorld,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" |- Videos: %v\n", res.InsertedIDs)

	// Resolutions
	res360p := Resolution{
		ID:    primitive.NewObjectID(),
		Value: "360p",
	}
	res480p := Resolution{
		ID:    primitive.NewObjectID(),
		Value: "480p",
	}
	res720p := Resolution{
		ID:    primitive.NewObjectID(),
		Value: "720p",
	}

	res, err = db.Collection("Resolutions").InsertMany(context.TODO(), []interface{}{
		&res360p,
		&res480p,
		&res720p,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" |- Resolutions: %v\n", res.InsertedIDs)

	// Playlists
	watchedPlayListOfChanwit := Playlist{
		ID:      primitive.NewObjectID(),
		Title:   "Watched",
		OwnerID: chanwit.ID,
		WatchVideos: []WatchVideo{
			{
				WatchedTime:  time.Now(),
				VideoID:      saLecture4.ID,
				ResolutionID: res720p.ID,
			},
			{
				WatchedTime:  time.Now(),
				VideoID:      helloWorld.ID,
				ResolutionID: res720p.ID,
			},
		},
	}
	musicPlayListOfChanwit := Playlist{
		ID:          primitive.NewObjectID(),
		Title:       "Music",
		OwnerID:     chanwit.ID,
		WatchVideos: []WatchVideo{},
	}
	watchedPlayListOfName := Playlist{
		ID:      primitive.NewObjectID(),
		Title:   "Watched",
		OwnerID: name.ID,
		WatchVideos: []WatchVideo{
			{
				WatchedTime:  time.Now(),
				VideoID:      helloWorld.ID,
				ResolutionID: res480p.ID,
			},
		},
	}

	res, err = db.Collection("Playlists").InsertMany(context.TODO(), []interface{}{
		&watchedPlayListOfChanwit,
		&musicPlayListOfChanwit,
		&watchedPlayListOfName,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" |- Playlists: %v\n", res.InsertedIDs)
}

func SetupDatabase() {
	client := ConnectDatabase()
	fmt.Println(reflect.TypeOf(client))

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	ClearDatabase()
	CreateDatabase()

}
