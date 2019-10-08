package database

import (
	"encoding/json"
	"fmt"
	"log"

	"tessa/config"
)

//Newsletters ...
type Newsletters struct {
	Fields

	DaysValid int

	Sendsms, Sendmail bool

	Ordering uint64

	Image, FromEmail, FromName, Replyto, Subject, URL,
	Sms, SmsMessage, SMTP, SenderId, Filepath string
}

//ToMap ...
func (table *Newsletters) ToMap() (mapInterface map[string]interface{}) {
	jsonTable, _ := json.Marshal(table)
	json.Unmarshal(jsonTable, &mapInterface)
	return
}

//FillStruct ...
func (table *Newsletters) FillStruct(tableMap map[string]interface{}) error {
	jsonTable, _ := json.Marshal(tableMap)
	if err := json.Unmarshal(jsonTable, &table); err != nil {
		return err
	}
	return nil
}

//Create ...
func (table *Newsletters) Create(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlInsert(table, tableMap); sqlQuery != "" {
		if err := config.Get().Postgres.Get(&table.ID, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Update ...
func (table *Newsletters) Update(tableMap map[string]interface{}) {
	if sqlQuery, sqlParams := table.sqlUpdate(table, tableMap); sqlQuery != "" {
		if _, err := config.Get().Postgres.Exec(sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//GetByID ...
func (table *Newsletters) GetByID(tableMap map[string]interface{}, searchParams *SearchParams) {
	if sqlQuery, sqlParams := table.sqlSelect(table, tableMap, searchParams); sqlQuery != "" {
		sqlParams = append(sqlParams, searchParams.ID)
		sqlQuery += fmt.Sprintf("id = $%v ", len(sqlParams))
		if err := config.Get().Postgres.Get(table, sqlQuery, sqlParams...); err != nil {
			log.Println(err.Error())
		}
	}
}

//Search ...
func (table *Newsletters) Search(tableMap map[string]interface{}, searchParams *SearchParams) (list []Newsletters) {
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
