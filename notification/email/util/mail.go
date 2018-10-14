package notification

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	//"log"
)

func SendMail(token string) {
	fmt.Println("---- ", token)
	from := mail.NewEmail("Saurabh", "skcse03@gmail.com")
	to := mail.NewEmail("pb", "testing mail")
	plainTextContent := " password Reset" + token
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>" + token
	subject := "reset password"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("API-key of Send grid")
	if response, err := client.Send(message); err == nil {
		fmt.Println(response)
	}
}
