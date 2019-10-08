package api

import (
	"compress/gzip"
	"context"
	"io"
	"strings"

	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

//***** GENERIC FUNCTIONS TO Handle HTTP Requests and Resposnse *****//

//Router ...
type Router struct { // Router struct would carry the httprouter instance,
	*httprouter.Router //so its methods could be verwritten and replaced with methds with wraphandler
}

//Get ...
func (router *Router) Get(path string, handler http.Handler) {
	router.GET(path, wrapHandler(handler)) // Get is an endpoint to only accept requests of method GET
}

//Post is an endpoint to only accept requests of method POST
func (router *Router) Post(path string, handler http.Handler) {
	router.POST(path, wrapHandler(handler))
}

//Put is an endpoint to only accept requests of method PUT
func (router *Router) Put(path string, handler http.Handler) {
	router.PUT(path, wrapHandler(handler))
}

//Delete is an endpoint to only accept requests of method DELETE
func (router *Router) Delete(path string, handler http.Handler) {
	router.DELETE(path, wrapHandler(handler))
}

//NewRouter is a wrapper that makes the httprouter struct a child of the router struct
func NewRouter() *Router {
	return &Router{httprouter.New()}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func gzipWrite(dataBytes []byte, httpRes http.ResponseWriter) {
	httpRes.Header().Set("Content-Encoding", "gzip")
	gzipHandler := gzip.NewWriter(httpRes)
	defer gzipHandler.Close()
	httpResGzip := gzipResponseWriter{Writer: gzipHandler, ResponseWriter: httpRes}
	httpResGzip.Write(dataBytes)
}

func wrapHandler(httpHandler http.Handler) httprouter.Handle {
	return func(httpRes http.ResponseWriter, httpReq *http.Request, httpParams httprouter.Params) {
		ctx := context.WithValue(httpReq.Context(), "params", httpParams)
		httpReq = httpReq.WithContext(ctx)

		if !strings.Contains(httpReq.Header.Get("Accept-Encoding"), "gzip") {
			httpHandler.ServeHTTP(httpRes, httpReq)
			return
		}

		httpRes.Header().Set("Content-Encoding", "gzip")
		gzipHandler := gzip.NewWriter(httpRes)
		defer gzipHandler.Close()
		httpResGzip := gzipResponseWriter{Writer: gzipHandler, ResponseWriter: httpRes}
		httpHandler.ServeHTTP(httpResGzip, httpReq)
	}
}

//***** GENERIC FUNCTIONS TO Handle HTTP Requests and Resposnse *****//

//StartRouter ...
func StartRouter() {

	totalUsers := 0
	if config.Get().Postgres.Get(&totalUsers, "select count(id) from users"); totalUsers == 0 {
		utils.SaveFileToPath("adminurl", "config", []byte(config.Get().Adminurl))
		database.Init("/all")
	}

	middlewares := alice.New()
	router := NewRouter()

	apiHandler(middlewares, router)

	router.NotFound = middlewares.ThenFunc(
		func(httpRes http.ResponseWriter, httpReq *http.Request) {
			frontend := strings.Split(httpReq.URL.Path[1:], "/")
			switch frontend[0] {
			case "logout":

				apiAuthLogout(httpRes, httpReq)
				http.Redirect(httpRes, httpReq, "/", http.StatusTemporaryRedirect)

			default:
				fileServe(httpRes, httpReq)

			}
		})

	allowOriginFunc := func(origin string) bool {
		return true
	}
	mainHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  allowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Accept-Encoding",
			"Content-Type", "X-Auth-Token", "X-CSRF-Token",
			"Authorization", "*"},
	}).Handler(router)

	sMessage := "serving @ " + config.Get().Address
	println(sMessage)
	log.Println(sMessage)
	log.Fatal(http.ListenAndServe(config.Get().Address, mainHandler))
}
