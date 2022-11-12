package cloud

import (
	"context"
	"log"

	"firebase.google.com/go/v4/storage"
	"github.com/iugstav/colatech-api/internal/firebase"
)

func CloudStorage() *storage.Client {
	app := firebase.InitializeApp()

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
