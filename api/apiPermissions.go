package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerPermissions(middlewares alice.Chain, router *Router) {
	router.Get("/api/permissions", middlewares.ThenFunc(apiPermissionsGet))
	router.Post("/api/permissions", middlewares.ThenFunc(apiPermissionsPost))
	router.Post("/api/permissions/search", middlewares.ThenFunc(apiPermissionsSearch))

	router.Post("/api/permissions/get", middlewares.ThenFunc(apiPermissionUserGet))
	router.Post("/api/permissions/duplicate", middlewares.ThenFunc(apiPermissionsDuplicate))

	router.Get("/api/permissions/delete", middlewares.ThenFunc(apiPermissionsDelete))
}

func apiPermissionsGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Permissions{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()

		tableMap["RolesID"] = apiFollowerRole("permissionroles", table.ID)
		tableMap["UsernamesID"] = apiFollowerUsername("permissionusernames", table.ID)

		message.Body = tableMap
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiPermissionsPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiSecurePost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Permissions{}
		table.FillStruct(tableMap)

		if table.ID == 0 {
			table.FillStruct(tableMap)
			table.Create(table.ToMap())
		} else {
			table.Update(tableMap)
		}
		message.Body = table.ID
		message.Message = RecordSaved

		if tableMap["RolesID"] != nil && len(tableMap["RolesID"].(map[string]interface{})) > 0 {
			followerList := make(map[uint64]string)
			for sID, sFullname := range tableMap["RolesID"].(map[string]interface{}) {
				mapID, _ := strconv.ParseUint(sID, 0, 64)
				followerList[mapID] = sFullname.(string)
			}
			apiFollowerListSave(table.Code, table.Title, "permissionroles", table.ID, followerList)
		}

		if tableMap["UsernamesID"] != nil && len(tableMap["UsernamesID"].(map[string]interface{})) > 0 {
			followerList := make(map[uint64]string)
			for sID, sFullname := range tableMap["UsernamesID"].(map[string]interface{}) {
				mapID, _ := strconv.ParseUint(sID, 0, 64)
				followerList[mapID] = sFullname.(string)
			}
			apiFollowerListSave(table.Code, table.Title, "permissionusernames", table.ID, followerList)
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiPermissionsSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Permissions{}
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

func apiPermissionsDuplicate(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	message.Message = "Could not Duplicate Post"
	if message.Code == http.StatusOK {
		table := database.Permissions{}
		table.GetByID(table.ToMap(), formSearch)

		table.Workflow = "draft"
		table.Title += " - DUPLICATE"
		table.Code += "-duplicate"

		mapRolesID := apiFollowerRole("permissionroles", table.ID)
		mapUsernamesID := apiFollowerUsername("permissionusernames", table.ID)

		table.ID = 0
		table.Create(table.ToMap())
		message.Body = table.ID
		message.Message = "Permissions Post Duplicated!!"

		apiFollowerListSave(table.Code, table.Title, "permissionusernames", table.ID, mapUsernamesID)
		apiFollowerListSave(table.Code, table.Title, "permissionroles", table.ID, mapRolesID)
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiPermissionUserGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Permissions{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}

		var searchList []map[string]interface{}
		claims := utils.VerifyJWT(httpRes, httpReq)
		userID := uint64(claims["ID"].(float64))

		searchResults := table.Search(table.ToMap(), formSearch)
		for _, result := range searchResults {

			lApplyPermission := false
			UsernamesID := apiFollowerUsername("permissionusernames", result.ID)
			if UsernamesID[userID] != "" {
				lApplyPermission = true
			} else {
				RolesID := apiFollowerRole("permissionroles", result.ID)
				for roleID, _ := range RolesID {
					UsernamesID := apiFollowerUsername("roleusernames", roleID)
					if UsernamesID[userID] != "" {
						lApplyPermission = true
						break
					}
				}
			}

			if lApplyPermission {
				tableMap := result.ToMap()
				for indexName := range tableMap {
					switch indexName {
					default:
						delete(tableMap, indexName)
					case "Workflow", "Action",
						"Method", "Code", "ID":
					}
				}
				searchList = append(searchList, tableMap)
			}
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiPermissionsDelete(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		if formSearch.ID > 0 {
			sqlParams := []interface{}{formSearch.ID}
			sqlQuery := "delete from permissions where id = $1 "
			config.Get().Postgres.Exec(sqlQuery, sqlParams...)
			message.Message = "Permission Deleted!!"
		} else {
			if message.Message == "" {
				message.Message = "Unable to delete Permission!!"
			}
		}

	}
	json.NewEncoder(httpRes).Encode(message)
}
