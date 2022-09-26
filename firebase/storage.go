package snippets

import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

func (firebase *FirebaseApp) UploadFile(fileInput []byte, filename string) (string, error) {
	id := uuid.New()

	client, err := firebase.app.Storage(context.Background())
	if err != nil {
		return "", err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", err
	}

	object := bucket.Object(filename)
	writer := object.NewWriter(context.Background())
	writer.ObjectAttrs.ContentType = "application/pdf"
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(fileInput)); err != nil {
		return "", err
	}

	return bucket.SignedURL(filename, &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Hour),
	})

}

func (firebase *FirebaseApp) GetLinkFromFileName(filename string) (string, error) {
	client, err := firebase.app.Storage(context.Background())
	if err != nil {
		return "", err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", err
	}

	return bucket.SignedURL(filename, &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Hour),
	})
}
