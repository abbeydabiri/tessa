package api

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
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

		message.Code = http.StatusInternalServerError
		if table.Title == "" {
			message.Message += "Title is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		var userID, walletID uint64
		if claims := utils.VerifyJWT(httpRes, httpReq); claims != nil {
			if claims["ID"] != nil {
				userID = uint64(claims["ID"].(float64))
			}

			if claims["WalletID"] != nil {
				walletID = uint64(claims["WalletID"].(float64))
			}
		}

		account := database.Accounts{}
		if table.WalletID == 0 || table.AccountID == 0 {
			config.Get().Postgres.Get(&account, "select * from accounts where userid = $1 and walletid = $2 limit 1", userID, walletID)

			if account.ID > uint64(0) {
				table.WalletID = walletID
				table.AccountID = account.ID
			}
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

		// amount := big.NewInt(0)
		// if amount.Cmp(table.Amount) == 0 {
		if table.Amount == 0 {
			message.Message += "Amount is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		//Check if From Address is mine then it is a debit
		if table.FromAddress == "" {
			message.Message += "From Address is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		//Check if To Address is mine then it is a credit
		if table.ToAddress == "" {
			message.Message += "To Address is required \n"
			json.NewEncoder(httpRes).Encode(message)
			return
		}

		//From and To Address cannot be the same

		if table.ID == 0 {

			//Update Balance

			//Update Balance

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
