package api

import (
	"encoding/json"
	// "log"
	"net/http"

	"github.com/justinas/alice"

	"tessa/database"
	"tessa/config"
	"tessa/utils"
)

func apiHandlerDashboard(middlewares alice.Chain, router *Router) {
	router.Get("/api/dashboard", middlewares.ThenFunc(apiDashboardGet))
	router.Post("/api/dashboard", middlewares.ThenFunc(apiDashboardPost))
}

func apiDashboardGet(httpRes http.ResponseWriter, httpReq *http.Request) {
	message := apiGeneric(httpRes, httpReq)
	if message.Code == http.StatusOK {
		claims := utils.VerifyJWT(httpRes, httpReq)
		if claims["ProfileID"] == nil {
			json.NewEncoder(httpRes).Encode(message)
			return
		}


		formSearch := &database.SearchParams{ID: uint64(claims["ProfileID"].(float64))}

		table := database.Profiles{}
		table.GetByID(table.ToMap(), formSearch)

		tableMap := table.ToMap()
		delete(tableMap, "ID")
		delete(tableMap, "Createdate")
		delete(tableMap, "Updatedate")


		
		address := ""
		sqlQueryAddress := "select address from accounts where profileid = $1 limit 1"
		config.Get().Postgres.Get(&address, sqlQueryAddress, table.ID)
		tableMap["Address"] = address


		message.Body = tableMap

	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiDashboardPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	tableMap, message := apiGenericPost(httpRes, httpReq)
	if message.Code == http.StatusOK {
		table := database.Profiles{}
		table.FillStruct(tableMap)

		for indexName := range tableMap {
			switch indexName {
			default:
				delete(tableMap, indexName)
			case "Mobile", "Address",
				"Street", "Description", "Image":
			}
		}

		if table.Image = utils.SaveBase64Image(table.Image, ""); table.Image != "" {
			tableMap["Image"] = table.Image
		}

		claims := utils.VerifyJWT(httpRes, httpReq)
		tableMap["ID"] = uint64(claims["ID"].(float64))

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
