package database

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"tessa/config"
)

//Tasks ..
type Tasks struct {
	Fields
	StartDate, NextDate, StopDate time.Time
	
	Priority, Filepath string

	FromID, ThroughID, Interval uint64

	Mon, Tue, Wed, Thu,
	Fri, Sat, Sun, Daily bool
}

//ToMap ...
func (table *Tasks) ToMap() (mapInterface map[string]interface{}) {
	jsonTable, _ := json.Marshal(table)
	json.Unmarshal(jsonTable, &mapInterface)
	return
}

//FillStruct ...
func (table *Tasks) FillStruct(tableMap map[string]interface{}) error {
	jsonTable, _ := json.Marshal(tableMap)
	if err := json.Unmarshal(jsonTable, &table); err != nil {
		return err
	}
	return nil
}

//Create ...
func (table *Tasks) Create(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlInsert(table, tableMap); sqlQuery != "" {
		if err := config.Get().Postgres.Get(&table.ID, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Update ...
func (table *Tasks) Update(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlUpdate(table, tableMap); sqlQuery != "" {
		if _, err := config.Get().Postgres.Exec(sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//GetByID ...
func (table *Tasks) GetByID(tableMap map[string]interface{}, searchParams *SearchParams) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		sqlParams = append(sqlParams, searchParams.ID)
		sqlQuery += fmt.Sprintf("id = $%v ", len(sqlParams))
		if err := config.Get().Postgres.Get(table, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Search ...
func (table *Tasks) Search(tableMap map[string]interface{}, searchParams *SearchParams) (list []Tasks) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		searchParams.Text = "%" + searchParams.Text + "%"

		//Use the createdby to
		sqlParams = append(sqlParams, searchParams.Text)
		sqlQuery += fmt.Sprintf("lower(%v) like lower($%v)  ", searchParams.Field, len(sqlParams))

		if searchParams.UserID > 0 {
			sqlParams = append(sqlParams, searchParams.UserID)
			sqlQuery += fmt.Sprintf("and (createdby = $%v or updatedby = $%v or fromid = $%v or throughid = $%v or ownerid = $%v or id in (select bucketid from followers where bucket='tasks' and userid = $%v)) ",
				len(sqlParams), len(sqlParams), len(sqlParams), len(sqlParams), len(sqlParams), len(sqlParams))
		}
		sqlQuery += " order by id desc "

		sqlParams = append(sqlParams, searchParams.Limit)
		sqlQuery += fmt.Sprintf("limit $%v ", len(sqlParams))

		sqlParams = append(sqlParams, searchParams.Skip)
		sqlQuery += fmt.Sprintf("offset $%v ", len(sqlParams))
		if err := config.Get().Postgres.Select(&list, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
	return
}
