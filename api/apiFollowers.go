package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
)

type apiFollowerStruct struct {
	ID, UserID, BucketID uint64

	Code, Title, Workflow,
	Description, Bucket string
}

func apiHandlerFollowers(middlewares alice.Chain, router *Router) {
	router.Post("/api/followers/clear", middlewares.ThenFunc(apiFollowerListClear))
	router.Post("/api/followers/reset", middlewares.ThenFunc(apiFollowerListReset))
}

func apiFollowerListClear(httpRes http.ResponseWriter, httpReq *http.Request) {
	message := apiSecure(httpRes, httpReq)
	if message.Code == http.StatusOK {

		var formStruct struct {
			Bucket   string
			BucketID uint64
		}

		if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
			message.Code = http.StatusInternalServerError
			message.Message = "Error Decoding Form Values: " + err.Error()
		}

		if message.Code == http.StatusOK {
			formStruct.Bucket = strings.ToLower(formStruct.Bucket)
			sqlParams := []interface{}{formStruct.Bucket, formStruct.BucketID}
			sqlQuery := "delete from followers where bucket = $1 and bucketid = $2"
			config.Get().Postgres.Exec(sqlQuery, sqlParams...)
			message.Message = "list has been cleared!"
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiFollowerListReset(httpRes http.ResponseWriter, httpReq *http.Request) {
	message := apiSecure(httpRes, httpReq)
	if message.Code == http.StatusOK {

		var formStruct struct {
			Bucket   string
			BucketID uint64
			Workflow string
		}

		if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
			message.Code = http.StatusInternalServerError
			message.Message = "Error Decoding Form Values: " + err.Error()
		}

		if message.Code == http.StatusOK {
			if formStruct.Workflow == "" {
				formStruct.Workflow = "pending"
			}
			formStruct.Bucket = strings.ToLower(formStruct.Bucket)

			sqlParams := []interface{}{formStruct.Workflow, formStruct.Bucket, formStruct.BucketID}
			sqlQuery := "update followers set workflow = $1 where workflow != 'opened' and bucket = $2 and bucketid = $3"
			config.Get().Postgres.Exec(sqlQuery, sqlParams...)
			message.Message = "list has been reset!"

			//Get The Recipients:
			type recipientList struct{ RecipientsID map[uint64]string }
			if formStruct.Bucket == "Newsletters" {
				message.Body = recipientList{
					RecipientsID: apiFollowerContact(formStruct.Bucket, formStruct.BucketID),
				}
			} else {
				message.Body = recipientList{
					RecipientsID: apiFollowerUsername(formStruct.Bucket, formStruct.BucketID),
				}
			}
		}
	}
	json.NewEncoder(httpRes).Encode(message)
}

func apiFollowerRole(Bucket string, BucketID uint64) (FollowersID map[uint64]string) {
	var sqlResults []struct {
		ID    uint64
		Title string
	}

	sqlQuery := "select roles.id as ID, roles.title as Title from "
	sqlQuery += "roles inner join followers on roles.id = followers.roleid where "
	sqlQuery += "followers.bucket = $1 and followers.bucketid = $2 "
	config.Get().Postgres.Select(&sqlResults, sqlQuery, strings.ToLower(Bucket), BucketID)

	FollowersID = make(map[uint64]string)
	for _, result := range sqlResults {
		FollowersID[result.ID] = result.Title
	}
	return
}

func apiFollowerUsername(Bucket string, BucketID uint64) (FollowersID map[uint64]string) {

	var sqlResults []struct {
		ID uint64
		Username, Workflow,
		Title string
	}

	sqlQuery := "select users.id as ID, users.username as Username, followers.workflow as Workflow, "
	sqlQuery += "users.title as Title from users inner join followers on users.id = followers.userid "
	sqlQuery += "where followers.bucket = $1 and followers.bucketid = $2 "
	config.Get().Postgres.Select(&sqlResults, sqlQuery, strings.ToLower(Bucket), BucketID)

	FollowersID = make(map[uint64]string)
	for _, result := range sqlResults {
		sDetails := ""
		if result.Workflow != "" {
			sDetails = fmt.Sprintf("[%v] | ", result.Workflow)
		}
		sDetails += fmt.Sprintf("%v - [%v]", result.Username, result.Title)
		FollowersID[result.ID] = sDetails
	}
	return
}

func apiFollowerContact(Bucket string, BucketID uint64) (FollowersID map[uint64]string) {

	var sqlResults []struct {
		ID uint64
		Fullname, Mobile, Email,
		Workflow, Description string
	}

	sqlQuery := "select profiles.id as ID, profiles.fullname as Fullname, profiles.mobile as Mobile, "
	sqlQuery += "profiles.email as Email, profiles.description as Description, followers.workflow as Workflow "
	sqlQuery += "from profiles inner join followers on profiles.id = followers.userid "
	sqlQuery += "where followers.bucket = $1 and followers.bucketid = $2"
	config.Get().Postgres.Select(&sqlResults, sqlQuery, strings.ToLower(Bucket), BucketID)

	FollowersID = make(map[uint64]string)
	for _, result := range sqlResults {
		sDetails := ""
		if result.Workflow != "" {
			sDetails = fmt.Sprintf("[%v] | ", result.Workflow)
		}

		if result.Description != "" {
			sDetails = fmt.Sprintf("%s [%v] | ", sDetails, result.Description)
		}

		sDetails += fmt.Sprintf("%v - %v %v", result.Fullname, result.Mobile, result.Email)

		FollowersID[result.ID] = sDetails
	}
	return
}

func apiFollowsList(Bucket string, UserID uint64) (BucketsID map[uint64]string) {
	var sqlResults []struct {
		BucketID uint64
		Workflow string
	}
	sqlQuery := "select bucketid, workflow from followers where bucket=$1 and userid = $2 "
	config.Get().Postgres.Select(&sqlResults, sqlQuery, strings.ToLower(Bucket), UserID)

	BucketsID = make(map[uint64]string)
	for _, result := range sqlResults {
		BucketsID[result.BucketID] = result.Workflow
	}
	return
}

func apiFollowerListSave(Code, Title, Bucket string, BucketID uint64, FollowersID map[uint64]string) (statusMessage string) {

	var sqlResults []struct {
		ID, UserID uint64
		Workflow   string
	}
	sqlQuery := "select id, userid, workflow from followers where bucket=$1 and bucketid = $2 "
	if err := config.Get().Postgres.Select(&sqlResults, sqlQuery, strings.ToLower(Bucket), BucketID); err != nil {
		log.Println(err.Error())
	}

	searchList := make(map[string]struct {
		ID, UserID uint64
		Workflow   string
	})

	for _, result := range sqlResults {
		searchList[fmt.Sprintf("%v", result.UserID)] = result
	}

	var followersListSave []map[string]interface{}
	for fID := range FollowersID {
		if len(searchList) > 0 && searchList[fmt.Sprintf("%v", fID)].ID != 0 {
			delete(searchList, fmt.Sprintf("%v", fID))
			continue
		}

		follower := database.Followers{}
		follower.ID = 0
		follower.Workflow = "pending"

		follower.Code = Code
		follower.Title = Title

		switch Bucket {
		case "permissionroles":
			follower.RoleID = fID
		default:
			follower.UserID = fID
		}
		follower.Bucket = Bucket
		follower.BucketID = BucketID

		followersListSave = append(followersListSave, follower.ToMap())
	}

	//Delete Removed Followers
	var sqlDeleteValues []uint64
	for _, delFollower := range searchList {
		sqlDeleteValues = append(sqlDeleteValues, delFollower.ID)
	}
	sqlQuery = "delete from followers where bucket=$1 and bucketid = $2 and id in ($3)"
	config.Get().Postgres.Exec(sqlQuery, strings.ToLower(Bucket), BucketID, sqlDeleteValues)

	log.Println(strings.ToLower(Bucket))
	log.Println(BucketID)
	log.Println(sqlDeleteValues)
	//Delete Removed Followers

	//Update Existing Followers
	sqlQuery = "update followers set code= $1, title = $2 where bucket=$1 and bucketid = $2"
	config.Get().Postgres.Exec(sqlQuery, strings.ToLower(Bucket), BucketID, sqlDeleteValues)
	//Update Existing Followers

	//Create New Follower
	if len(followersListSave) > 0 {
		sqlBulkInsert := database.SQLBulkInsert(&database.Followers{}, followersListSave)
		config.Get().Postgres.Exec(sqlBulkInsert)
	}
	//Create New Followers

	return
}
