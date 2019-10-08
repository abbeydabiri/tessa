package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/justinas/alice"

	"tessa/database"
	"tessa/utils"
)

func apiHandlerSmtps(middlewares alice.Chain, router *Router) {
	router.Get("/api/smtp", middlewares.ThenFunc(apiSMTPGet))
	router.Post("/api/smtp", middlewares.ThenFunc(apiSMTPPost))
	router.Post("/api/smtp/search", middlewares.ThenFunc(apiSMTPSearch))
	router.Post("/api/smtp/search-code", middlewares.ThenFunc(apiSMTPSearch))

	router.Post("/api/smtp/test", middlewares.ThenFunc(apiSMTPTest))
}

func apiSMTPTest(httpRes http.ResponseWriter, httpReq *http.Request) {
	message := apiSecure(httpRes, httpReq)
	if message.Code == http.StatusOK {

		var formStruct struct {
			FromName, ReplyTo,
			SendTo, From,
			Subject, Message,
			Username, Password,
			Server string
			Port int
		}

		message.Message = "Email not sent"
		if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
			message.Code = http.StatusInternalServerError
			message.Message = "Error Decoding Form Values: " + err.Error()
		}

		if message.Code == http.StatusOK {
			message.Message = ""

			if formStruct.FromName == "" {
				message.Message += "From Name is missing!\n"
			}

			emailRE := regexp.MustCompile(`^[a-zA-Z0-9][-_.a-zA-Z0-9]*@[-_.a-zA-Z0-9]+?$`)
			if !emailRE.MatchString(formStruct.SendTo) {
				message.Message += "Send To is missing!\n"
			}

			if formStruct.From == "" {
				message.Message += "From Email is missing!\n"
			}

			if formStruct.Subject == "" {
				message.Message += "Subject is missing!\n"
			}

			if formStruct.Message == "" {
				message.Message += "Message is missing!\n"
			}

			if formStruct.Username == "" {
				message.Message += "Username is missing!\n"
			}

			if formStruct.Server == "" {
				message.Message += "Server is missing!\n"
			}

			if formStruct.Port == 0 {
				message.Message += "Port is invalid!\n"
			}

			if strings.HasSuffix(message.Message, "\n") {
				message.Message = message.Message[:len(message.Message)-2]
			}

			if message.Message == "" {
				//Send Test meail
				mailer := utils.Email{}
				mailer.To = formStruct.SendTo
				mailer.From = formStruct.From
				mailer.Replyto = formStruct.ReplyTo
				mailer.FromName = formStruct.FromName
				mailer.Subject = formStruct.Subject
				mailer.Message = formStruct.Message

				mySMTP := utils.SMTP{
					Port: formStruct.Port, Server: formStruct.Server,
					Username: formStruct.Username, Password: formStruct.Password,
				}

				if sError := utils.SendEmail(mailer, mySMTP); sError != "" {
					message.Message = fmt.Sprintf("Email Error: %s", sError)
				} else {
					message.Message = fmt.Sprintf("Email Sent to %v", mailer.To)
				}
				//Send Test meail
			}

		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiSMTPGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Smtps{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiSMTPPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiSecurePost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Smtps{}
		table.FillStruct(tableMap)

		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
			message.Body = table.ID
		}

		message.Body = table.ID
		message.Message = RecordSaved
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiSMTPSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Smtps{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}
		var searchList []interface{}
		searchResults := table.Search(table.ToMap(), formSearch)
		for _, result := range searchResults {
			tableMap := result.ToMap()
			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}
