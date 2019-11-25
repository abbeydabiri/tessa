package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerNewsletters(middlewares alice.Chain, router *Router) {
	router.Get("/api/newsletters", middlewares.ThenFunc(apiNewsletterGet))
	router.Post("/api/newsletters", middlewares.ThenFunc(apiNewsletterPost))
	router.Post("/api/newsletters/search", middlewares.ThenFunc(apiNewsletterSearch))

	router.Post("/api/newsletter/duplicate", middlewares.ThenFunc(apiNewsletterDuplicate))

	router.GET("/fav/:jwtoken/logo.ico", apiNewsLetterOpenned)
	// go apiJobCheckTransactions()
}

func apiNewsLetterOpenned(httpRes http.ResponseWriter, httpReq *http.Request, httpParams httprouter.Params) {
	jwtToken := httpParams.ByName("jwtoken")
	if jwtClaims := utils.ValidateJWT(jwtToken); jwtClaims != nil {
		UserID := uint64(0)
		if jwtClaims["uid"] != nil {
			UserID = uint64(jwtClaims["uid"].(float64))
		}

		NewsletterID := uint64(0)
		if jwtClaims["nid"] != nil {
			NewsletterID = uint64(jwtClaims["nid"].(float64))
		}

		newsletter := database.Newsletters{}
		formSearch := database.SearchParams{ID: NewsletterID}
		newsletter.GetByID(newsletter.ToMap(), &formSearch)

		//Update Newsletter Recipient if Exists
		sqlQuery := "update followers set workflow='opened' where bucket = 'newsletters' and  bucketid = $1 and userid = $2"
		config.Get().Postgres.Exec(sqlQuery, NewsletterID, UserID)
		//Update Newsletter Recipient if Exists

		//Save Bucket Hit and Reference Newsletter + User
		tableHits := database.Hits{}
		tableHitsMap := make(map[string]interface{})

		tableHits.URL = fmt.Sprintf("%v%v", httpReq.Host, httpReq.URL.String())
		if jwtClaims["url"] != nil {
			tableHits.URL = jwtClaims["url"].(string)
		}
		tableHits.UserAgent = httpReq.UserAgent()
		tableHits.IPAddress = httpReq.RemoteAddr
		tableHits.Code = utils.GetUnixTimestamp()

		tableHits.Title = jwtClaims["Email"].(string)
		if newsletter.Title != "" {
			tableHits.Title = fmt.Sprintf("%v newsletter: %v ", tableHitsMap["Title"], newsletter.Title)
		}

		emailPixelRE := regexp.MustCompile(`^[a-zA-Z0-9][-_.a-zA-Z0-9]*@[-_.a-zA-Z0-9]+?$`)
		if emailPixelRE.MatchString(tableHitsMap["Title"].(string)) {
			tableHits.Title = fmt.Sprintf("%v IP: %v ", tableHitsMap["Title"], httpReq.RemoteAddr)
		}

		tableHits.Description = fmt.Sprintf("Referrer: %v", httpReq.Referer())
		tableHits.Create(tableHits.ToMap())
		//Save Bucket Hit and Reference Newsletter + User
	}
}

func apiNewsletterDuplicate(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	message.Message = "Could not Duplicate Newsletter"
	if message.Code == http.StatusOK {
		table := database.Newsletters{}
		table.GetByID(table.ToMap(), formSearch)

		table.ID = 0
		table.Workflow = "draft"
		table.Title += " - DUPLICATE"

		fileBytes, _ := config.Asset(table.Filepath)
		if fileBytes != nil {
			filepath := strings.ToLower(fmt.Sprintf("NW_%v_%v", time.Now().Format("060102"), utils.RandomString(6)))
			table.Filepath = utils.SaveFile(filepath, "newsletter", fileBytes)
		}
		table.Create(table.ToMap())
		message.Body = table.ID
		message.Message = "Newsletter Duplicated!!"
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiNewsletterSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Newsletters{}
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

func apiNewsletterGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Newsletters{}
		table.GetByID(table.ToMap(), formSearch)

		if table.ID > 0 {
			tableMap := table.ToMap()

			tableMap["RecipientsID"] = apiFollowerContact("newsletters", table.ID)
			FileByte, _ := config.Asset(table.Filepath)
			tableMap["File"] = fmt.Sprintf("%s", FileByte)

			message.Body = tableMap
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiNewsletterPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiSecurePost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Newsletters{}
		table.FillStruct(tableMap)

		if table.Image = utils.SaveBase64Image(table.Image, ""); table.Image != "" {
			tableMap["Image"] = table.Image
		}

		if tableMap["File"] != nil {
			filepath := strings.ToLower(fmt.Sprintf("NW_%v_%v", time.Now().Format("060102"), utils.RandomString(6)))
			if table.Filepath = utils.SaveFile(filepath, "newsletter", []byte(tableMap["File"].(string))); table.Filepath != "" {
				tableMap["Filepath"] = table.Filepath
			}
		}

		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}

		if tableMap["RecipientsID"] != nil {
			followerList := make(map[uint64]string)
			for sID, sFullname := range tableMap["RecipientsID"].(map[string]interface{}) {
				mapID, _ := strconv.ParseUint(sID, 0, 64)
				followerList[mapID] = sFullname.(string)
			}
			apiFollowerListSave(table.Code, table.Title, "newsletters", table.ID, followerList)
		}

		message.Body = table.ID
		message.Message = RecordSaved
	}
	json.NewEncoder(httpRes).Encode(message)
}
