package api

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"

	"tessa/database"
)

func apiHandlerAccountTokens(middlewares alice.Chain, router *Router) {
	router.Get("/api/accounttokens", middlewares.ThenFunc(apiAccountTokensGet))
	router.Post("/api/accounttokens", middlewares.ThenFunc(apiAccountTokensPost))
	router.Post("/api/accounttokens/search", middlewares.ThenFunc(apiAccountTokensSearch))
}

func apiAccountTokensGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.AccountTokens{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiAccountTokensPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.AccountTokens{}
		table.FillStruct(tableMap)

		message.Code = http.StatusInternalServerError
		if table.Title == "" {
			message.Message += "Title is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.WalletID == 0 {
			message.Message += "Wallet ID is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.AccountID == 0 {
			message.Message += "Account ID is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		if table.TokenID == 0 {
			message.Message += "Token ID is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}
		message.Code = http.StatusOK

		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Message = "Token Purchased!!"
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiAccountTokensSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.AccountTokens{}
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
