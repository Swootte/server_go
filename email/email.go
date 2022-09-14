package email

import (
	"context"
	"os"
	"server/errors"
	"server/graph/model"
	"server/utils"
	"time"

	"github.com/hanzoai/gochimp3"
	"github.com/mailgun/mailgun-go/v4"
)

var mg mailgun.MailgunImpl

func init() {
	utils.LoadEnv()
	_mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
	_mg.SetAPIBase(mailgun.APIBaseEU)
	mg = *_mg
}

func initmailchimp() *gochimp3.ListResponse {
	client := gochimp3.New(os.Getenv("MAILCHIMP_KEY"))
	list, err := client.GetList(os.Getenv("MAILCHIMP_LIST_ID"), nil)
	if err != nil {
		errors.SaveError(context.Background(), err, "28", "email", "", "")
		panic(err)
	}

	return list
}

func SendEmail(email *model.EmailInput, ip string) (bool, error) {
	message := mg.NewMessage("contact@swootte.com", email.Name+" "+"Initial contact", email.Message, email.Message)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Send the message with a 10 second timeout
	_, _, err := mg.Send(ctx, message)

	if err != nil {
		errors.SaveError(ctx, err, "43", "email", "", "")
		return false, err
	}
	_result := true
	return _result, nil
}

func SendPDFQrCodeFile() {

}

func SendPDFQrCodeFilebis() {

}

func SendVerificationEmail() {

}

func SuscribeToNewsLetter(email string, ip string) (bool, error) {
	list := initmailchimp()
	req := &gochimp3.MemberRequest{
		EmailAddress: email,
		Status:       "subscribed",
	}

	if _, err := list.CreateMember(req); err != nil {
		errors.SaveError(context.Background(), err, "70", "email", "", "")
		return false, err
	}

	return true, nil
}
