package database

import (
	"encoding/json"
	"fmt"
	"log"

	"tessa/config"
)

//Networks ...
type Networks struct {
	Fields

	URL string

	Priority int
}

//Setup ...
func (table *Networks) Setup() {

	log.Println("Setting Up Networks")

	table.Title = "Rinkeby"
	table.URL = "https://rinkeby.infura.io"
	table.Workflow = "enabled"
	table.Priority = 1
	table.Create(table.ToMap())

	table.Title = "Ropsten"
	table.URL = "https://ropsten.infura.io"
	table.Workflow = "enabled"
	table.Priority = 2
	table.Create(table.ToMap())

	table.Title = "Kovan"
	table.URL = "https://kovan.infura.io"
	table.Workflow = "enabled"
	table.Priority = 3
	table.Create(table.ToMap())

	table.Title = "Mainnet"
	table.URL = "https://mainnet.infura.io"
	table.Workflow = "enabled"
	table.Priority = 4
	table.Create(table.ToMap())

}

//ToMap ...
func (table *Networks) ToMap() (mapInterface map[string]interface{}) {
	jsonTable, _ := json.Marshal(table)
	json.Unmarshal(jsonTable, &mapInterface)
	return
}

//FillStruct ...
func (table *Networks) FillStruct(tableMap map[string]interface{}) error {
	jsonTable, _ := json.Marshal(tableMap)
	if err := json.Unmarshal(jsonTable, &table); err != nil {
		return err
	}
	return nil
}

//Create ...
func (table *Networks) Create(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlInsert(table, tableMap); sqlQuery != "" {
		if err := config.Get().Postgres.Get(&table.ID, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Update ...
func (table *Networks) Update(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlUpdate(table, tableMap); sqlQuery != "" {
		if _, err := config.Get().Postgres.Exec(sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//GetByID ...
func (table *Networks) GetByID(tableMap map[string]interface{}, searchParams *SearchParams) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		sqlParams = append(sqlParams, searchParams.ID)
		sqlQuery += fmt.Sprintf("id = $%v ", len(sqlParams))
		if err := config.Get().Postgres.Get(table, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Search ...
func (table *Networks) Search(tableMap map[string]interface{}, searchParams *SearchParams) (list []Networks) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		searchParams.Text = "%" + searchParams.Text + "%"
		sqlParams = append(sqlParams, searchParams.Text)
		sqlQuery += fmt.Sprintf("lower(%v) like lower($%v) order by id desc ", searchParams.Field, len(sqlParams))

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
