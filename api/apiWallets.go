package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerWallets(middlewares alice.Chain, router *Router) {
	router.Get("/api/wallets", middlewares.ThenFunc(apiWalletsGet))
	router.Post("/api/wallets", middlewares.ThenFunc(apiWalletsPost))
	router.Post("/api/wallets/search", middlewares.ThenFunc(apiWalletsSearch))
	router.Post("/api/wallets/select", middlewares.ThenFunc(apiWalletsSelect))
}

func apiWalletsSelect(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")
	message := &Message{Code: http.StatusInternalServerError}

	jwtClaims := utils.VerifyJWT(httpRes, httpReq)
	if jwtClaims == nil {
		message.Body = map[string]string{"Redirect": "/"}
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	var formStruct struct {
		WalletID uint64
		Title    string
	}
	if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
		message.Message = "Error Decoding Form Values " + err.Error()
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	sqlQuery := "select id as walletid, title from wallets where userid = $1 and walletid = $2"
	config.Get().Postgres.Get(&formStruct, sqlQuery, uint64(jwtClaims["ID"].(float64)), formStruct.WalletID)
	if formStruct.WalletID == 0 {
		message.Message = "Wallet ID is invalid"
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	message.Code = http.StatusOK
	message.Message = fmt.Sprintf("%s selected!!", formStruct.Title)

	cookieExpires := time.Now().Add(time.Minute * 60)
	jwtClaims["WalletID"] = formStruct.WalletID
	jwtClaims["exp"] = cookieExpires.Unix()

	if jwtToken, err := utils.GenerateJWT(jwtClaims); err == nil {
		cookieMonster := &http.Cookie{
			Name: config.Get().COOKIE, Value: jwtToken, Expires: cookieExpires, Path: "/",
		}
		http.SetCookie(httpRes, cookieMonster)
		httpReq.AddCookie(cookieMonster)
	}

	json.NewEncoder(httpRes).Encode(message)
}

func apiWalletsGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Wallets{}

		table.GetByID(table.ToMap(), formSearch)

		//check jwtClaims and filter results
		lSkip := true
		jwtClaims := utils.VerifyJWT(httpRes, httpReq)
		if apiBlock("admin", jwtClaims) {
			lSkip = false
		} else {
			var userID uint64
			if jwtClaims != nil && jwtClaims["ID"] != nil {
				userID = uint64(jwtClaims["ID"].(float64))
			}
			if userID == table.UserID {
				lSkip = false
			}
		}
		//check jwtClaims and filter results

		if !lSkip {
			tableMap := table.ToMap()
			delete(tableMap, "UserID")
			delete(tableMap, "Mnemonic")
			delete(tableMap, "ProfileID")
			message.Body = tableMap
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiWalletsPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Wallets{}
		table.FillStruct(tableMap)

		message.Message = RecordSaved
		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {

			//check jwtClaims and filter results
			lSkip := true
			jwtClaims := utils.VerifyJWT(httpRes, httpReq)
			if apiBlock("admin", jwtClaims) {
				lSkip = false
			} else {
				var userID, curUserID uint64
				if jwtClaims != nil && jwtClaims["ID"] != nil {
					userID = uint64(jwtClaims["ID"].(float64))
				}

				sqlQuery := "select userid from wallets where id = $1 limit 1"
				config.Get().Postgres.Get(&curUserID, sqlQuery, table.ID)
				if userID == curUserID {
					lSkip = false
				}
			}
			//check jwtClaims and filter results

			if !lSkip {
				table.Update(tableMap)
			}
		}
		message.Body = table.ID
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiWalletsSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Wallets{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}
		var searchList []interface{}
		var searchResults []database.Wallets

		//check jwtClaims and filter results
		var userID uint64
		lSearchExtra := true
		jwtClaims := utils.VerifyJWT(httpRes, httpReq)
		if apiBlock("admin", jwtClaims) {
			lSearchExtra = false
		} else {
			if jwtClaims != nil && jwtClaims["ID"] != nil {
				userID = uint64(jwtClaims["ID"].(float64))
			}
		}
		//check jwtClaims and filter results
		if !lSearchExtra {
			searchResults = table.Search(table.ToMap(), formSearch)
		} else {
			extra := fmt.Sprintf("userid = %d and ", userID)
			searchResults = table.SearchExtra(table.ToMap(), formSearch, extra)
		}

		for _, result := range searchResults {
			tableMap := result.ToMap()

			delete(tableMap, "Mnemonic")
			// delete(tableMap, "UserID" )
			// delete(tableMap, "ProfileID" )

			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}
