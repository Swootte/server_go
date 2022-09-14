package user

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
)

const (
	User_adminUIDTEST = "5fDqKkcrrXgg9g8KfosQaaPM1b12"
	Admin_pinTEST     = "111111"
	UserPinCode       = "140403"
)

func CreateUserTest(c *client.Client, options client.Option, idToken string) (*struct{ CreateUser *model.UserCreated }, error) {
	var _response struct {
		CreateUser *model.UserCreated
	}
	fcmtoken := "sqdsq"
	invitedBy := ""
	_user := model.UserInput{
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		Email:       gofakeit.Email(),
		PinCode:     UserPinCode,
		Password:    "sqdqsdqsd",
		Phonenumber: "+1" + gofakeit.Phone(),
		Token:       "sqdqsd",
		Country:     "US",
		DisplayName: &fcmtoken,
		FcmToken:    &fcmtoken,
		InvitedBy:   &invitedBy,
		Adress: &model.AdressInput{
			Title: &gofakeit.Address().Street,
			Location: &model.LocationInput{
				Latitude:  &gofakeit.Address().Latitude,
				Longitude: &gofakeit.Address().Longitude,
			},
			IsChosed: new(bool),
			City:     &gofakeit.Address().City,
		},
		BirthDate: &model.BirthDate{
			Day:   10,
			Month: 6,
			Year:  1990,
			Iso:   "dd",
		},
	}

	err := c.Post(
		`mutation CreateUser($user: UserInput) {
			createUser(user: $user){
				user {
					_id
					first_name
					last_name
					keypair{
						publicKey
						secretKey
					}
				}
				customToken
			}
		}
		`,
		&_response,
		client.Var("user", _user),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return &_response, err
}

func LoadBalanceTest(c *client.Client, options client.Option, idToken string) (*struct{ LoadBalance model.Wallet }, error) {
	var _response struct{ LoadBalance model.Wallet }
	err := c.Post(
		`
		query LoadBalance {
			loadBalance {
					amount
					address
					isFrozen
				}
			}
		`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return &_response, err
}

func DeleteUserTest(c *client.Client, options client.Option, idToken string) (struct{ DeleteUser bool }, error) {
	var _response struct{ DeleteUser bool }
	err := c.Post(
		`
		mutation DeleteUser {
			deleteUser
		}
		`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func UpdateFcmTokenTest(c *client.Client, options client.Option, idToken string, fcmToken string) (*struct{ UpdateFcmToken bool }, error) {
	var _response struct{ UpdateFcmToken bool }
	err := c.Post(
		`
		mutation UpdateFcmToken($fcmToken: String) {
			updateFcmToken(fcmToken: $fcmToken)
		}
		`,
		&_response,
		client.Var("fcmToken", fcmToken),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return &_response, err
}

func ChangePinCodeTest(c *client.Client, options client.Option, idToken string, newPin string) (struct{ ChangePinCode bool }, error) {
	var _response struct{ ChangePinCode bool }
	err := c.Post(
		`
		mutation ChangePinCode($newPin: String) {
			changePinCode(newPin: $newPin)
		}
		`,
		&_response,
		client.Var("newPin", newPin),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func UpdateProfilPictureTest(c *client.Client, options client.Option, idToken string, link string) (struct{ UpdateProfilePicture bool }, error) {
	var _response struct{ UpdateProfilePicture bool }
	err := c.Post(
		`
		mutation UpdateProfilePicture($link: String!) {
			updateProfilePicture(link: $link)
		}
		`,
		&_response,
		client.Var("link", link),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func UserExistTest(c *client.Client, options client.Option, idToken string) (*struct{ UsersExist model.User }, error) {
	var _response struct {
		UsersExist model.User
	}
	err := c.Post(
		`
		query UsersExist {
			usersExist {
				_id
				photoUrl
				fcmToken
				permissions
			}
		}
		`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return &_response, err
}
