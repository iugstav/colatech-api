package auth

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/iugstav/colatech-api/internal/firebase"
)

func GetAuthClient() *auth.Client {
	app, err := firebase.InitializeApp().Auth(context.Background())
	if err != nil {
		panic(err)
	}

	return app
}
