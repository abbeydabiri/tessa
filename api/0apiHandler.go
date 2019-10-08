package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

const (
	//IsRequired ...
	IsRequired = " is Required \n" //isRequired is required

	//RecordSaved ...
	RecordSaved = "Record Saved" //recordSaved is saved

	//TitleRequired ...
	TitleRequired = "Title is Required \n"

	//WorkflowRequired ...
	WorkflowRequired = "Status is Required \n"
)

//Message ...
type Message struct {
	Code           int
	Message, Error string
	Body           interface{}
}

// JSONTime ...
type JSONTime time.Time

//MarshalJSON ...
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("02/01/2006 03:04:05 PM"))
	return []byte(stamp), nil
}

func verifyID(httpRes http.ResponseWriter, httpReq *http.Request, claims jwt.MapClaims) {
	if claims == nil {
		http.Redirect(httpRes, httpReq, "/", http.StatusTemporaryRedirect)
		return
	}

	if claims["ID"] == nil {
		http.Redirect(httpRes, httpReq, "/", http.StatusTemporaryRedirect)
		return
	}
}

func apiHandler(middlewares alice.Chain, router *Router) {

	// router.POST("/db/init/*table", func(httpRes http.ResponseWriter, httpReq *http.Request, httpParams httprouter.Params) {
	router.GET("/db/init/*table", func(httpRes http.ResponseWriter, httpReq *http.Request, httpParams httprouter.Params) {
		// claims := utils.VerifyJWT(httpRes, httpReq)
		// if claims == nil {
		// 	http.Redirect(httpRes, httpReq, "/404.html", http.StatusTemporaryRedirect)
		// 	return
		// }

		// if claims["IsAdmin"].(bool) && claims["Username"].(string) == "root" {
		errMessages := database.Init(httpParams.ByName("table"))

		message := new(Message)
		message.Code = http.StatusOK
		message.Message = strings.Join(errMessages, "\n")
		json.NewEncoder(httpRes).Encode(message)
		return
		// }

		// http.Redirect(httpRes, httpReq, "/404.html", http.StatusTemporaryRedirect)
	})

	router.GET("/dashboard/*page", func(httpRes http.ResponseWriter, httpReq *http.Request, httpParams httprouter.Params) {
		claims := utils.VerifyJWT(httpRes, httpReq)
		verifyID(httpRes, httpReq, claims)
		fileServe(httpRes, httpReq)
		return
	})

	//Authentication Functions --> Below
	apiHandlerAuth(middlewares, router)
	apiHandlerImport(middlewares, router)
	apiHandlerDashboard(middlewares, router)
	apiHandlerWeb3(middlewares, router)

	//Other API Functions
	apiHandlerAccounts(middlewares, router)
	apiHandlerAccountTokens(middlewares, router)
	apiHandlerActivations(middlewares, router)
	apiHandlerCampaigns(middlewares, router)
	apiHandlerCurrencys(middlewares, router)
	apiHandlerContacts(middlewares, router)
	apiHandlerFollowers(middlewares, router)
	apiHandlerHits(middlewares, router)
	apiHandlerNetworks(middlewares, router)
	apiHandlerNewsletters(middlewares, router)
	apiHandlerPermissions(middlewares, router)
	apiHandlerProfiles(middlewares, router)
	apiHandlerRoles(middlewares, router)
	apiHandlerSeocontents(middlewares, router)
	apiHandlerSettings(middlewares, router)
	apiHandlerSmtps(middlewares, router)
	apiHandlerTasks(middlewares, router)
	apiHandlerTokens(middlewares, router)
	apiHandlerTransactions(middlewares, router)
	apiHandlerUsers(middlewares, router)
	apiHandlerWallets(middlewares, router)
}

func cut(name string) string {
	name = strings.TrimSuffix(name, "/")
	dir, _ := path.Split(name)
	return dir
}

func fileServe(httpRes http.ResponseWriter, httpReq *http.Request) {

	folderPath := "uiweb"
	switch httpReq.Host {
	case config.Get().Appurl:
		folderPath = "uiapp"
	case config.Get().Adminurl:
		folderPath = "uiadmin"
	}

	urlPath := strings.Replace(httpReq.URL.Path, "//", "/", -1)
	if strings.HasSuffix(urlPath, "/") {
		urlPath = path.Join(urlPath, "index.html")
	}

	var err error
	var dataBytes []byte

	if dataBytes, err = config.Asset(folderPath + urlPath); err != nil {
		for urlPath != "/" {
			urlPath = cut(urlPath)
			newPath := path.Join(urlPath, "index.html")
			if dataBytes, err = config.Asset(folderPath + newPath); err == nil {
				break
			}
		}
	}

	httpRes.Header().Set("Cache-Control", "max-age=0, must-revalidate")
	httpRes.Header().Set("Pragma", "no-cache")
	httpRes.Header().Set("Expires", "0")

	httpRes.Header().Add("Content-Type", config.ContentType(urlPath))
	if !strings.Contains(httpReq.Header.Get("Accept-Encoding"), "gzip") {
		httpRes.Write(dataBytes)
		return
	}
	gzipWrite(dataBytes, httpRes)
}

func addSEOContent(urlPath string, dataBytes []byte, httpReq *http.Request) []byte {

	if !strings.Contains(urlPath, "index.html") {
		return dataBytes
	}

	var seoPage struct {
		SeoTitle, SeoContent string
	}

	searchResults := []database.Seocontents{}
	config.Get().Postgres.Select(&searchResults, "select * from seocontents where workflow = 'enabled'")

	httpReqPath := httpReq.Host + httpReq.URL.String()
	for _, result := range searchResults {
		lUpdate := false
		switch result.Filter {
		case "HasPrefix":
			if strings.HasPrefix(httpReqPath, result.URL) {
				lUpdate = true
			}
		case "HasSuffix":
			if strings.HasSuffix(httpReqPath, result.URL) {
				lUpdate = true
			}
		case "Contains":
			if strings.Contains(httpReqPath, result.URL) {
				lUpdate = true
			}
		}

		if lUpdate {
			if len(result.Code) > 0 {
				seoPage.SeoTitle = result.Code
			}
			if len(result.Description) > 0 {
				seoPage.SeoContent = result.Description
			}
		}
	}

	if seoPage.SeoTitle == "" && seoPage.SeoContent == "" {
		return dataBytes
	}

	var err error
	var pageBytes bytes.Buffer
	var pageTemplate *template.Template
	if pageTemplate, err = template.New("index").Parse(string(dataBytes)); err != nil {
		log.Printf(err.Error())
		return dataBytes
	}

	if err = pageTemplate.Execute(&pageBytes, seoPage); err != nil {
		log.Printf(err.Error())
		return dataBytes
	}

	return pageBytes.Bytes()
}
