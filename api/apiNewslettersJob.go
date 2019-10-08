package api

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiJobSendingNewsletters() {
	for {
		newsletterList := []database.Newsletters{}
		if err := config.Get().Postgres.Select(&newsletterList, "select * from newsletters where workflow = 'sending'"); err != nil {
			log.Println(err.Error())
		}

		//Loop through all eligible Users Newsletter,
		for _, newsletter := range newsletterList {

			//Check if either SMTP or SMS is enabled
			if newsletter.Sendsms || newsletter.Sendmail {

				var smtpAccounts []database.Smtps
				sqlQuery := "select * from smtps where workflow = 'enabled' and code = $1"
				if err := config.Get().Postgres.Select(&smtpAccounts, sqlQuery, newsletter.SMTP); err != nil {
					log.Println(err.Error())
				}

				smtpToUse := database.Smtps{}
				for _, smtpAcct := range smtpAccounts {
					cFormat := "Mon, 02 Jan 2006 15:04:05 WAT"
					smtpUpdatedate := smtpAcct.Updatedate.Format(cFormat)
					if utils.GetDifferenceInSeconds("", smtpUpdatedate) > int(smtpAcct.Delay) {
						smtpToUse = smtpAcct
						break
					}
				}

				var smsAccounts []database.Sms
				sqlQuery = "select * from sms where workflow = 'enabled' and code = $1"
				if err := config.Get().Postgres.Select(&smsAccounts, sqlQuery, newsletter.Sms); err != nil {
					log.Println(err.Error())
				}

				smsToUse := database.Sms{}
				for _, smsAcct := range smsAccounts {
					if smsAcct.Used < smsAcct.Units {
						smsToUse = smsAcct
						break
					}
					smsAcct.Workflow = "disabled"
					smsAcct.Description = fmt.Sprintf("SMS Units exhansted! Units:%v | Used:%v ",
						smsAcct.Units, smsAcct.Used)
					smsAcct.Update(smsAcct.ToMap())
				}

				//Get Recipients for the SMS or SMTP
				//Using smtpToUse.Rate, pick a certain number of recipients from followers
				rate := 1
				if int(smtpToUse.Rate) > 0 {
					rate = int(smtpToUse.Rate)
				}

				var recipientList []database.Followers
				sqlQuery = "select * from followers where workflow in ('pending','pending-email','pending-sms') and bucketid = $1 limit $2"
				if err := config.Get().Postgres.Select(&recipientList, sqlQuery, newsletter.ID, rate); err != nil {
					log.Println(err.Error())
				}

				recipientProfile := database.Profiles{}
				formSearch := new(database.SearchParams)

				if len(recipientList) > 0 {
					for _, recipient := range recipientList {
						formSearch.ID = recipient.UserID
						recipientProfile.GetByID(recipientProfile.ToMap(), formSearch)
						if recipientProfile.ID == 0 {
							recipient.Update(map[string]interface{}{
								"ID":          recipient.ID,
								"Workflow":    "failed",
								"Description": "Missing Contact",
							})
							continue
						}

						//Send Emails
						if newsletter.Sendmail && (recipient.Workflow == "pending" || recipient.Workflow == "pending-email") {
							if len(smtpAccounts) > 0 {
								emailRE := regexp.MustCompile(`^[a-zA-Z0-9][-_.a-zA-Z0-9]*@[a-zA-Z0-9.-]+\.[a-zA-Z0-9.-]+?$`)
								if newsletter.Sendmail && emailRE.MatchString(recipientProfile.Email) && smtpToUse.Workflow == "enabled" {
									if errorMsg, smtpTemplate := utils.GetTemplate(newsletter.Filepath); errorMsg == "" {
										var msgBytes bytes.Buffer

										mailer := utils.Email{}
										mailer.To = recipientProfile.Email
										mailer.From = newsletter.FromEmail
										mailer.Replyto = newsletter.Replyto
										mailer.FromName = newsletter.FromName

										jwtClaims := jwt.MapClaims{}
										jwtClaims["nid"] = newsletter.ID
										jwtClaims["uid"] = recipientProfile.ID
										jwtClaims["Email"] = recipientProfile.Email
										jwtClaims["Mobile"] = recipientProfile.Mobile
										jwtClaims["Fullname"] = recipientProfile.Fullname
										jwtClaims["Firstname"] = recipientProfile.Firstname
										jwtClaims["Lastname"] = recipientProfile.Lastname

										jwtClaims["exp"] = time.Now().AddDate(0, 0,
											newsletter.DaysValid).Unix()

										jwtToken, _ := utils.GenerateJWT(jwtClaims)

										newsletterRecipient := recipientProfile
										newsletterRecipient.Code = jwtToken + "/"

										if err := smtpTemplate.Execute(&msgBytes, newsletterRecipient); err == nil {
											if emailSubject, err := template.New("template").Parse(newsletter.Subject); err == nil {

												mailer.Message = msgBytes.String()
												if newsletter.URL != "" {
													mailer.Message += fmt.Sprintf(` <img src="%v/pixelthumb-%v-it.png"/>`,
														newsletter.URL, jwtToken)
												}

												var subjectBytes bytes.Buffer
												if err := emailSubject.Execute(&subjectBytes, recipientProfile); err == nil {
													mailer.Subject = subjectBytes.String()

													mySMTP := utils.SMTP{
														Port: int(smtpToUse.Port), Server: smtpToUse.Server,
														Username: smtpToUse.Username, Password: smtpToUse.Password,
													}

													sError := utils.SendEmail(mailer, mySMTP)
													if sError != "" {
														smtpToUse.Workflow = "disabled"
														smtpToUse.Description = sError
													} else {
														//Mark Follower as Sent
														recipient.Workflow = "sent"
														recipient.Description += "Email Sent "
													}
													//Update SMTP Account
													smtpToUse.Update(smtpToUse.ToMap())
												} else {
													newsletter.Workflow = "paused"
													newsletter.Description = "Please fix Newsletter Subject: " + newsletter.Subject
													newsletter.Update(newsletter.ToMap())
												}
											} else {
												newsletter.Workflow = "paused"
												newsletter.Description = "Error Parsing: Please fix Newsletter Subject: " + newsletter.Subject
												newsletter.Update(newsletter.ToMap())
											}
										} else {
											newsletter.Workflow = "paused"
											newsletter.Description = "Please fix Newsletter Email Message: "
											newsletter.Update(newsletter.ToMap())
										}
									} else {
										newsletter.Workflow = "paused"
										newsletter.Description = "Email Message is badly formed " + errorMsg
										newsletter.Update(newsletter.ToMap())
									}
								} else {
									if !emailRE.MatchString(recipientProfile.Email) {
										recipient.Workflow = "failed"
										recipient.Description += "| Invalid Email "
									}
								}
							} else {
								newsletter.Workflow = "paused"
								newsletter.Description = "Please enable SMTP Accounts with code: " + newsletter.SMTP
								newsletter.Update(newsletter.ToMap())
							}
						}

						// Send SMS Messages
						if newsletter.Sendsms && (recipient.Workflow == "pending" || recipient.Workflow == "pending-sms") {
							if len(smsAccounts) > 0 {
								mobile := regexp.MustCompile(`[^0-9]+`).
									ReplaceAllString(recipientProfile.Mobile, "")

								if newsletter.Sendsms && len(mobile) > 10 && newsletter.SmsMessage != "" &&
									smsToUse.Used <= smsToUse.Units {
									if smsTemplate, err := template.New("template").Parse(newsletter.SmsMessage); err == nil {
										var msgBytes bytes.Buffer
										if err := smsTemplate.Execute(&msgBytes, recipientProfile); err == nil {

											// smsURL := strings.Replace(smsToUse.URL, "%s", newsletter.SenderId, 1)
											go utils.SendSMSTwilio(newsletter.SenderId, mobile, msgBytes.String())

											//Mark Follower as Sent
											recipient.Workflow = "sent"
											recipient.Description += "| SMS Sent "

											//Update SMS Account
											smsToUse.Used++
											smsToUse.Update(smsToUse.ToMap())
										} else {
											newsletter.Workflow = "paused"
											newsletter.Description = "SMS Message parsing error: " + err.Error()
											newsletter.Update(newsletter.ToMap())
										}
									} else {
										newsletter.Workflow = "paused"
										newsletter.Description = "SMS Message is badly formed: " + err.Error()
										newsletter.Update(newsletter.ToMap())
									}
								} else {
									if len(mobile) > 10 {
										recipient.Workflow = "failed"
										recipient.Description += "| SMS Invalid Mobile"
									}

									if newsletter.SmsMessage == "" {
										newsletter.Workflow = "paused"
										newsletter.Description = "SMS Message is empty"
										newsletter.Update(newsletter.ToMap())
									}
								}
							} else {
								newsletter.Workflow = "paused"
								newsletter.Description = "Please enable SMS Accounts with code: " + newsletter.Sms
								newsletter.Update(newsletter.ToMap())
							}
						}

						if recipient.Workflow != "pending" {
							if smtpToUse.Workflow == "disabled" {
								recipient.Workflow = "pending-email"
							}
							recipient.Update(recipient.ToMap())
						}
					}
				} else {
					newsletter.Workflow = "completed"
					newsletter.Description = "Newsletter sent to all followers please check recipient list for each recipient status"
					newsletter.Update(newsletter.ToMap())
				}

			} else {
				newsletter.Workflow = "disabled"
				newsletter.Description = "Please enable SendMail or SendSMS else newsletter will be disabled"
				newsletter.Update(newsletter.ToMap())
				//Update This Newsletter
			}
		}
		//Pick 10 eligible Recipients each (if non Update newsletter as complete)

		<-time.Tick(time.Second * 10)
	}
}
