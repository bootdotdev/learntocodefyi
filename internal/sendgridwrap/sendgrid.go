package sendgridwrap

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	host = "https://api.sendgrid.com"
)

// Client -
type Client struct {
	APIKey   string
	Platform string
}

// NewClient -
func NewClient(apiKey, platform string) Client {
	return Client{
		APIKey:   apiKey,
		Platform: platform,
	}
}

// SendMagicLink -
func (c Client) SendMagicLink(toEmail, toName, loginLink string) error {
	const templateID = "d-a8edcea9806b4cdbbb72ff85b8efd93c"
	const fromEmail = "no-reply@learntocode.fyi"
	const fromName = "LearnToCode.fyi"

	if c.Platform != "PROD" {
		log.Println("Would send email:")
		log.Println("To:", toEmail)
		log.Println("To Name:", toName)
		log.Println("Login Link:", loginLink)
		return nil
	}

	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail(fromName, fromEmail))

	m.SetTemplateID(templateID)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(toName, toEmail),
	}
	p.AddTos(tos...)
	p.SetDynamicTemplateData("loginLink", loginLink)
	m.AddPersonalizations(p)

	mailSettings := mail.NewMailSettings()
	bypassListManagementSetting := mail.NewSetting(true)
	mailSettings.SetBypassListManagement(bypassListManagementSetting)
	m.SetMailSettings(mailSettings)

	falseVal := false
	trueVal := true
	m.SetTrackingSettings(&mail.TrackingSettings{
		SubscriptionTracking: &mail.SubscriptionTrackingSetting{
			Enable: &falseVal,
		},
		BypassListManagement: &mail.Setting{
			Enable: &trueVal,
		},
	})

	request := sendgrid.GetRequest(c.APIKey, "/v3/mail/send", host)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.MakeRequest(request)
	if err != nil {
		return err
	}
	return nil
}
