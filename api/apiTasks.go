package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerTasks(middlewares alice.Chain, router *Router) {
	router.Get("/api/tasks", middlewares.ThenFunc(apiTaskGet))
	router.Post("/api/tasks", middlewares.ThenFunc(apiTaskPost))
	router.Post("/api/tasks/search", middlewares.ThenFunc(apiTaskSearch))

	router.Post("/api/tasks/read", middlewares.ThenFunc(apiTaskRead))
	router.Post("/api/tasks/duplicate", middlewares.ThenFunc(apiTaskDuplicate))
}

func apiTaskRead(httpRes http.ResponseWriter, httpReq *http.Request) {
	message := apiSecure(httpRes, httpReq)
	if message.Code == http.StatusOK {
		var formStruct struct {
			UserID, TaskID uint64
		}

		message.Message = "Task not located"
		if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
			message.Code = http.StatusInternalServerError
			message.Message = "Error Decoding Form Values: " + err.Error()
		}

		if message.Code == http.StatusOK {
			if formStruct.UserID > uint64(0) && formStruct.TaskID > uint64(0) {
				sqlParams := []interface{}{formStruct.UserID, formStruct.TaskID}
				sqlQuery := "update followers set workflow='read' where bucket = 'tasks' and userid = $1 and bucketid = $2"
				config.Get().Postgres.Exec(sqlQuery, sqlParams...)
				message.Message = "Marked Read"
			}
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTaskDuplicate(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	message.Message = "Could not Duplicate Post"
	if message.Code == http.StatusOK {
		table := database.Tasks{}
		table.GetByID(table.ToMap(), formSearch)

		table.ID = 0
		table.Workflow = "draft"
		table.Title += " - DUPLICATE"
		table.Code = utils.TitleToURL(fmt.Sprintf("TSK%v-%v-%v", table.Title, time.Now().Format("06-01-02"), utils.RandomString(3)))

		fileBytes, _ := config.Asset(table.Filepath)
		if fileBytes != nil {
			table.Filepath = utils.SaveFile(table.Code, "tasks", fileBytes)
		}
		table.Create(table.ToMap())
		message.Body = table.ID
		message.Message = "Task Post Duplicated!!"
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTaskSearch(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureSearch(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tasks{}
		if formSearch.Field == "" {
			formSearch.Field = "Title"
		}

		//Only Admins allowwed to view all tasks
		claims := utils.VerifyJWT(httpRes, httpReq)
		formSearch.UserID = uint64(claims["ID"].(float64))
		if !apiBlock("admin", claims) {
			formSearch.UserID = 0
		}
		//Only Admins allowwed to view all tasks
		followingTasks := apiFollowsList("tasks", uint64(claims["ID"].(float64)))

		var searchList []interface{}
		searchResults := table.Search(table.ToMap(), formSearch)
		for _, result := range searchResults {
			tableMap := result.ToMap()
			if followingTasks[result.ID] != "" {
				tableMap["Workflow"] = followingTasks[result.ID]
			}
			searchList = append(searchList, tableMap)
		}
		message.Body = searchList
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTaskGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	formSearch, message := apiSecureGet(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tasks{}
		table.GetByID(table.ToMap(), formSearch)

		lShow := false
		claims := utils.VerifyJWT(httpRes, httpReq)
		UserID := uint64(claims["ID"].(float64))
		if !apiBlock("admin", claims) {
			lShow = true
		}

		if !lShow {
			if table.Createdby == UserID {
				lShow = true
			}

			if table.Updatedby == UserID {
				lShow = true
			}

			if table.FromID == UserID {
				lShow = true
			}

			if table.ThroughID == UserID {
				lShow = true
			}

			if !lShow {
				if followingTasks := apiFollowsList("tasks", uint64(claims["ID"].(float64))); followingTasks[table.ID] != "" {
					lShow = true
				}
			}
		}

		if lShow {
			tableMap := table.ToMap()

			config.Get().Postgres.Get(tableMap["From"], database.SQLProfile, table.FromID)
			config.Get().Postgres.Get(tableMap["Through"], database.SQLProfile, table.ThroughID)

			if !table.StartDate.IsZero() {
				tableMap["StartDate"], tableMap["StartTime"] = utils.DateTimeSplit(table.StartDate)
			}
			if !table.NextDate.IsZero() {
				tableMap["NextDate"], tableMap["NextTime"] = utils.DateTimeSplit(table.NextDate)
			}
			if !table.StopDate.IsZero() {
				tableMap["StopDate"], tableMap["StopTime"] = utils.DateTimeSplit(table.StopDate)
			}
			tableMap["UsernamesID"] = apiFollowerUsername("tasks", table.ID)

			FileByte, _ := config.Asset(table.Filepath)
			tableMap["File"] = fmt.Sprintf("%s", FileByte)

			message.Body = tableMap
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiTaskPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiSecurePost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Tasks{}
		table.FillStruct(tableMap)

		lSave := false
		if table.ID == 0 {
			lSave = true
		} else {
			claims := utils.VerifyJWT(httpRes, httpReq)
			UserID := uint64(claims["ID"].(float64))
			if !apiBlock("admin", claims) {
				lSave = true
			}

			if !lSave {
				tableExisting := database.Tasks{}
				formSearch := new(database.SearchParams)
				formSearch.ID = table.ID
				tableExisting.GetByID(tableExisting.ToMap(), formSearch)

				if table.Createdby == UserID {
					lSave = true
				}

				if table.Updatedby == UserID {
					lSave = true
				}

				if table.FromID == UserID {
					lSave = true
				}

				if table.ThroughID == UserID {
					lSave = true
				}

				if !lSave {
					if followingTasks := apiFollowsList("tasks", uint64(claims["ID"].(float64))); followingTasks[table.ID] != "" {
						lSave = true
					}
				}
			}
		}

		if !lSave {
			message.Message = "You do not have permission to save"
		} else {
			if tableMap["StartDate"] != nil && tableMap["StartTime"] != nil {
				tableMap["StartDate"] = utils.DateTimeMerge(
					tableMap["StartDate"].(string), tableMap["StartTime"].(string), config.Get().Timezone)
			}

			if tableMap["NextDate"] != nil && tableMap["NextTime"] != nil {
				tableMap["NextDate"] = utils.DateTimeMerge(
					tableMap["NextDate"].(string), tableMap["NextTime"].(string), config.Get().Timezone)
			}

			if tableMap["StopDate"] != nil && tableMap["StopTime"] != nil {
				tableMap["StopDate"] = utils.DateTimeMerge(
					tableMap["StopDate"].(string), tableMap["StopTime"].(string), config.Get().Timezone)
			}

			if table.ID == 0 || tableMap["Code"] == nil || tableMap["Code"].(string) == "" {
				tableMap["Code"] = utils.TitleToURL(fmt.Sprintf("%v-%v-%v", table.Title, time.Now().Format("06-01-02"), utils.RandomString(3)))
			}

			if tableMap["File"] != nil {
				if table.Filepath = utils.SaveFile(tableMap["Code"].(string), "tasks", []byte(tableMap["File"].(string))); table.Filepath != "" {
					tableMap["Filepath"] = table.Filepath
				}
			}

			if table.ID == 0 {
				table.FillStruct(tableMap)
				table.Create(table.ToMap())
			} else {
				table.Update(tableMap)
			}

			if tableMap["UsernamesID"] != nil && len(tableMap["UsernamesID"].(map[string]interface{})) > 0 {
				followerList := make(map[uint64]string)
				for sID, sFullname := range tableMap["UsernamesID"].(map[string]interface{}) {
					log.Printf(" sID: %v - sFullname: %v \n", sID, sFullname)
					mapID, _ := strconv.ParseUint(sID, 0, 64)
					followerList[mapID] = sFullname.(string)
				}
				apiFollowerListSave(table.Code, table.Title, "tasks", table.ID, followerList)
			}

			message.Body = table.ID
			message.Message = RecordSaved
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}
