package auth

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/iugstav/colatech-api/pkg/user"
)

func CreateUser(data *user.CreateFirebaseUserData, client *auth.Client) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		UID(data.ID).
		Email(data.Email).
		Password(data.BrutePassword).
		PhotoURL(data.ImageURL)

	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func DeleteUser(id string, client *auth.Client) error {
	err := client.DeleteUser(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}
