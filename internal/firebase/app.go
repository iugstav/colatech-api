package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitializeApp() *firebase.App {
	envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal(envError)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	pathToKey := fmt.Sprintf("%s/%s", dir, os.Getenv("PATH_TO_FIREBASE_KEY"))

	config := &firebase.Config{
		StorageBucket: os.Getenv("FIREBASE_STORAGE_BUCKET_NAME"),
	}
	opt := option.WithCredentialsFile(pathToKey)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	return app
}
