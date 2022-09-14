package notification

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func LoadAllNotificationTest(c *client.Client, options client.Option, idToken string) (struct{ LoadNotification []model.Notification }, error) {
	var _response struct{ LoadNotification []model.Notification }
	err := c.Post(
		`
        query LoadNotification {
            loadNotification {
                _id
                text
                type
            }
        }
	`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
