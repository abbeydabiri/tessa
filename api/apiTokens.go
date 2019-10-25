package api

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"

	"tessa/database"
)

func apiHandlerTokens(middlewares alice.Chain, router *Router) {
	router.Get("/api/tokens", middlewares.ThenFunc(apiTokensGet))
	router.Post("/api/tokens", middlewares.ThenFunc(apiTokensPost))
	router.Post("/api/tokens/search", middlewares.ThenFunc(apiTokensSearch))
}

func apiTokensGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tokens{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTokensPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tokens{}
		table.FillStruct(tableMap)

		message.Code = http.StatusInternalServerError
		if table.Company == "" {
			message.Message += "Company is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.RC == "" {
			message.Message += "Company RC is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.Title == "" {
			message.Message += "Token Title is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.Symbol == "" {
			message.Message += "Token Symbol is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.Project == "" {
			message.Message += "Token Project is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.ProjectCost < 1.00 {
			message.Message += "Token Project Cost is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.MaxTotalSupply < 1 {
			message.Message += "Token Total Supply is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Code = http.StatusOK
		message.Message = RecordSaved
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTokensSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tokens{}
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
