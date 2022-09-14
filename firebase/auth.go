package snippets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"server/utils"

	"firebase.google.com/go/v4/auth"

	"io/ioutil"
	"net/http"
)

func init() {
	utils.LoadEnv()
}

const (
	verifyCustomTokenURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s"
)

type FirebaseAuth struct {
	auth *auth.Client
}

type PostBody struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

func (firebase *FirebaseApp) Connect() *FirebaseAuth {
	// Access auth service from the default app
	client, err := firebase.app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return &FirebaseAuth{
		auth: client,
	}
}

func postRequest(url string, req []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func GetIdToken(customtoken string) string {
	payload := map[string]interface{}{
		"token":             customtoken,
		"returnSecureToken": true,
	}

	req, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := postRequest(fmt.Sprintf(verifyCustomTokenURL, os.Getenv("FIREBASE_API_KEY")), req)

	if err != nil {
		fmt.Println(err)
	}
	var respBody struct {
		IDToken string `json:"idToken"`
	}
	if err := json.Unmarshal(resp, &respBody); err != nil {
		fmt.Println(err)
	}
	return respBody.IDToken

}

func (firebase *FirebaseApp) CreateCustomToken(ctx context.Context, customUID string) (string, error) {
	// [START create_custom_token_golang]
	client, err := firebase.app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(ctx, customUID)
	return token, err
}

func (firebase *FirebaseApp) VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error) {
	// [START verify_id_token_golang]
	client, err := firebase.app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)

	return token, err
}

func (fbClient *FirebaseAuth) CreateSmallUserForTest(ctx context.Context) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{})
	u, err := fbClient.auth.CreateUser(ctx, params)
	return u, err
}

func (fbClient *FirebaseAuth) CreateUser(ctx context.Context, email string, phonenumber string, password string, displayname string, photoUrl string, uid string) (*auth.UserRecord, error) {
	// [START create_user_golang]

	params := (&auth.UserToUpdate{}).
		Email(email).
		EmailVerified(false).
		PhoneNumber(phonenumber).
		Password(password).
		DisplayName(displayname).
		PhotoURL(photoUrl).
		Disabled(false)
	u, err := fbClient.auth.UpdateUser(ctx, uid, params)
	// [END create_user_golang]
	return u, err
}

func DeleteInBulk(ctx context.Context, client *auth.Client, users []string) bool {
	deleteUsersResult, err := client.DeleteUsers(ctx, users)
	if err != nil {
		log.Fatalf("error deleting users: %v\n", err)
	}

	log.Printf("Successfully deleted %d users", deleteUsersResult.SuccessCount)
	log.Printf("Failed to delete %d users", deleteUsersResult.FailureCount)
	for _, err := range deleteUsersResult.Errors {
		log.Printf("%v", err)
	}
	return true
}

func (fbClient *FirebaseAuth) UpdateUser(ctx context.Context, uid string) {
	params := (&auth.UserToUpdate{}).
		Email("user@example.com").
		EmailVerified(true).
		PhoneNumber("+15555550100").
		Password("newPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(true)
	u, err := fbClient.auth.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
}

func (fbClient *FirebaseAuth) UpdateUserPassword(ctx context.Context, uid string, passsword string) bool {
	params := (&auth.UserToUpdate{}).
		Password("newPassword")
	u, err := fbClient.auth.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	return true
}

func (fbClient *FirebaseAuth) DeleteUser(ctx context.Context, uid string) bool {
	err := fbClient.auth.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	return true
}
