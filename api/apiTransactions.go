package api

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"

	"tessa/database"
)

func apiHandlerTransactions(middlewares alice.Chain, router *Router) {
	router.Get("/api/transactions", middlewares.ThenFunc(apiTransactionsGet))
	router.Post("/api/transactions", middlewares.ThenFunc(apiTransactionsPost))
	router.Post("/api/transactions/search", middlewares.ThenFunc(apiTransactionsSearch))
}

func apiTransactionsGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Transactions{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTransactionsPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Transactions{}
		table.FillStruct(tableMap)

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

func apiTransactionsSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Transactions{}
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
