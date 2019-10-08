package api

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"

	"tessa/database"
)

func apiHandlerCurrencys(middlewares alice.Chain, router *Router) {
	router.Get("/api/currencys", middlewares.ThenFunc(apiCurrencysGet))
	router.Post("/api/currencys", middlewares.ThenFunc(apiCurrencysPost))
	router.Post("/api/currencys/search", middlewares.ThenFunc(apiCurrencysSearch))
}

func apiCurrencysGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Currencys{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiCurrencysPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Currencys{}
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

func apiCurrencysSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Currencys{}
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
