package snippets

import (
	"context"
	"log"
	"os"
	"server/utils"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

/*
initialise firebase App
*/

func init() {
	utils.LoadEnv()
}

type FirebaseApp struct {
	app *firebase.App
}

func Connect() *FirebaseApp {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT")))
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_ID"), StorageBucket: os.Getenv("STORAGEBUCKET")}
	app, err := firebase.NewApp(context.Background(), config, opt)

	if err != nil {
		log.Fatalf("errir initializing app: %v\n", err)
	}
	return &FirebaseApp{
		app: app,
	}
}
