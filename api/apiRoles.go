package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
)

func apiHandlerRoles(middlewares alice.Chain, router *Router) {
	router.Get("/api/roles", middlewares.ThenFunc(apiRolesGet))
	router.Post("/api/roles", middlewares.ThenFunc(apiRolesPost))
	router.Post("/api/roles/search", middlewares.ThenFunc(apiRolesSearch))

	router.Get("/api/roles/delete", middlewares.ThenFunc(apiRolesDelete))
}

func apiRolesGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Roles{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		tableMap["UsernamesID"] = apiFollowerUsername("roleusernames", table.ID)
		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiRolesPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiSecurePost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Roles{}
		table.FillStruct(tableMap)

		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Message = RecordSaved

		if tableMap["UsernamesID"] != nil && len(tableMap["UsernamesID"].(map[string]interface{})) > 0 {
			followerList := make(map[uint64]string)
			for sID, sFullname := range tableMap["UsernamesID"].(map[string]interface{}) {
				mapID, _ := strconv.ParseUint(sID, 0, 64)
				followerList[mapID] = sFullname.(string)
			}
			apiFollowerListSave(table.Code, table.Title, "roleusernames", table.ID, followerList)
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiRolesSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Roles{}
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

func apiRolesDelete(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		if formSearch.ID > 0 {
			sqlParams := []interface{}{formSearch.ID}
			sqlQuery := "delete from roles where id = $1 "
			config.Get().Postgres.Exec(sqlQuery, sqlParams...)
			message.Message = "Role Deleted!!"
		} else {
			if message.Message == "" {
				message.Message = "Unable to delete Role!!"
			}
		}

	}
	json.NewEncoder(httpRes).Encode(message)
}
