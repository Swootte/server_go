package snippets

import (
	"context"

	"firebase.google.com/go/v4/messaging"
)

func (firebase *FirebaseApp) SendNotificationMEssage(token string, title string, body string, _type string, transactionId string) (bool, error) {
	//create an id
	client, err := firebase.app.Messaging(context.Background())
	if err != nil {
		return false, err
	}

	client.Send(context.Background(), &messaging.Message{
		Data: map[string]string{
			"type":          _type,
			"transactionId": transactionId,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Data:     map[string]string{},
			Notification: &messaging.AndroidNotification{
				Sound:             "default",
				NotificationCount: new(int),
			},
			FCMOptions: &messaging.AndroidFCMOptions{},
		},
		Webpush: &messaging.WebpushConfig{},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "5",
				"apns-topic":    "io.flutter.plugins.firebase.messaging",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge:            new(int),
					Sound:            "default",
					ContentAvailable: true,
				},
			},
		},
		Token: token,
	})

	return true, nil
}
