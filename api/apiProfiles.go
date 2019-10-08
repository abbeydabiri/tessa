package api

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/justinas/alice"

	"tessa/database"
	"tessa/utils"
)

func apiHandlerProfiles(middlewares alice.Chain, router *Router) {
	router.Get("/api/profiles", middlewares.ThenFunc(apiProfilesGet))
	router.Post("/api/profiles", middlewares.ThenFunc(apiProfilesPost))
	router.Post("/api/profiles/search", middlewares.ThenFunc(apiProfilesSearch)) // router.Post("/api/profiles/welcomemail", middlewares.ThenFunc(apiProfilesPostWelcomeMail))
}

func apiProfilesGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Profiles{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiProfilesPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Profiles{}
		table.FillStruct(tableMap)


		fullname := ""
		if table.Title != "" {
			fullname  = table.Title
		}

		if table.Firstname != "" {
			fullname = fmt.Sprintf("%s %s", fullname, table.Firstname)
		}

		if table.Lastname != "" {
			fullname = fmt.Sprintf("%s %s", fullname, table.Lastname)
		}
		table.Fullname = fullname

		if table.Fullname == "" {
			message.Message += "Fullname is required \n"
			message.Code = http.StatusInternalServerError
			json.NewEncoder(httpRes).Encode(message)
			return
		}


		if table.Mobile == "" {
			message.Message += "Mobile is required \n"
			message.Code = http.StatusInternalServerError
			json.NewEncoder(httpRes).Encode(message)
			return
		}


		if table.Image = utils.SaveBase64Image(table.Image, ""); table.Image != "" {
			tableMap["Image"] = table.Image
		}



		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Message = RecordSaved
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiProfilesSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Profiles{}
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
