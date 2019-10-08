package utils

import (
	"bytes"
	"fmt"
	"log"

	"html/template"

	"tessa/config"
	"tessa/database"
)

//Email structure of mail
type Email struct {
	From, FromName, Replyto,
	To, Subject, Message string
	BCC, CC []string
}

func SendEmail(email Email, mySMTP SMTP) string {

	if email.To == "" || email.Subject == "" || email.Message == "" {
		return "either To, Subject or Message fields is blank/empty"
	}

	if mySMTP.Port == 0 || mySMTP.Server == "" || mySMTP.Username == "" || mySMTP.Password == "" {
		return "either Port, Server or Username or Password fields is/are blank/empty"
	}

	if email.From == "" {
		email.From = mySMTP.Username
	}

	if email.FromName == "" {
		email.FromName = mySMTP.Username
	}

	emailSender := fmt.Sprintf("%s <%s>", email.FromName, email.From)

	var messageList []EMailMessage
	messageList = append(messageList,
		EMailMessage{
			Attachment: "",
			To:         email.To,
			From:       emailSender,
			Cc:         email.CC, Bcc: email.BCC, Replyto: email.Replyto,
			Subject: email.Subject,
			Content: email.Message,
		})
	mailer := Mailer{mySMTP, messageList}

	//log.Printf(" - - -- - - - -- - -- - --- - \n Mail:  %v \n\n", mailer)
	// return ""

	sMessage := mailer.CheckMail()
	if len(sMessage) > 0 {
		log.Printf(sMessage)
		return sMessage
	}

	sMessage = mailer.SendMail()
	if len(sMessage) > 0 {
		sMessage = fmt.Sprintf("To: %v\nFrom: %v\nReply: %v\nSubject: %v\n %v",
			email.To, emailSender, email.Replyto, email.Subject, sMessage)
		log.Printf(sMessage)
		log.Printf(email.Message)
		return sMessage
	}

	return sMessage
}

//SendNewsletterEmail all mails get sent as newsletter i.e an existing template
func SendNewsletterEmail(emailStruct Email, mailTemplateValues interface{}, newsletterCode string) {

	newsletterList := []database.Newsletters{}
	sqlQuery := "select id, filepath, subject, fromemail, fromname, replyto, smtp from newsletters where sendmail = true and workflow = 'send' code = $1"
	if err := config.Get().Postgres.Select(&newsletterList, sqlQuery, newsletterCode); err != nil {
		log.Println(err.Error())
	}

	if len(newsletterList) > 0 {
		newsletter := newsletterList[0]
		if newsletter.ID == 0 {
			log.Println("Error finding Newsletter %+v", newsletter)
			return
		}

		errMessage, mailTemplate := GetTemplate(newsletter.Filepath)
		if errMessage != "" {
			log.Println("Error Getting Email Template " + errMessage)
		} else {
			var subjectBytes bytes.Buffer
			var messageBytes bytes.Buffer

			if err := mailTemplate.Execute(&messageBytes, mailTemplateValues); err != nil {
				log.Println("Error Generating Email Message " + err.Error())
			} else {

				if emailStruct.Message == "" {
					emailStruct.Message = messageBytes.String()
				} else {
					emailStruct.Message = messageBytes.String() + emailStruct.Message
				}

				if emailSubject, err := template.New("template").Parse(newsletter.Subject); err == nil {
					if err := emailSubject.Execute(&subjectBytes, mailTemplateValues); err == nil {
						emailStruct.Subject = subjectBytes.String()
					}
				}

				emailStruct.From = newsletter.FromEmail
				emailStruct.FromName = newsletter.FromName
				emailStruct.Replyto = newsletter.Replyto

				smtpToUse := database.Smtps{}
				formSearch := database.SearchParams{Field: "Code",
					Text: newsletter.SMTP, Workflow: "enabled"}
				smtpAccounts := smtpToUse.Search(smtpToUse.ToMap(), &formSearch)

				for _, smtpAcct := range smtpAccounts {
					cFormat := "Mon, 02 Jan 2006 15:04:05 WAT"
					smtpUpdatedate := smtpAcct.Updatedate.Format(cFormat)
					if GetDifferenceInSeconds("", smtpUpdatedate) >= int(smtpAcct.Delay) {
						smtpToUse = smtpAcct
						break
					}
					smtpToUse = smtpAcct
				}

				if smtpToUse.ID > 0 {
					utilsSMTP := SMTP{}
					utilsSMTP.Port = int(smtpToUse.Port)
					utilsSMTP.Server = smtpToUse.Server
					utilsSMTP.Username = smtpToUse.Username
					utilsSMTP.Password = smtpToUse.Password

					sError := SendEmail(emailStruct, utilsSMTP)
					if sError != "" {
						smtpToUse.Workflow = "disabled"
						smtpToUse.Description = sError
					}
					//Update SMTP Account
					smtpToUse.Update(smtpToUse.ToMap())
				} else {
					errorMsg := "NO valid SMTP Found to Use for mail -> " + emailStruct.To
					config.Get().Postgres.Exec("update newsletters set description = '$1' where id = $2", errorMsg, newsletter.ID)
					log.Printf("%s  - newsletter -> %+v \n", errorMsg, newsletter)
				}
			}
		}
	} else {
		log.Println("Email Newsletter not found " + newsletterCode)
	}
}
