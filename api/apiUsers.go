package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"
	"golang.org/x/crypto/bcrypt"

	"tessa/config"
	"tessa/database"
)

func apiHandlerUsers(middlewares alice.Chain, router *Router) {
	router.Get("/api/users", middlewares.ThenFunc(apiUsersGet))
	router.Post("/api/users", middlewares.ThenFunc(apiUsersPost))
	router.Post("/api/users/search", middlewares.ThenFunc(apiUsersSearch))
}

func apiUsersGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Users{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()
		Profile := ""
		config.Get().Postgres.Get(&Profile, database.SQLProfile, table.ProfileID)
		tableMap["Profile"] = Profile

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiUsersPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiSecurePost(httpRes, httpReq)

	if message.Code != http.StatusOK {
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	table := database.Users{}
	table.FillStruct(tableMap)

	if table.Username == "" {
		message.Message += "Username is required \n"
		message.Code = http.StatusInternalServerError
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	if tableMap["Password"] != nil && tableMap["Password"].(string) != "" {
		passwordHash, _ := bcrypt.GenerateFromPassword(
			[]byte(tableMap["Password"].(string)), bcrypt.DefaultCost)
		tableMap["Password"] = base64.StdEncoding.EncodeToString(passwordHash)
	}

	if table.ID == 0 {
		table.FillStruct(tableMap)
		// Checks if user has already been created
		searchResults := table.Search(table.ToMap(), &database.SearchParams{Field: "username", Text: tableMap["Username"].(string)})
		if len(searchResults) > 0 {
			message.Code = http.StatusBadRequest
			message.Message = "User already exist!"
			message.Error = "Duplicate User"
			json.NewEncoder(httpRes).Encode(message)
			return
		}
		table.Create(table.ToMap())
	} else {
		table.Update(tableMap)
	}
	message.Body = table.ID
	message.Message = RecordSaved
	json.NewEncoder(httpRes).Encode(message)
}

func apiUsersSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Users{}
		if formSearch.Field == "" {
			formSearch.Field = "Username"
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
