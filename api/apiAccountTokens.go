package api

import (
	"encoding/json"
	"net/http"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
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

		Token := database.Tokens{}
		config.Get().Postgres.Get(&Token, "select * from tokens where id = $1 limit 1", table.TokenID)
		tableMap["Token"] = Token

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

		var userID, walletID uint64
		if claims := utils.VerifyJWT(httpRes, httpReq); claims != nil {

			if claims["ID"] != nil {
				userID = uint64(claims["ID"].(float64))
			}

			if claims["WalletID"] != nil {
				walletID = uint64(claims["WalletID"].(float64))
			}
		}

		if table.WalletID == 0 || table.AccountID == 0 {
			accountID := uint64(0)
			config.Get().Postgres.Get(&accountID, "select id from accounts where userid = $1 and walletid = $2 limit 1", userID, walletID)

			if accountID > uint64(0) {
				table.AccountID = accountID
				table.WalletID = walletID
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
		message.Code = http.StatusOK
		
		//If Token Account already exists do not replace
		accountID := uint64(0)
		sqlCheck := "select id from accounttokens where tokenid = $1 and walletid = $2 and accountid = $3 limit 1"
		config.Get().Postgres.Get(&accountID, sqlCheck, table.TokenID, table.WalletID, table.AccountID )
		if accountID > uint64(0) {
			message.Body = table.ID
			message.Message = ""
			json.NewEncoder(httpRes).Encode(message)
			return
		}
		//If Token Account already exists do not replace



		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Balance = float64(0.0)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Message = "Token Account Created!!"
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

			Token := database.Tokens{}
			config.Get().Postgres.Get(&Token, "select * from tokens where id = $1 limit 1", result.TokenID)
			tableMap["Token"] = Token
			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}
