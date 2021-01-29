package util

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type MySnsData struct {
	Facebook string
	Twitter string
	Instagram string
	Line string
}

func getFirebaseClient(ctx context.Context) (*firestore.Client, error) {
	projectID := "snsqrcodegenerator"

	opt := option.WithCredentialsFile("Users/hoge/秘密鍵.json")

	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SaveUserItem(m MySnsData) (MySnsData, error) {
	ctx := context.Background()

	client, err := getFirebaseClient(ctx)
	if err != nil {
		return m, err
	}

	_, err = client.Collection("users").Doc("1").Set(ctx, map[string]interface{}{
		"facebook" : m.Facebook,
		"twitter" : m.Twitter,
		"instagram" : m.Instagram,
		"line" : m.Line,
	})
	if err != nil {
		return m, err
	}
	return m, nil
}