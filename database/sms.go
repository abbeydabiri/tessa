package database

import (
	"encoding/json"
	"fmt"
	"log"

	"tessa/config"
)

//Sms ...
type Sms struct {
	Fields
	Url         string
	Units, Used int
}

//ToMap ...
func (table *Sms) ToMap() (mapInterface map[string]interface{}) {
	jsonTable, _ := json.Marshal(table)
	json.Unmarshal(jsonTable, &mapInterface)
	return
}

//FillStruct ...
func (table *Sms) FillStruct(tableMap map[string]interface{}) error {
	jsonTable, _ := json.Marshal(tableMap)
	if err := json.Unmarshal(jsonTable, &table); err != nil {
		return err
	}
	return nil
}

//Create ...
func (table *Sms) Create(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlInsert(table, tableMap); sqlQuery != "" {
		if err := config.Get().Postgres.Get(&table.ID, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Update ...
func (table *Sms) Update(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlUpdate(table, tableMap); sqlQuery != "" {
		if _, err := config.Get().Postgres.Exec(sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//GetByID ...
func (table *Sms) GetByID(tableMap map[string]interface{}, searchParams *SearchParams) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		sqlParams = append(sqlParams, searchParams.ID)
		sqlQuery += fmt.Sprintf("id = $%v ", len(sqlParams))
		if err := config.Get().Postgres.Get(table, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Search ...
func (table *Sms) Search(tableMap map[string]interface{}, searchParams *SearchParams) (list []Sms) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		searchParams.Text = "%" + searchParams.Text + "%"
		sqlParams = append(sqlParams, searchParams.Text)
		sqlQuery += fmt.Sprintf("%v like $%v order by id desc ", searchParams.Field, len(sqlParams))

		if searchParams.Limit == 0 {
			searchParams.Limit = 100
		}
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
