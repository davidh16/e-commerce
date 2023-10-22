package services

import (
	"bytes"
	"context"
	"e-commerce/config"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
	"time"
)

func (s Service) UploadMediaToBucket(ctx context.Context, file multipart.File, path string) (string, error) {

	cfg := config.GetConfig()

	opt := option.WithCredentialsFile("password-lock-486ee-firebase-adminsdk-xtd5c-cc43257771.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return "", err
	}

	// Initialize Firebase Storage client
	client, err := app.Storage(context.Background())
	if err != nil {
		return "", err
	}

	// Create a storage reference
	storageRef, err := client.Bucket(cfg.StorageBucket)
	if err != nil {
		return "", err
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	object := storageRef.Object(path)

	// Upload the file to Firebase Storage
	wc := object.NewWriter(ctxWithTimeout)
	_, err = io.Copy(wc, file)
	if err != nil {
		return "", err
	}
	err = wc.Close()
	if err != nil {
		return "", err
	}

	return path, nil
}

func (s Service) DownloadMedia(ctx context.Context, path string) ([]byte, error) {

	cfg := config.GetConfig()

	opt := option.WithCredentialsFile("password-lock-486ee-firebase-adminsdk-xtd5c-cc43257771.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	// Initialize Firebase Storage client
	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	// Create a storage reference
	storageRef, err := client.Bucket(cfg.StorageBucket)
	if err != nil {
		return nil, err
	}

	object := storageRef.Object(path)
	if err != nil {
		return nil, err
	}

	reader, err := object.NewReader(context.Background())
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	// Copy the remote file's data to the buffer.
	_, err = io.Copy(&buffer, reader)
	if err != nil {
		log.Fatalf("Error reading file content: %v", err)
	}

	return buffer.Bytes(), nil
}
