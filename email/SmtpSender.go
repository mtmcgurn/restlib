// Copyright 2019 Reaction Engineering International. All rights reserved.
// Use of this source code is governed by the MIT license in the file LICENSE.txt.

package email

import (
	"encoding/json"
	"github.com/domodwyer/mailyak"
	"github.com/reaction-eng/restlib/configuration"
	"github.com/reaction-eng/restlib/utils"
	"html/template"
	"log"
	"net/smtp"
	"path/filepath"
	"time"
)

/**
Simple struct for email
*/
type SmtpSender struct {
	smtpServer   string
	smtpUser     string
	smtpPassword string
	smtpFrom     string
	smtpPort     string
}

//Provide a method to make a new AnimalRepoSql
func NewSmtpSender(configFile ...string) *SmtpSender {

	//Load up the config
	config, err := configuration.NewConfiguration(configFile...)

	if err != nil {
		log.Fatal(err)
	}

	sender := SmtpSender{
		smtpServer:   config.GetStringFatal("smtp_server"),
		smtpPort:     config.GetStringFatal("smtp_port"),
		smtpUser:     config.GetStringFatal("smtp_user"),
		smtpPassword: config.GetStringFatal("smtp_password"),
		smtpFrom:     config.GetStringFatal("smtp_from"),
	}

	return &sender

}

/**
Get all of the news
*/
func (repo *SmtpSender) SendEmail(email *HeaderInfo, body string, attachments map[string][]*utils.Base64File) error {

	// Create a new email - specify the SMTP host and auth
	mail := mailyak.New(repo.smtpServer+repo.smtpPort,
		smtp.PlainAuth("", repo.smtpUser, repo.smtpPassword, repo.smtpServer)) //authentication

	//Set the to info
	mail.To(email.To...)
	//If there are any bcc
	if email.Bcc != nil {
		mail.Bcc(email.Bcc...)
	}
	mail.Subject(email.Subject)
	mail.From(repo.smtpFrom)
	if len(email.ReplyTo) > 0 {
		mail.ReplyTo(email.ReplyTo)
	}

	//Set the body
	mail.Plain().Set(body)

	//Now Send
	return mail.Send()

}

func formatInTimeZone(dateTime *time.Time, timeZone string, format string) string {
	if dateTime == nil {
		return ""
	}

	currentTimeZone, _ := time.LoadLocation(timeZone)
	//Convert the time
	timeInZone := dateTime.In(currentTimeZone)

	return timeInZone.Format(format)

}

/**
Get all of the news
*/
func (repo *SmtpSender) SendEmailTemplateString(email *HeaderInfo, templateString string, data interface{}, attachments map[string][]*utils.Base64File) error {
	// Create a new email - specify the SMTP host and auth
	mail := mailyak.New(repo.smtpServer+repo.smtpPort,
		smtp.PlainAuth("", repo.smtpUser, repo.smtpPassword, repo.smtpServer)) //authentication

	//Set the to info
	mail.To(email.To...)
	//If there are any bcc
	if email.Bcc != nil {
		mail.Bcc(email.Bcc...)
	}
	mail.Subject(email.Subject)
	mail.From(repo.smtpFrom)
	if len(email.ReplyTo) > 0 {
		mail.ReplyTo(email.ReplyTo)
	}

	//Define a template function map for general time
	funcMap := template.FuncMap{
		"now":              time.Now,
		"formatInTimeZone": formatInTimeZone,
	}

	//Execute the table file
	t := template.New("Basic Table Template").Funcs(funcMap)

	//Parse the file
	t, err := t.Parse(templateString)
	if err != nil {
		return err
	}

	//Now add the html table
	err = t.Execute(mail.HTML(), data)
	if err != nil {
		return err
	}

	//Set an error
	tryJsonString, _ := json.Marshal(data)
	mail.Plain().Set(string(tryJsonString))

	//March over each attachment and add it
	for _, values := range attachments {
		for _, value := range values {
			//Save it to the mail
			mail.Attach(value.GetName(), value.GetDataReader())
		}
	}

	//Now Send
	return mail.Send()
}

/**
Get all of the news
*/
func (repo *SmtpSender) SendEmailTemplateFile(email *HeaderInfo, templateFile string, data interface{}, attachments map[string][]*utils.Base64File) error {
	// Create a new email - specify the SMTP host and auth
	mail := mailyak.New(repo.smtpServer+repo.smtpPort,
		smtp.PlainAuth("", repo.smtpUser, repo.smtpPassword, repo.smtpServer)) //authentication

	//Set the to info
	mail.To(email.To...)

	//If there are any bcc
	if email.Bcc != nil {
		mail.Bcc(email.Bcc...)
	}

	mail.Subject(email.Subject)
	mail.From(repo.smtpFrom)
	if len(email.ReplyTo) > 0 {
		mail.ReplyTo(email.ReplyTo)
	}

	//Define a template function map for general time
	funcMap := template.FuncMap{
		"now":              time.Now,
		"formatInTimeZone": formatInTimeZone,
	}

	//Parse the file
	t, err := template.New(filepath.Base(templateFile)).Funcs(funcMap).ParseFiles(templateFile)
	if err != nil {
		return err
	}
	//Now add the html table
	err = t.Execute(mail.HTML(), data)
	if err != nil {
		return err
	}

	//Set an error
	tryJsonString, _ := json.Marshal(data)
	mail.Plain().Set(string(tryJsonString))

	//March over each attachment and add it
	for _, values := range attachments {
		for _, value := range values {
			//Save it to the mail
			mail.Attach(value.GetName(), value.GetDataReader())
		}
	}

	//Now Send
	return mail.Send()
}

/**
Get all of the news
*/
func (repo *SmtpSender) SendEmailTable(email *HeaderInfo, tableData TableInfo, attachments map[string][]*utils.Base64File) error {

	// Create a new email - specify the SMTP host and auth
	mail := mailyak.New(repo.smtpServer+repo.smtpPort,
		smtp.PlainAuth("", repo.smtpUser, repo.smtpPassword, repo.smtpServer)) //authentication

	//Set the to info
	mail.To(email.To...)
	//If there are any bcc
	if email.Bcc != nil {
		mail.Bcc(email.Bcc...)
	}
	mail.Subject(email.Subject)
	mail.From(repo.smtpFrom)
	if len(email.ReplyTo) > 0 {
		mail.ReplyTo(email.ReplyTo)
	}
	//Execute the table file
	t := template.New("Basic Table Template")

	//Add the required functions
	t.Funcs(template.FuncMap{
		"GetTable": TableizeData,
		"GetTitle": GetTableTitle,
	})

	//Parse the file
	t, err := t.Parse(getTableHtml())
	if err != nil {
		return err
	}

	//Now add the html table
	err = t.Execute(mail.HTML(), tableData)
	if err != nil {
		return err
	}

	//Set an error
	mail.Plain().Set("HTML Email Required")

	//March over each attachment and add it
	for _, values := range attachments {
		for _, value := range values {

			//Save it to the mail
			mail.Attach(value.GetName(), value.GetDataReader())

		}
	}

	//Now Send
	return mail.Send()

}

/**
Function to tableize the data
*/
func getTableHtml() string {

	return `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<style>
				.header{
  					background-color: #aed957;
  					font-size: 18px;
  					text-align: center;
  					font-weight: bold;
				}
				.title{
  					text-align: left;
  					font-size: 15px;
					font-weight: bold;
  					background-color: gray;

				}
				.content{
  			  		background-color: white;

				}
			</style>
		</head>
		<body>
		<h1>{{. | GetTitle}}</h1>
		{{. | GetTable }}

		</body>
		</html>

	`

}

/**
Function to tableize the data
*/
func GetTableTitle(args ...interface{}) string {

	//check to see if it is a tableInfo
	tableInfo, ok := args[0].(TableInfo)

	//If is is a table
	if ok {
		return tableInfo.GetTitle()
	} else {
		return "Unknown Table Title"
	}

}

/**
Function to tableize the data
*/
func TableizeData(args ...interface{}) template.HTML {
	//check to see if it is a tableInfo
	tableInfo, ok := args[0].(TableInfo)

	//If is is a table
	if ok {
		html := template.HTML(tableizeTalbeInfo(tableInfo))
		return html
	} else {
		return "Unknown Table Title"
	}
}

/**
Function to tableize the data
*/
func tableizeTalbeInfo(info TableInfo) string {
	//Check to see if it has children
	html := ""

	//Add the table header
	html += `<table width="99%" border="0" cellpadding="1" cellspacing="0" bgcolor="#EAEAEA">`

	//Now add the header
	html += `<tr class="header"><td>` + info.GetTitle() + `</td></tr>`

	//Now march over each data, if it contains another table add it
	for _, childInfo := range info.GetChildren() {
		//If it is a node just add it
		if childInfo.IsNode() {
			html += `<tr><td>`
			html += tableizeTalbeInfo(childInfo)
			html += `</td></tr>`

		} else {
			//Add the table row
			html += `<tr class="title"><td><strong>` + childInfo.GetTitle() + `</strong></td></tr>`

			//If it is a node add the children
			html += `<tr><td>` + childInfo.GetValue() + `</td></tr>`

		}

	}

	//Close the ui segment
	html += `</table>`

	return html
}
