package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerContacts(middlewares alice.Chain, router *Router) {
	router.Get("/api/contacts", middlewares.ThenFunc(apiContactsGet))
	router.Post("/api/contacts", middlewares.ThenFunc(apiContactsPost))
	router.Post("/api/contacts/search", middlewares.ThenFunc(apiContactsSearch))
}

func apiContactsGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Contacts{}

		table.GetByID(table.ToMap(), formSearch)

		//check jwtClaims and filter results
		lSkip := true
		jwtClaims := utils.VerifyJWT(httpRes, httpReq)
		if apiBlock("admin", jwtClaims) {
			var userID uint64
			if jwtClaims != nil && jwtClaims["ID"] != nil {
				userID = uint64(jwtClaims["ID"].(float64))
			}
			if userID == table.UserID {
				lSkip = false
			}
		} else {
			lSkip = false
		}
		//check jwtClaims and filter results

		if !lSkip {
			tableMap := table.ToMap()
			delete(tableMap, "UserID")
			message.Body = tableMap
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiContactsPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Contacts{}
		table.FillStruct(tableMap)

		fullname := ""
		if table.Title != "" {
			fullname = table.Title
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

		//Find Crypto Address tied to Mobile:
		var walletAddress string
		sqlQuery := "select address from accounts where userid = (select id from users where username = $1) limit 1"

		log.Println(sqlQuery)
		log.Println(table.Mobile)
		config.Get().Postgres.Get(&walletAddress, sqlQuery, table.Mobile)
		if walletAddress == "" {
			walletAddress = "0x0"
		}
		table.Address = walletAddress
		//Find Crypto Address tied to Mobile:

		//check jwtClaims and filter results
		lSkip := true
		jwtClaims := utils.VerifyJWT(httpRes, httpReq)
		if apiBlock("admin", jwtClaims) {
			if jwtClaims != nil && jwtClaims["ID"] != nil {
				table.UserID = uint64(jwtClaims["ID"].(float64))
			}
		} else {
			lSkip = false
		}
		//check jwtClaims and filter results

		if table.ID == 0 {
			table.Create(table.ToMap())
		} else {
			var curUserID uint64
			sqlQuery := "select userid from contacts where id = $1 limit 1"
			config.Get().Postgres.Get(&curUserID, sqlQuery, table.ID)
			if table.UserID == curUserID {
				lSkip = false
			}

			if !lSkip {
				table.Update(table.ToMap())
			}
		}
		message.Message = RecordSaved
		message.Body = table.ID
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiContactsSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiGenericSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Contacts{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}
		var searchList []interface{}
		var searchResults []database.Contacts

		//check jwtClaims and filter results
		var userID uint64
		lSearchExtra := true
		jwtClaims := utils.VerifyJWT(httpRes, httpReq)
		if apiBlock("admin", jwtClaims) { //not admin
			if jwtClaims != nil && jwtClaims["ID"] != nil {
				userID = uint64(jwtClaims["ID"].(float64))
			}
		} else {
			lSearchExtra = false
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

			// delete(tableMap, "Mnemonic")
			// delete(tableMap, "UserID" )
			// delete(tableMap, "ProfileID" )

			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}
