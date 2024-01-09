package sendgridwrap

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	host = "https://api.sendgrid.com"
)

const (
	LearnToCodeFYIListID = "ca9ecbeb-2741-4365-9e5e-600058798398"
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
	const templateID = "d-6e010242fdb54c819d93ce43f0e3a51c"
	const fromEmail = "hello@learntocode.fyi"
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
	resp, err := sendgrid.MakeRequest(request)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("bad response from sendgrid. code: %v, body %v", resp.StatusCode, resp.Body)
	}

	return nil
}

// UpsertContact -
// https://docs.sendgrid.com/api-reference/contacts/add-or-update-a-contact
func (c Client) UpsertContact(email string, listIDs []string) error {
	if c.Platform != "PROD" {
		log.Println("Would upsert contact:")
		log.Println("To:", email)
		log.Println("List IDs:", listIDs)
		return nil
	}

	request := sendgrid.GetRequest(c.APIKey, "/v3/marketing/contacts", host)
	request.Method = "PUT"

	type contactModel struct {
		Email string `json:"email"`
	}

	type reqModel struct {
		ListIDs  []string       `json:"list_ids,omitempty"`
		Contacts []contactModel `json:"contacts"`
	}

	contact := contactModel{
		Email: email,
	}

	dat, err := json.Marshal(reqModel{
		ListIDs:  listIDs,
		Contacts: []contactModel{contact},
	})
	if err != nil {
		return err
	}
	request.Body = dat
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if response.StatusCode > 299 {
		return fmt.Errorf("bad response from sendgrid. code: %v, body %v", response.StatusCode, response.Body)
	}
	return nil
}
