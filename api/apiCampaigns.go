package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"strings"
	"time"

	"github.com/justinas/alice"

	jwt "github.com/dgrijalva/jwt-go"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerCampaigns(middlewares alice.Chain, router *Router) {
	router.Get("/api/campaigns", middlewares.ThenFunc(apiCampaignGet))
	router.Post("/api/campaigns", middlewares.ThenFunc(apiCampaignPost))
	router.Post("/api/campaigns/search", middlewares.ThenFunc(apiCampaignSearch))
	router.Post("/api/campaigns/duplicate", middlewares.ThenFunc(apiCampaignDuplicate))
	router.Post("/api/campaigns/link", middlewares.ThenFunc(apiCampaignLink))
}

func apiCampaignDuplicate(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	message.Message = "Could not Duplicate Campaign"
	if message.Code == http.StatusOK {
		table := database.Campaigns{}
		table.GetByID(table.ToMap(), formSearch)

		table.ID = 0
		table.Workflow = "draft"
		table.Title += " - DUPLICATE"

		table.Code = strings.ToLower(fmt.Sprintf("CMP%v-%v-%v", formSearch.OwnerID,
			time.Now().Format("060102"), utils.RandomString(3)))

		fileBytes, _ := config.Asset(table.Campaign)
		if fileBytes != nil {
			table.Campaign = utils.SaveFileToPath("campaign.zip", "campaigns/"+table.Code, fileBytes)
			utils.ZipExtract(table.Campaign, strings.Replace(table.Campaign, "campaign.zip", "", 1))
		}

		table.Create(table.ToMap())
		message.Body = table.ID
		message.Message = "Campaign Duplicated!!"
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiCampaignLink(httpRes http.ResponseWriter, httpReq *http.Request) {
	message := apiGeneric(httpRes, httpReq)
	if message.Code == http.StatusOK {

		var formStruct struct {
			TargetID, CampaignID uint64
			ExpiryDate, ExpiryTime,
			Prefix, Sendmailto,
			AccessToken string

			Sendmail, Random,
			Shorten bool
		}

		if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
			message.Message = err.Error() + " \n"
		}

		if formStruct.ExpiryDate == "" {
			message.Message += "Expiry Date is Missing \n"
		}

		if formStruct.ExpiryTime == "" {
			message.Message += "Expiry Date is Missing \n"
		}

		if formStruct.CampaignID == 0 {
			message.Message += "No Campaign Found \n"
		}

		if strings.HasSuffix(message.Message, "\n") {
			message.Code = http.StatusInternalServerError
			message.Message = message.Message[:len(message.Message)-2]
		}

		tableProfile := database.Profiles{}
		if message.Code == http.StatusOK && formStruct.TargetID > uint64(0) {
			tableProfile.GetByID(tableProfile.ToMap(), &database.SearchParams{ID: formStruct.TargetID})
		}

		tableCampaign := database.Campaigns{}
		if message.Code == http.StatusOK && formStruct.CampaignID > uint64(0) {
			tableCampaign.GetByID(tableCampaign.ToMap(), &database.SearchParams{ID: formStruct.CampaignID})
		}

		if message.Code == http.StatusOK && tableCampaign.ID > 0 && tableProfile.ID > 0 {
			jwtClaims := jwt.MapClaims{}
			if tableProfile.Email != "" {
				jwtClaims["Email"] = tableProfile.Email
			}
			if tableProfile.Mobile != "" {
				jwtClaims["Mobile"] = tableProfile.Mobile
			}
			if tableProfile.Fullname != "" {
				jwtClaims["Fullname"] = tableProfile.Fullname
			}
			if tableProfile.Title != "" {
				jwtClaims["Title"] = tableProfile.Title
			}
			if tableProfile.Firstname != "" {
				jwtClaims["Firstname"] = tableProfile.Firstname
			}
			if tableProfile.Lastname != "" {
				jwtClaims["Lastname"] = tableProfile.Lastname
			}

			if formStruct.Random {
				jwtClaims["rand"] = utils.RandomString(4)
			}
			jwtClaims["cid"] = formStruct.CampaignID
			mergedDate := utils.DateTimeMerge(formStruct.ExpiryDate,
				formStruct.ExpiryTime, config.Get().Timezone)
			jwtClaims["exp"] = mergedDate.Unix()

			jwtToken, _ := utils.GenerateJWT(jwtClaims)
			campaignTitle := "Click Here"
			campaignURL := fmt.Sprintf(`%s/%s/`, formStruct.Prefix, jwtToken)

			//Shorten Link if requested and access token exists
			if formStruct.Shorten && formStruct.AccessToken != "" {
				apiURL := fmt.Sprintf("https://api-ssl.bitly.com/v3/shorten?format=txt&access_token=%s&longURL=%s",
					url.QueryEscape(formStruct.AccessToken), url.QueryEscape(campaignURL))
				apiHttpReq, _ := http.NewRequest("GET", apiURL, nil)
				apiClient := &http.Client{}
				apiResp, apiErr := apiClient.Do(apiHttpReq)
				if apiErr != nil {
					message.Message = "URL Shortener Error: " + apiErr.Error()
					formStruct.Sendmail = false
				} else {
					apiRespBody, _ := ioutil.ReadAll(apiResp.Body)
					campaignTitle = string(apiRespBody)
					campaignURL = campaignTitle
					apiResp.Body.Close()
				}
			}
			//Shorten Link if requested and access token exists

			message.Code = http.StatusOK
			message.Body = map[string]string{"url": campaignURL, "title": campaignTitle}

			//Send Generated Link via Custom Email
			/*
				smtpToUse := database.Smtps{}
				emailRE := regexp.MustCompile(`^[a-zA-Z0-9][-_.a-zA-Z0-9]*@[-_.a-zA-Z0-9]+?$`)
				if tableCampaign.SMTP != "" && formStruct.Sendmail && emailRE.MatchString(formStruct.Sendmailto) {

					var smtpAccounts []database.Smtps
					sqlQuery := "select * from smtps where workflow = 'enabled' and campaignid = $1 and ownerid = $2 and code = $3"
					config.Get().Postgres.Select(&smtpAccounts, sqlQuery, tableCampaign.CampaignID, tableCampaign.OwnerID, tableCampaign.Smtp)

					for _, smtpAcct := range smtpAccounts {
						cFormat := "Mon, 02 Jan 2006 15:04:05 WAT"
						smtpUpdatedate := smtpAcct.Updatedate.Format(cFormat)
						smtpToUse = smtpAcct
						if utils.GetDifferenceInSeconds("", smtpUpdatedate) > int(smtpAcct.Delay) {
							break
						}
					}

					if smtpToUse.ID > uint64(0) {
						var messageBytes bytes.Buffer
						var emailTemplate *template.Template
						templateFilePath := ""
						sqlQuery := "select filepath from newsletter where sendmail = true and workflow = 'send' and code = 'email_campaign_link' and campaignid = $1 and ownerid = $2 limit 1"
						config.Get().Postgres.Get(&templateFilePath, sqlQuery, tableCampaign.CampaignID, tableCampaign.OwnerID)
						message.Message, emailTemplate = utils.GetTemplate(templateFilePath)

						//Create Custom PixelTracker jwtToken
						jwtClaimsPixel := jwt.MapClaims{}
						jwtClaimsPixel["URL"] = campaignURL
						jwtClaimsPixel["Email"] = fmt.Sprintf("%s openned: campaign-link-mail %s ",
							tableCampaign.SmtpTo, tableCampaign.Title)
						jwtClaimsPixel["exp"] = mergedDate.Unix()
						jwtTokenPixel, _ := utils.GenerateJWT(jwtClaimsPixel)
						//Create Custom PixelTracker jwtToken

						var msgStruct struct {
							URL string
						}
						msgStruct.URL = campaignURL

						if message.Message == "" {
							if err := emailTemplate.Execute(&messageBytes, msgStruct); err != nil {
								log.Println("Error Generating Email Message " + err.Error())
								message.Message = fmt.Sprintf("ERROR GENERATING EMAIL - Please Fix!!!\nError: %s", err.Error())
							} else {
								mailer := utils.Email{}
								mailer.To = formStruct.Sendmailto
								mailer.From = smtpToUse.Username
								mailer.Replyto = "team@no-reply.com"
								mailer.FromName = "TEAM NO-REPLY"
								mailer.Subject = fmt.Sprintf("Your -%v- Campaign Is Ready!", tableCampaign.Title)

								mailer.Message = fmt.Sprintf(`%v <img src="%v/pixelthumb-%v-it.png"/>`,
									messageBytes.String(), formStruct.Prefix, jwtTokenPixel)

								mySMTP := utils.SMTP{
									Port: int(smtpToUse.Port), Server: smtpToUse.Server,
									Username: smtpToUse.Username, Password: smtpToUse.Password,
								}

								sError := utils.SendEmail(mailer, mySMTP)
								if sError != "" {
									smtpToUse.Workflow = "disabled"
									smtpToUse.Description = sError
									message.Message = fmt.Sprintf("SMTP [%s] is DISABLED - Please Fix!!!\nError: %s", tableCampaign.Smtp, sError)
								} else {
									message.Message = fmt.Sprintf("Generated Link sent to %v", mailer.To)
								}
								smtpToUse.Update(smtpToUse.ToMap())
							}
						}
					} else {
						message.Message = fmt.Sprintf("Could not send email! \nSMTP [%s] is DISABLED!! \nPlease Fix!!!", tableCampaign.Smtp)
					}
				}
			*/
			//Send Generated Link via Custom Email
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiCampaignSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Campaigns{}
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

func apiCampaignGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Campaigns{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		if !table.StartDate.IsZero() {
			tableMap["StartDate"], tableMap["StartTime"] = utils.DateTimeSplit(table.StartDate)
		}

		if !table.StopDate.IsZero() {
			tableMap["StopDate"], tableMap["StopTime"] = utils.DateTimeSplit(table.StopDate)
		}

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiCampaignPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Campaigns{}
		table.FillStruct(tableMap)

		claims := utils.VerifyJWT(httpRes, httpReq)
		if claims == nil {
			message.Code = http.StatusInternalServerError
			message.Message = "Invalid JWT Claims"
		} else {
			if table.ID == 0 || tableMap["Code"] == nil || tableMap["Code"].(string) == "" {
				tableMap["Code"] = strings.ToLower(fmt.Sprintf("CMP%v-%v-%v", claims["ID"], time.Now().Format("060102"), utils.RandomString(3)))
				if !apiBlock("partner", claims) || !apiBlock("customer", claims) {
					tableMap["Workflow"] = "draft"
				}
			}
		}

		if table.ID == 0 && message.Code == http.StatusOK {
			switch {
			case strings.HasPrefix(table.Campaign, "data:application/zip"),
				strings.HasPrefix(table.Campaign, "data:multipart/x-zip"),
				strings.HasPrefix(table.Campaign, "data:application/zip-compressed"),
				strings.HasPrefix(table.Campaign, "data:application/x-zip-compressed"),
				strings.HasPrefix(table.Campaign, "data:application/octet-stream"):
			default:
				message.Code = http.StatusInternalServerError
				message.Message += "No Campaign File selected (missing zip)"
			}
		}

		if message.Code == http.StatusOK {
			if tableMap["Campaign"] != nil {
				if table.Campaign = utils.SaveCustomBase64File(tableMap["Campaign"].(string), "campaigns/"+tableMap["Code"].(string), "campaign.zip"); table.Campaign != "" {
					utils.ZipExtract(table.Campaign, strings.Replace(table.Campaign, "campaign.zip", "", 1))
					tableMap["Campaign"] = table.Campaign
				}
			}
			if table.Image = utils.SaveBase64Image(table.Image, ""); table.Image != "" {
				tableMap["Image"] = table.Image
			}

			if tableMap["StartTime"] == nil {
				tableMap["StartTime"] = ""
			}

			if tableMap["StartDate"] != nil {
				table.StartDate = utils.DateTimeMerge(tableMap["StartDate"].(string), tableMap["StartTime"].(string), config.Get().Timezone)
			} else {
				table.StartDate = time.Now()
			}

			if tableMap["StopTime"] == nil {
				tableMap["StopTime"] = ""
			}

			if tableMap["StopDate"] != nil {
				table.StopDate = utils.DateTimeMerge(tableMap["StopDate"].(string), tableMap["StopTime"].(string), config.Get().Timezone)
			} else {
				table.StopDate = table.StartDate.AddDate(0, 0, 1)
			}

			tableMap["StartDate"] = table.StartDate
			tableMap["StopDate"] = table.StopDate

			if table.ID == 0 {
				table.FillStruct(tableMap)
				table.Create(table.ToMap())
			} else {
				table.Update(tableMap)
			}

			message.Body = table.ID
			message.Message = RecordSaved
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}
