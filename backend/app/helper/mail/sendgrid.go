package mail

import (
"encoding/json"
"errors"
"fmt"

sendgrid "github.com/sendgrid/sendgrid-go"
)

type Emailer interface{
	Send( *EmailContent) error
}

type EmailUser struct{
	Name 	string 	`json:"name"`
	Email 	string 	`json:"email"`
}


func (eu *EmailUser) String() string{
	c,_ := json.Marshal(eu)
	return string(c)
}

//EmailContent
type EmailContent struct{
	Id 					int 		`json:"id"`					//id of order
	Subject 			string		`json:"subject"`			//subject of email
	FromUser			*EmailUser	`json:"from_user"`			//Email of my company
	ToUser				*EmailUser	`json:"to_user"`			//Email of client
	PlainContent		string		`json:"plain_content"`		//Body content of email
	HtmlContent			string		`json:"html_content"`		//Html show
}

//func change EmailContent to string of json
func (ec *EmailContent) String() string{
	c,_ := json.Marshal(ec)
	return string(c)
}

//func Validate will check email content is available or not
func (ec *EmailContent) Validate() error{
	if ec==nil || ec.FromUser==nil || ec.ToUser==nil || ec.PlainContent==""{
		return errors.New("wrong content email")
	}
	return nil
}

//Struct Sendgrid implement send email to destination email via method Send
type Sendgrid struct{
	ApiKey 		string				`json:"api_key"`
	Client 		*sendgrid.Client
}

//fun NewSendgrid create new Sendgrid
func NewSendgrid(apiKey string) *Sendgrid{
	client := sendgrid.NewSendClient(apiKey)
	return &Sendgrid{
		ApiKey: apiKey,
		Client: client,
	}
}

//func Send mail
func (sg *Sendgrid) Send(ec *EmailContent) error{
	if err := ec.Validate(); err != nil {
		return errors.New("Error check validate when sending")
	}
	//code send email
	//TODO Un-comment that
	from := EmailUser{ec.FromUser.Name, ec.FromUser.Email}
	subject := ec.Subject
	to := EmailUser{ec.ToUser.Name, ec.ToUser.Email}
	plainTextContent := ec.PlainContent
	htmlContent := ec.HtmlContent
	message := EmailContent{1,subject,&from, &to, plainTextContent, htmlContent}
	response, err := sg.Client.send(message)
	if err != nil {
		fmt.Println("Cannot send email due to error: ", err)
		return err
	}
	fmt.Printf("Email sent with Response code: %v, Response body: %v, Response headers: %v\n", response.StatusCode, response.Body, response.Headers)
	return nil
}